package main

import (
	"context"
	"log"
	"net"
	"notifaction-service/config"
	notifactionpb "notifaction-service/protos/notifaction"
	"notifaction-service/service"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load(".")

	// Initialize Gin router and WebSocket route
	// r := gin.Default()
	// r.GET("/ws", handleWebSocket)

	// Initialize Notification Service
	notificationServer := service.NewNotificationServer()

	// Set up gRPC server
	lis, err := net.Listen("tcp", ":"+cfg.NotificationServicePort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpcInterceptor),
	)
	reflection.Register(grpcServer)

	notifactionpb.RegisterNotificationServiceServer(grpcServer, notificationServer)

	// Start HTTP server for WebSocket
	// go func() {
	// 	log.Printf("Starting HTTP server for WebSocket on port 7070\n")
	// 	if err := r.Run(":7070"); err != nil {
	// 		log.Fatalf("Failed to run HTTP server: %v", err)
	// 	}
	// }()

	// Start gRPC server in a goroutine
	go func() {
		log.Printf("Starting GRPC server on port :%s\n", cfg.NotificationServicePort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve GRPC server: %v", err)
		}
	}()

	// Handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	// Graceful shutdown of gRPC server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	grpcServer.GracefulStop()

	log.Println("Shutting down gracefully...")
	<-ctx.Done()
	log.Println("Shutdown complete")
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

// func handleWebSocket(c *gin.Context) {
// 	// Extract WebSocket connection from the request
// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		log.Printf("Error while upgrading connection: %v", err)
// 		return
// 	}
// 	defer conn.Close()

// 	// Handle WebSocket connection
// 	for {
// 		messageType, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Printf("Error while reading message: %v", err)
// 			break
// 		}

// 		// Handle received message
// 		log.Printf("Received message: %s", msg)
// 		err = conn.WriteMessage(messageType, msg)
// 		if err != nil {
// 			log.Printf("Error while writing message: %v", err)
// 			break
// 		}
// 	}
// }

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }
