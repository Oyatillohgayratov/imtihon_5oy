package notifaction

import (
	"booking-service/config"
	"booking-service/internal/user"
	notifactionpb "booking-service/protos/notifaction"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func SendEmail(userId, message string) (string, error) {
	cfg := config.Load(".")

	port := fmt.Sprintf("%s:%s", cfg.NotificationServiceHost, cfg.NotificationServicePort)
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", fmt.Errorf("failed to connect to user service: %v", err)
	}
	defer conn.Close()

	client := notifactionpb.NewNotificationServiceClient(conn)

	r, err := user.GetUser(userId)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %v", err)
	}
	if r == nil {
		return "", fmt.Errorf("user not found")
	}

	res, err := client.SendEmail(context.Background(), &notifactionpb.SendEmailRequest{
		To:      r.Email,
		Subject: "HOTEL",
		Body:    "asalom aleykum " + r.Username + "\n" + message,
	})
	if err != nil {
		return "", fmt.Errorf("failed to send email: %v", err)
	}

	return res.Status, nil
}

func BroadcastMessage(userid, message string) error {
	cfg := config.Load(".")

	port := fmt.Sprintf("%s:%s", cfg.NotificationServiceHost, cfg.NotificationServicePort)
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to user service: %v", err)
	}
	defer conn.Close()

	client := notifactionpb.NewNotificationServiceClient(conn)

	_, err = client.NotifyWebSocket(context.Background(), &notifactionpb.NotifyWebSocketRequest{})
	if err != nil {
		return fmt.Errorf("failed to notify websocket: %v", err)
	}
	fmt.Printf("Sent message to websocket: %s\n", message)
	return nil
}
