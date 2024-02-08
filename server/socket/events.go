package socket

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"bomberman-dom/models"
)

const (
	EventSendMessage        = "send_message"         // Event for sending messages
	EventOnlineUserList     = "online_users_list"    // Event for receiving connected user list
	EventReceiveMessage     = "receive_message"      // Event for receiving messages
	EventClientInfoMessage  = "client_info"          // Displays username, id on user connect
	GameEventError          = "game_error"           // For backend error logs
	GameEventInput          = "game_input"           // Input - up, down, left, right, place
	GameEventGameState      = "game_state"           // State - lobby, start, end
	GameEventBomb           = "game_bomb"            // Bomb - place, explode
	GameEventObstacle       = "game_obstacle"        // Obstacles - boxes, powerups
	GameEventPowerup        = "game_powerup"         // Powerup - pickup
	GameEventExplosion      = "game_event"           // Explosion - appear, disappear
	GameEventPlayerMotion   = "game_player_position" // Player motion - position
	GameEventPlayerHealth   = "game_player_health"   // Player health - health
	GameEventPlayerCreation = "game_player_creation" // Player creation - X, Y
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event models.Event, c *Client) error

func SendMessageHandler(event models.Event, c *Client) error {
	var chatEvent models.SendMessageEvent
	if err := json.Unmarshal(event.Payload, &chatEvent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	chatEvent.Message = strings.TrimSpace(chatEvent.Message)
	if chatEvent.Message == "" {
		return nil
	}

	var broadMessage models.SendMessageEvent

	broadMessage.SentTime = time.Now()
	broadMessage.Message = chatEvent.Message
	broadMessage.SenderID = c.id

	data, err := json.Marshal(broadMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	var outgoingEvent models.Event = models.Event{
		Type:    EventReceiveMessage,
		Payload: data,
	}

	for client := range c.manager.clients {
		client.egress <- outgoingEvent
	}

	return nil
}

func (m *Manager) BroadcastClient(event models.Event, clientId int) {
	m.Lock()
	defer m.Unlock()

	user := m.GetClientById(clientId)
	for client := range m.clients {
		if client.username == user.username {
			client.egress <- event
		}
	}
}

func (m *Manager) BroadcastAllClients(event models.Event) {
	m.Lock()
	defer m.Unlock()

	for client := range m.clients {
		client.egress <- event
	}
}

func broadcastClientInfo(m *Manager, username string) {
	m.Lock()
	defer m.Unlock()
	clientInfoEvent := m.GetConnectedClient(username)
	for client := range m.clients {
		if client.username == username {
			client.egress <- clientInfoEvent
		}
	}
}

func broadcastOnlineUserList(m *Manager) {
	m.Lock()
	defer m.Unlock()
	onlineUsersListEvent := m.GetConnectedClients()
	for client := range m.clients {
		client.egress <- onlineUsersListEvent
	}
}
