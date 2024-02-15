package socket

import (
	"bomberman-dom/models"
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type (
	ClientList map[*Client]bool
	OnlineList map[string]*Client
	Client     struct {
		connection *websocket.Conn
		manager    *Manager
		egress     chan models.Event
		id         int // Unique identifier for the client
		username   string
		lobby      *Lobby
	}
)

var (
	pongWait       = 10 * time.Second
	pingInterval   = (pongWait * 9) / 10
	maxMessageSize = 512
)

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	manager.UserId++
	return &Client{
		connection: conn,
		manager:    manager,
		egress:     make(chan models.Event),
		id:         manager.UserId,
	}
}

func (c *Client) readMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()

	c.connection.SetReadLimit(int64(maxMessageSize))

	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		return
	}
	c.connection.SetPongHandler(c.pongHandler)

	for {
		_, payload, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break
		}
		var request models.Event
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("error marshalling message: %v", err)
			continue
		}

		if err := c.manager.routeEvent(request, c); err != nil {
			log.Println("Error handling Message: ", err)
		}
	}
}

func (c *Client) pongHandler(pongMsg string) error {
	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}

func (c *Client) writeMessages() {
	ticker := time.NewTicker(pingInterval)

	defer func() {
		ticker.Stop()
		c.manager.removeClient(c)
	}()

	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed: ", err)
				}
				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				log.Println(err)
				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println(err)
			}

		case <-ticker.C:
			if err := c.connection.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println("writemsg: ", err)
				return
			}
		}
	}
}
