package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"hotel-service/config"
	connpostgresql "hotel-service/pkg/connPostgreSql"
	hotelpb "hotel-service/protos/hotel"
	"hotel-service/service"
	"hotel-service/storage/postgresql"

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
	storage := postgresql.NewPostgresStorage(db)

	hotelservice := service.NewHotelServer(storage)

	lis, err := net.Listen("tcp", ":"+cfg.HotelServicePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcInterceptor),
	)
	reflection.Register(s)

	hotelpb.RegisterHotelServiceServer(s, hotelservice)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		fmt.Printf("Hotel Service listening on port :%s\n", cfg.HotelServicePort)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.GracefulStop()

	fmt.Println("Shutting down gracefully...")
	<-ctx.Done() 
	fmt.Println("Shutdown complete")

	// fmt.Printf("Hotel Service listening on port :%s\n", cfg.HotelServicePort)
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
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
