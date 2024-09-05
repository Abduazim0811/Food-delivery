package notificationservice

import (
	"context"
	"log"
	"notification-service/internal/http/handler"
	"notification-service/notificationproto"
)

type Service struct {
	notificationproto.UnimplementedNotificationServer
	W handler.HandlerWebSocket
}

func (s *Service) AddUser(ctx context.Context, req *notificationproto.AddUserRequest) (*notificationproto.EmailResponse, error) {

	conn, err := s.W.GetUserConnection(req.UserId)
	if err != nil {
		log.Println("Error retrieving user connection:", err)
		return nil, err
	}
	if conn == nil {
		return &notificationproto.EmailResponse{
			Message: "User connection not found",
		}, nil
	}

	if err := s.W.AddUser(req.UserId, conn); err != nil {
		log.Println("Error adding user:", err)
		return nil, err
	}

	return &notificationproto.EmailResponse{
		Message: "User added successfully",
	}, nil
}
