package service

import (
	"context"
	"fmt"
	"user-service/pkg/hash"
	"user-service/pkg/token"
	pb "user-service/protos/user"
	"user-service/storage/postgresql"
)

type UserServer struct {
	storage postgresql.Storage
	pb.UnimplementedUserServiceServer
}

func NewUserService(storage postgresql.Storage) *UserServer {
	return &UserServer{storage: storage}
}

func (s *UserServer) Register(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	hashpassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	userID, err := s.storage.InsertUser(ctx, &pb.RegisterUserRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: hashpassword,
	})

	if err != nil {
		return &pb.RegisterUserResponse{}, err
	}

	return &pb.RegisterUserResponse{
		UserID:   fmt.Sprintf("%d", userID),
		Username: req.Username,
		Email:    req.Email,
	}, nil
}

func (s *UserServer) Login(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	check, err := s.storage.LoginSql(ctx, req)
	if err != nil {
		return nil, err
	}
	if !check {
		return nil, err
	}

	gentoken, err := token.GenerateToken(req.Username)
	if err != nil {
		return nil, err
	}

	return &pb.LoginUserResponse{
		Token:     gentoken,
		ExpiresIn: int32(token.TokenTTL.Seconds()),
	}, nil
}

func (s *UserServer) CheckUserID(ctx context.Context, req *pb.CheckUserIDRequest) (*pb.CheckUserIDResponse, error) {
	valid, err := s.storage.CheckUserIDSql(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.CheckUserIDResponse{Valid: valid}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.storage.GetUserSql(ctx, req.UserID)
	if err == nil {
		return user, nil
	}
	return nil, err
}
