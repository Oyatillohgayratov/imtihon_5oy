package hotel

import (
	"booking-service/config"
	hotelpb "booking-service/protos/hotel"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CheckHotelID(req string) (bool, error) {
	cfg := config.Load(".")
	port := fmt.Sprintf("%s:%s", cfg.HotelServiceHost, cfg.HotelServicePort)

	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return false, err
	}

	client := hotelpb.NewHotelServiceClient(conn)

	res, err := client.CheckHotelID(context.Background(), &hotelpb.CheckHotelIDRequest{Id: req})
	if err != nil {
		return false, err
	}

	return res.Valid, nil
}

func GetRoomID(hotelid, roomtype string) (string, error) {
	cfg := config.Load(".")
	port := fmt.Sprintf("%s:%s", cfg.HotelServiceHost, cfg.HotelServicePort)

	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}

	client := hotelpb.NewHotelServiceClient(conn)

	res, err := client.GetRoomID(context.Background(), &hotelpb.GetRoomIDRequest{HotelID: hotelid, RoomType: roomtype})
	if err != nil {
		return "", err
	}

	return res.RoomID, nil
}

func Updateavailability(roomID string, update bool) (bool, error) {
	cfg := config.Load(".")
	port := fmt.Sprintf("%s:%s", cfg.HotelServiceHost, cfg.HotelServicePort)

	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return false, err
	}

	client := hotelpb.NewHotelServiceClient(conn)

	res, err := client.UdateRoomAvailability(context.Background(), &hotelpb.UdateRoomAvailabilityRequest{
		RoomID: roomID,
		Update: update,
	})
	if err != nil {
		return false, err
	}
	return res.Updated, nil
}
