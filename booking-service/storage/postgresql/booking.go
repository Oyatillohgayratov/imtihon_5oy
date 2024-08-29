package postgres

import (
	bookingpb "booking-service/protos/booking"
	"context"
	"database/sql"
	"fmt"
)

type Storage interface {
	InsertBooking(ctx context.Context, booking *bookingpb.CreateBookingRequest) (int, error)
	GetBookingDetails(ctx context.Context, bookingID string) (*bookingpb.GetBookingDetailsResponse, error)
	UpdateBooking(ctx context.Context, booking *bookingpb.UpdateBookingRequest) error
	CancelBooking(ctx context.Context, bookingID string) error
	ListUserBookings(ctx context.Context, userID string) ([]*bookingpb.Booking, error)

}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{
		db: db,
	}
}
func (s *PostgresStorage) InsertBooking(ctx context.Context, booking *bookingpb.CreateBookingRequest) (int, error) {
	var bookingID int
	err := s.db.QueryRowContext(ctx,
		`INSERT INTO bookings (user_id, hotel_id, room_type, check_in_date, check_out_date, total_amount, status) 
		 VALUES ($1, $2, $3, $4, $5, $6, $7) 
		 RETURNING id`,
		booking.UserID, booking.HotelID, booking.RoomType, booking.CheckInDate, booking.CheckOutDate, booking.TotalAmount, "Confirmed").Scan(&bookingID)

	if err != nil {
		return 0, err
	}

	return bookingID, nil
}

func (s *PostgresStorage) GetBookingDetails(ctx context.Context, bookingID string) (*bookingpb.GetBookingDetailsResponse, error) {
	query := `
		SELECT id, user_id, hotel_id, room_type, check_in_date, check_out_date, total_amount, status
		FROM bookings
		WHERE id = $1
	`
	row := s.db.QueryRowContext(ctx, query, bookingID)

	var booking bookingpb.GetBookingDetailsResponse
	if err := row.Scan(&booking.BookingID, &booking.UserID, &booking.HotelID, &booking.RoomType, &booking.CheckInDate, &booking.CheckOutDate, &booking.TotalAmount, &booking.Status); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("booking with ID %s not found", bookingID)
		}
		return nil, err
	}

	return &booking, nil
}

func (s *PostgresStorage) UpdateBooking(ctx context.Context, booking *bookingpb.UpdateBookingRequest) error {
	_, err := s.db.ExecContext(ctx,
		`UPDATE bookings
		 SET check_in_date = $2, check_out_date = $3, total_amount = $4, status = $5
		 WHERE id = $1`,
		booking.BookingID, booking.CheckInDate, booking.CheckOutDate, booking.TotalAmount, booking.Status)

	return err
}

func (s *PostgresStorage) CancelBooking(ctx context.Context, bookingID string) error {
	_, err := s.db.ExecContext(ctx,
		`UPDATE bookings
		 SET status = 'Cancelled'
		 WHERE id = $1`,
		bookingID)
	return err
}

func (s *PostgresStorage) ListUserBookings(ctx context.Context, userID string) ([]*bookingpb.Booking, error) {
	rows, err := s.db.QueryContext(ctx,
		`SELECT id, hotel_id, room_type, check_in_date, check_out_date, total_amount, status
		 FROM bookings
		 WHERE user_id = $1`,
		userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []*bookingpb.Booking
	for rows.Next() {
		var booking bookingpb.Booking
		if err := rows.Scan(&booking.BookingID, &booking.HotelID, &booking.RoomType, &booking.CheckInDate, &booking.CheckOutDate, &booking.TotalAmount, &booking.Status); err != nil {
			return nil, err
		}
		bookings = append(bookings, &booking)
	}

	return bookings, nil
}

func (s *PostgresStorage) BookinsRooms()  {

}

func (s *PostgresStorage) CheckRooms()  {
	
}
