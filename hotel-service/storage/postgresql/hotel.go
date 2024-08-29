package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	hotelpb "hotel-service/protos/hotel"

	_ "github.com/lib/pq"
)

type Storage interface {
	InsertHotel(ctx context.Context, hotel *hotelpb.Hotel) (int, error)
	ListHotels(ctx context.Context) ([]*hotelpb.Hotel, error)
	GetHotelDetails(ctx context.Context, hotelID string) (*hotelpb.Hotel, error)
	CheckRoomAvailability(ctx context.Context, hotelID string) ([]*hotelpb.RoomAvailability, error)
	CheckHotelIDSql(ctx context.Context, hotelID string) (bool, error)
	UpdateRoomAvailability(ctx context.Context, ID string, availability bool) error
	GetRoomIDSql(ctx context.Context, req *hotelpb.GetRoomIDRequest) (*hotelpb.GetRoomIDResponse, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{
		db: db,
	}
}

func (s *PostgresStorage) InsertHotel(ctx context.Context, hotel *hotelpb.Hotel) (int, error) {
	if hotel.Rating < 0.0 || hotel.Rating > 9.9 {
		return 0, fmt.Errorf("rating must be between 0.0 and 9.9")
	}
	var hotelID int
	err := s.db.QueryRowContext(ctx,
		`INSERT INTO hotels (name, location, rating, address) 
		 VALUES ($1, $2, $3, $4) 
		 RETURNING id`,
		hotel.Name, hotel.Location, hotel.Rating, hotel.Address).Scan(&hotelID)

	if err != nil {
		return 0, err
	}

	for _, room := range hotel.Rooms {
		if room.PricePerNight < 0 || room.PricePerNight > 9999999999.99 {
			return 0, fmt.Errorf("price per night must be between 0.00 and 9999999999.99")
		}
		_, err := s.db.ExecContext(ctx,
			`INSERT INTO rooms (hotel_id, room_type, price_per_night, availability) 
			 VALUES ($1, $2, $3, $4)`,
			hotelID, room.RoomType, room.PricePerNight, room.Availability)
		if err != nil {
			return 0, err
		}
	}

	return hotelID, nil
}

func (s *PostgresStorage) ListHotels(ctx context.Context) ([]*hotelpb.Hotel, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, name, location, rating, address FROM hotels")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hotels []*hotelpb.Hotel
	for rows.Next() {
		var hotel hotelpb.Hotel
		if err := rows.Scan(&hotel.HotelID, &hotel.Name, &hotel.Location, &hotel.Rating, &hotel.Address); err != nil {
			return nil, err
		}
		hotels = append(hotels, &hotel)
	}

	return hotels, nil
}

func (s *PostgresStorage) GetHotelDetails(ctx context.Context, hotelID string) (*hotelpb.Hotel, error) {
	query := `
		SELECT id, name, location, rating, address
		FROM hotels
		WHERE id = $1
	`
	row := s.db.QueryRowContext(ctx, query, hotelID)

	var hotel hotelpb.Hotel
	if err := row.Scan(&hotel.HotelID, &hotel.Name, &hotel.Location, &hotel.Rating, &hotel.Address); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("hotel with ID %s not found", hotelID)
		}
		return nil, err
	}

	// Fetch rooms
	roomRows, err := s.db.QueryContext(ctx, "SELECT room_type, price_per_night, availability FROM rooms WHERE hotel_id = $1", hotelID)
	if err != nil {
		return nil, err
	}
	defer roomRows.Close()

	var rooms []*hotelpb.Room
	for roomRows.Next() {
		var room hotelpb.Room
		if err := roomRows.Scan(&room.RoomType, &room.PricePerNight, &room.Availability); err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}

	hotel.Rooms = rooms

	return &hotel, nil
}

func (s *PostgresStorage) CheckRoomAvailability(ctx context.Context, hotelID string) ([]*hotelpb.RoomAvailability, error) {
	query := `
		SELECT room_type, COUNT(*) as available_rooms
		FROM rooms
		WHERE hotel_id = $1 AND availability = true
		GROUP BY room_type
	`
	rows, err := s.db.QueryContext(ctx, query, hotelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roomAvailabilities []*hotelpb.RoomAvailability
	for rows.Next() {
		var roomAvailability hotelpb.RoomAvailability
		if err := rows.Scan(&roomAvailability.RoomType, &roomAvailability.AvailableRooms); err != nil {
			return nil, err
		}
		roomAvailabilities = append(roomAvailabilities, &roomAvailability)
	}

	return roomAvailabilities, nil
}

func (s *PostgresStorage) CheckHotelIDSql(ctx context.Context, hotelID string) (bool, error) {
	var exists bool
	err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM hotels WHERE id = $1)", hotelID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *PostgresStorage) UpdateRoomAvailability(ctx context.Context, ID string, availability bool) error {
	query := `
	    UPDATE rooms
	    SET availability = $1
	    WHERE id = $2
	`
	_, err := s.db.ExecContext(ctx, query, availability, ID)
	if err != nil {
		return fmt.Errorf("failed to update room availability: %v", err)
	}
	return nil
}

func (s *PostgresStorage) GetRoomIDSql(ctx context.Context, req *hotelpb.GetRoomIDRequest) (*hotelpb.GetRoomIDResponse, error) {
	query := `
        SELECT id
        FROM rooms
        WHERE hotel_id = $1 AND room_type = $2
    `

	var roomID string
	err := s.db.QueryRowContext(ctx, query, req.HotelID, req.RoomType).Scan(&roomID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("room not found for hotelID %s and roomType %s", req.HotelID, req.RoomType)
		}
		return nil, err
	}

	return &hotelpb.GetRoomIDResponse{
		RoomID: roomID,
	}, nil
}
