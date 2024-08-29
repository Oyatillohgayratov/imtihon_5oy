package service

import (
	"context"
	hotelpb "hotel-service/protos/hotel"
	"hotel-service/storage/postgresql"
)

type HotelServer struct {
	storage postgresql.Storage
	hotelpb.UnimplementedHotelServiceServer
}

func NewHotelServer(storage postgresql.Storage) *HotelServer {
	return &HotelServer{storage: storage}
}

func (s *HotelServer) GetHotels(ctx context.Context, req *hotelpb.GetHotelsRequest) (*hotelpb.GetHotelsResponse, error) {
	hotels, err := s.storage.ListHotels(ctx)
	if err != nil {
		return nil, err
	}

	return &hotelpb.GetHotelsResponse{Hotels: hotels}, nil
}

func (s *HotelServer) AddHotel(ctx context.Context, req *hotelpb.Hotel) (*hotelpb.AddHotelResponse, error) {
	hotelID, err := s.storage.InsertHotel(ctx, req)
	if err != nil {
		return nil, err
	}

	return &hotelpb.AddHotelResponse{
		HotelID: int32(hotelID),
	}, nil
}

func (s *HotelServer) GetHotelDetails(ctx context.Context, req *hotelpb.GetHotelDetailsRequest) (*hotelpb.GetHotelDetailsResponse, error) {
	hotelID := req.GetHotelID()
	hotel, err := s.storage.GetHotelDetails(ctx, hotelID)
	if err != nil {
		return nil, err
	}

	return &hotelpb.GetHotelDetailsResponse{Hotel: hotel}, nil
}

func (s *HotelServer) CheckRoomAvailability(ctx context.Context, req *hotelpb.CheckRoomAvailabilityRequest) (*hotelpb.CheckRoomAvailabilityResponse, error) {
	hotelID := req.GetHotelID()
	roomAvailabilities, err := s.storage.CheckRoomAvailability(ctx, hotelID)
	if err != nil {
		return nil, err
	}

	return &hotelpb.CheckRoomAvailabilityResponse{RoomAvailabilities: roomAvailabilities}, nil
}

func (s *HotelServer) CheckHotelID(ctx context.Context, req *hotelpb.CheckHotelIDRequest) (*hotelpb.CheckHotelIDResponse, error) {
	valid, err := s.storage.CheckHotelIDSql(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &hotelpb.CheckHotelIDResponse{Valid: valid}, nil
}

func (s *HotelServer) UdateRoomAvailability(ctx context.Context, req *hotelpb.UdateRoomAvailabilityRequest) (*hotelpb.UdateRoomAvailabilityResponse, error) {
	err := s.storage.UpdateRoomAvailability(ctx, req.RoomID, req.Update)
	if err != nil {
		return &hotelpb.UdateRoomAvailabilityResponse{Updated: false}, err
	}
	return &hotelpb.UdateRoomAvailabilityResponse{Updated: true}, nil
}

func (s *HotelServer) GetRoomID(ctx context.Context, req *hotelpb.GetRoomIDRequest) (*hotelpb.GetRoomIDResponse, error) {
	res, err := s.storage.GetRoomIDSql(ctx, req)
	if err!= nil {
        return nil, err
    }
	return res, nil
}
