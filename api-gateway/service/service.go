package service

import (
	"api-gateway/config"
	bookingpb "api-gateway/protos/booking"
	hotelpb "api-gateway/protos/hotel"
	userpb "api-gateway/protos/user"
	"log"

	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IClients interface {
	User() userpb.UserServiceClient
	Booking() bookingpb.BookingServiceClient
	Hotel() hotelpb.HotelServiceClient
}

type ServiceManager struct {
	userService    userpb.UserServiceClient
	bookingService bookingpb.BookingServiceClient
	hotelService   hotelpb.HotelServiceClient
}

func New(cfg config.Config) (*ServiceManager, error) {
	log.Printf("UserServiceHost: %s, UserServicePort: %s", cfg.UserServiceHost, cfg.UserServicePort)

	connUser, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial user service: %v", err)
	}

	connBooking, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.BookingServiceHost, cfg.BookingServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial booking service: %v", err)
	}

	connHotel, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", cfg.HotelServiceHost, cfg.HotelServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!= nil {
        return nil, fmt.Errorf("failed to dial hotel service: %v", err)
    }

	return &ServiceManager{
		userService: userpb.NewUserServiceClient(connUser),
		bookingService: bookingpb.NewBookingServiceClient(connBooking),
        hotelService: hotelpb.NewHotelServiceClient(connHotel),
	}, nil
}

func (s *ServiceManager) User() userpb.UserServiceClient {
	return s.userService
}

func (s *ServiceManager) Booking() bookingpb.BookingServiceClient {
	return s.bookingService
}

func (s *ServiceManager) Hotel() hotelpb.HotelServiceClient {
	return s.hotelService
}
