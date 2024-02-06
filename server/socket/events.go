package socket

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"bomberman-dom/models"
)

const (
	EventSendMessage       = "send_message" // Event for sending messages
	EventAmaBoy            = "ama_boy_next_door"
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

func UsernameHandler(event Event, c *Client) error {
	var adduserEvent models.AddUsernameEvent
	if err := json.Unmarshal(event.Payload, &adduserEvent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}
	adduserEvent.UserName = strings.TrimSpace(adduserEvent.UserName)

	if adduserEvent.UserName == "" || c.manager.Lobby.isUsernameExists(adduserEvent.UserName) {
		c.egress <- SerializeData(GameEventGameState, models.Response{Status: "Индус", Message: "ПОШЛИ НАХУЙ"})
		return fmt.Errorf("username is empty or username already taken")
	}

	if c.manager.Lobby.getAmountOfPlayers() > 4 {
		c.egress <- SerializeData(GameEventGameState, models.Response{Status: "Индус", Message: "СЕЛ НАХУЙ"})
		return fmt.Errorf("too many players")
	}

	c.username = adduserEvent.UserName

	c.manager.Lobby.addPlayer(c)

	c.egress <- SerializeData(GameEventGameState, models.ChangeState{State: "lobby"})
	broadcastClientInfo(c.manager, c)
	broadcastOnlineUserList(c.manager)
	return nil
}

func broadcastClientInfo(m *Manager, client *Client) {
	client.egress <- SerializeData(EventClientInfoMessage, models.ClientInfo{
		Username: client.username,
		Id:       client.id,
	})
}

func broadcastOnlineUserList(m *Manager) {
	onlineUsersListEvent := SerializeData(EventOnlineUserList, m.Lobby.userList)

	for client := range m.clients {
		client.egress <- onlineUsersListEvent
	}
}
