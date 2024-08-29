package service

import (
	"context"
	"fmt"

	"booking-service/internal/hotel"
	"booking-service/internal/notifaction"
	"booking-service/internal/user"
	bookingpb "booking-service/protos/booking"
	postgres "booking-service/storage/postgresql"
)

type BookingServer struct {
	storage postgres.Storage
	bookingpb.UnimplementedBookingServiceServer
}

func NewBookingServer(storage postgres.Storage) *BookingServer {
	return &BookingServer{
		storage: storage,
	}
}

func (s *BookingServer) CreateBooking(ctx context.Context, req *bookingpb.CreateBookingRequest) (*bookingpb.CreateBookingResponse, error) {
	r1, err := user.CheckUserID(req.UserID)
	if !r1 {
		if err != nil {
			return nil, fmt.Errorf("failed to check user ID: %v", err)
		}
		return nil, fmt.Errorf("invalid user ID")
	}
	r2, err := hotel.CheckHotelID(req.HotelID)
	if !r2 {
		if err != nil {
			return nil, fmt.Errorf("failed to check hotel ID: %v", err)
		}
		return nil, fmt.Errorf("invalid hotel ID")
	}
	roomid, err := hotel.GetRoomID(req.HotelID, req.RoomType)
	if err != nil {
		return nil, fmt.Errorf("failed to get room ID: %v", err)
	}

	bookingID, err := s.storage.InsertBooking(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create booking: %v", err)
	}

	r4, err := hotel.Updateavailability(roomid, false)
	if !r4 {
		if err != nil {
			return nil, fmt.Errorf("failed to update hotel availability: %v", err)
		}
		return nil, fmt.Errorf("failed to update hotel availability")
	}

	_, err = notifaction.SendEmail(req.UserID, "xona muvafaqiyatli bron qilindi")
	if err != nil {
		return nil, fmt.Errorf("failed to send notification: %v", err)
	}

	message := fmt.Sprintf("Booking confirmed! Booking ID: %d", bookingID)
	notifaction.BroadcastMessage(req.UserID, message)

	return &bookingpb.CreateBookingResponse{
		BookingID:    fmt.Sprintf("%d", bookingID),
		UserID:       req.UserID,
		HotelID:      req.HotelID,
		RoomType:     req.RoomType,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
		TotalAmount:  req.TotalAmount,
		Status:       "Confirmed",
	}, nil
}

func (s *BookingServer) GetBookingDetails(ctx context.Context, req *bookingpb.GetBookingDetailsRequest) (*bookingpb.GetBookingDetailsResponse, error) {
	booking, err := s.storage.GetBookingDetails(ctx, req.BookingID)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *BookingServer) UpdateBooking(ctx context.Context, req *bookingpb.UpdateBookingRequest) (*bookingpb.UpdateBookingResponse, error) {
	err := s.storage.UpdateBooking(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update booking: %v", err)
	}

	return &bookingpb.UpdateBookingResponse{
		BookingID:    req.BookingID,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
		TotalAmount:  req.TotalAmount,
		Status:       req.Status,
	}, nil
}

func (s *BookingServer) CancelBooking(ctx context.Context, req *bookingpb.CancelBookingRequest) (*bookingpb.CancelBookingResponse, error) {
	err := s.storage.CancelBooking(ctx, req.BookingID)
	if err != nil {
		return nil, fmt.Errorf("failed to cancel booking: %v", err)
	}
	r, err := s.GetBookingDetails(ctx, &bookingpb.GetBookingDetailsRequest{BookingID: req.BookingID})
	if err != nil {
		return nil, fmt.Errorf("failed to get booking details: %v", err)
	}
	roomid, err := hotel.GetRoomID(r.HotelID, r.RoomType)
	if err != nil {
		return nil, fmt.Errorf("failed to get room ID: %v", err)
	}

	r4, err := hotel.Updateavailability(roomid, true)
	if !r4 {
		if err != nil {
			return nil, fmt.Errorf("failed to update hotel availability: %v", err)
		}
		return nil, fmt.Errorf("failed to update hotel availability")
	}

	_, err = notifaction.SendEmail(r.UserID, "xona muvafaqiyatli chiqib ketildi")
	if err != nil {
		return nil, fmt.Errorf("failed to send notification: %v", err)
	}

	// message := fmt.Sprintf("Booking canceled. Booking ID: %s", req.BookingID)
	// notifaction.BroadcastMessage(r.UserID, message)

	return &bookingpb.CancelBookingResponse{
		Message:   "Booking successfully canceled",
		BookingID: req.BookingID,
	}, nil
}

func (s *BookingServer) ListUserBookings(ctx context.Context, req *bookingpb.ListUserBookingsRequest) (*bookingpb.ListUserBookingsResponse, error) {
	bookings, err := s.storage.ListUserBookings(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	_, err = notifaction.SendEmail(req.UserID, "xona muvafaqiyatli yangilandi qilindi")
	if err != nil {
		return nil, fmt.Errorf("failed to send notification: %v", err)
	}

	return &bookingpb.ListUserBookingsResponse{
		BookingList: bookings,
	}, nil
}
