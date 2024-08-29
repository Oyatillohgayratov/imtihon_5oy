package service

import (
	"context"
	"fmt"
	"net/smtp"

	pb "notifaction-service/protos/notifaction"
)

type NotificationServiceServer struct {
	pb.UnimplementedNotificationServiceServer
	// wsConn *websocket.Conn
}

func NewNotificationServer() *NotificationServiceServer {
	ns := &NotificationServiceServer{}
	// ns.connectWebSocket()
	return ns
}

// func (ns *NotificationServiceServer) connectWebSocket() {
// 	url := "ws://web-socket-service:7070/ws"

// 	var err error
// 	for {
// 		ns.wsConn, _, err = websocket.DefaultDialer.Dial(url, nil)
// 		if err != nil {
// 			log.Printf("Failed to connect to WebSocket server: %v. Retrying...", err)
// 			time.Sleep(5 * time.Second)
// 			continue
// 		}
// 		log.Println("Successfully connected to WebSocket server")
// 		break
// 	}
// }

func (ns *NotificationServiceServer) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	err := sendEmail(req.To, req.Subject, req.Body)
	if err != nil {
		return &pb.SendEmailResponse{Status: "Failed"}, err
	}
	return &pb.SendEmailResponse{Status: "Success"}, nil
}

// func (ns *NotificationServiceServer) NotifyWebSocket(ctx context.Context, req *pb.NotifyWebSocketRequest) (*pb.NotifyWebSocketResponse, error) {
// 	err := ns.SendWebSocketMessage(req.UserID, req.Message)
// 	if err != nil {
// 		log.Printf("Error sending WebSocket message: %v", err)
// 		return &pb.NotifyWebSocketResponse{Status: "Failed"}, err
// 	}

// 	return &pb.NotifyWebSocketResponse{Status: "Sent"}, nil
// }

// SendWebSocketMessage sends a message to the WebSocket server
// func (ns *NotificationServiceServer) SendWebSocketMessage(userID, message string) error {
// 	if ns.wsConn == nil {
// 		ns.connectWebSocket() // Reconnect if connection is lost
// 	}

// 	msg := map[string]string{
// 		"userID":  userID,
// 		"message": message,
// 	}
// 	data, err := json.Marshal(msg)
// 	if err != nil {
// 		return fmt.Errorf("failed to marshal message: %v", err)
// 	}

// 	err = ns.wsConn.WriteMessage(websocket.TextMessage, data)
// 	if err != nil {
// 		log.Printf("Failed to send message, reconnecting WebSocket: %v", err)
// 		ns.connectWebSocket() // Attempt to reconnect
// 		return err
// 	}

// 	log.Printf("Sent message to WebSocket: %s\n", message)
// 	return nil
// }

func sendEmail(to, subject, body string) error {
	from := "dilshoddilmurodov112@gmail.com" // Use environment variables
	pass := "xmxu rdhp pmdf pezk"            // Use environment variables
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, pass, smtpHost)
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}
