package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/twmb/franz-go/pkg/kgo"
)

type HandlerWebSocket struct {
	Map   map[string]*websocket.Conn
	Mutex *sync.Mutex
	Ctx   context.Context
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (u *HandlerWebSocket) HandleWebSocket(c *gin.Context) {
	fmt.Println("WebSocket is working")
	userID := c.Request.Header.Get("id")
	fmt.Println(userID)
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user ID"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket Upgrade error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Current connections map:", u.Map)

	kafkaReader, err := kgo.NewClient(
		kgo.SeedBrokers("broker:29092"),
		kgo.ConsumeTopics("order_updates"),
	)
	if err != nil {
		log.Println("Kafka client creation error:", err)
		return
	}
	defer kafkaReader.Close()

	for {
		fetches := kafkaReader.PollFetches(c.Request.Context())
		if fetches.IsClientClosed() {
			break
		}
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			message := record.Value
			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println("Error writing message to WebSocket:", err)
				return
			}
			fmt.Println("Message sent to WebSocket")
		}
	}
}

func (u *HandlerWebSocket) AddUser(userID string, conn *websocket.Conn) error {
	u.Mutex.Lock()
	defer u.Mutex.Unlock()
	if _, exists := u.Map[userID]; exists {
		log.Printf("User %s is already connected", userID)
		return errors.New("user already exists")
	}
	u.Map[userID] = conn
	log.Printf("User %s added to the map", userID)
	return nil
}

func (u *HandlerWebSocket) GetUserConnection(userID string) (*websocket.Conn, error) {
	u.Mutex.Lock()
	defer u.Mutex.Unlock()

	conn, exists := u.Map[userID]
	if !exists {
		return nil, fmt.Errorf("connection for user %s not found", userID)
	}
	return conn, nil
}
