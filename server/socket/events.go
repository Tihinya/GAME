package socket

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"bomberman-dom/models"
)

const (
	EventSendMessage       = "send_message"      // Event for sending messages
	EventOnlineUserList    = "online_users_list" // Event for receiving connected user list
	EventReceiveMessage    = "receive_message"   // Event for receiving messages
	EventClientInfoMessage = "client_info"       // Displays username, id on user connect
	GameEventNotification  = "game_notification" // For backend error logs (maybe?)
	GameEventMovePlayer    = "game_move"         // Move - up, down, left, right
	GameEventGameState     = "game_state"        // State - lobby, start, end
	GameEventBomb          = "game_bomb"         // Bomb - place, explode
	GameEventObstacle      = "game_obstacle"     // Obstacles - boxes, powerups
	GameEventPowerup       = "game_powerup"      // Powerup - pickup
	EventLoginHandler      = "register_user"     // Register user
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client) error

func MessageHandler(event Event, c *Client) error {
	var message models.MessageEvent
	if err := json.Unmarshal(event.Payload, &message); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	message.Message = strings.TrimSpace(message.Message)
	if message.Message == "" {
		return nil
	}

	message.SentTime = time.Now()

	// connection must have a name to send message
	if c.username == "" {
		return nil
	}

	message.Name = c.username

	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	for client := range c.manager.clients {
		client.egress <- Event{
			Type:    EventReceiveMessage,
			Payload: data,
		}
	}

	return nil
}

func broadcastClientInfo(m *Manager, client *Client) {
	client.egress <- SerializeData(EventClientInfoMessage, models.ClientInfo{
		Username: client.username,
		Id:       client.id,
	})
}

func broadcastOnlineUserList(m *Manager) {
	onlineUsersListEvent := m.GetConnectedClients()
	log.Println(onlineUsersListEvent)

	for client := range m.clients {
		client.egress <- onlineUsersListEvent
	}
}
