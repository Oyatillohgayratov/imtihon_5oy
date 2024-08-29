package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"booking-service/config"
	connpostgresql "booking-service/pkg/connPostgreSql"
	bookingpb "booking-service/protos/booking"
	"booking-service/service"
	postgres "booking-service/storage/postgresql"

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
	storage := postgres.NewPostgresStorage(db)

	bookingservice := service.NewBookingServer(storage)

	lis, err := net.Listen("tcp", ":"+cfg.BookingServicePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcInterceptor),
	)
	reflection.Register(s)

	bookingpb.RegisterBookingServiceServer(s, bookingservice)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		fmt.Printf("Hotel & Booking Services listening on port :%s\n", cfg.BookingServicePort)
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

	// fmt.Printf("Hotel & Booking Services listening on port :%s\n", cfg.BookingServicePort)
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
