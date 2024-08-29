package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"user-service/config"
	connpostgresql "user-service/pkg/connPostgreSql"
	pb "user-service/protos/user"
	"user-service/service"
	"user-service/storage/postgresql"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load(".")
	db, err := connpostgresql.ConnectToDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	postgresStorage := postgresql.NewPostgresStorage(db)
	userService := service.NewUserService(postgresStorage)

	lis, err := net.Listen("tcp", ":"+cfg.UserServicePort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcInterceptor),
	)
	reflection.Register(s)

	pb.RegisterUserServiceServer(s, userService)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Printf("UserService server is running on port :%s", cfg.UserServicePort)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.GracefulStop()

	log.Println("Shutting down gracefully...")
	<-ctx.Done() 
	log.Println("Shutdown complete")

	// log.Printf("UserService server is running on port :%s", cfg.UserServicePort)
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("Failed to serve: %v", err)
	// }
}

func grpcInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	m, err := handler(ctx, req)
	if err != nil {
		log.Printf("RPC failed with error: %v", err)
		return nil, err
	}
	return m, nil
}