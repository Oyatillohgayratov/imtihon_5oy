package user

import (
	"booking-service/config"
	userpb "booking-service/protos/user"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CheckUserID(req string) (bool, error) {
	cfg := config.Load(".")
	port := fmt.Sprintf("%s:%s", cfg.UserServiceHost, cfg.UserServicePort)

	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return false, err
	}

	client := userpb.NewUserServiceClient(conn)

	res, err := client.CheckUserID(context.Background(), &userpb.CheckUserIDRequest{Id: req})
	if err != nil {
		return false, err
	}

	return res.Valid, nil
}

func GetUser(id string) (*userpb.GetUserResponse, error) {
	cfg := config.Load(".")
	port := fmt.Sprintf("%s:%s", cfg.UserServiceHost, cfg.UserServicePort)

	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return &userpb.GetUserResponse{}, err
	}

	client := userpb.NewUserServiceClient(conn)

	res, err := client.GetUser(context.Background(), &userpb.GetUserRequest{UserID: id})
	if err != nil {
		return &userpb.GetUserResponse{}, err
	}

	return res, nil
}
