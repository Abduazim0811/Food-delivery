package storage

import (
	"context"
	"notification-service/internal/http/handler"
	notificationservice "notification-service/notification_service"

	"sync"

	"github.com/gorilla/websocket"
)

func NewWebSocket() *handler.HandlerWebSocket {

	ctx := context.Background()
	return &handler.HandlerWebSocket{
		Map:   make(map[string]*websocket.Conn),
		Mutex: &sync.Mutex{},
		Ctx:   ctx,
	}
}

func NewService() *notificationservice.Service {
	a := NewWebSocket()
	return &notificationservice.Service{W: *a}
}
