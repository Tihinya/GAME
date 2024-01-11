package socket

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"bomberman-dom/models"
)

const (
	EventSendMessage    = "send_message"
	EventOnlineUserList = "online_users_list"
	EventReceiveMessage = "receive_message"
	ClientInfoMessage   = "client_info"
)

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client) error

func SendMessageHandler(event Event, c *Client) error {
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

	var outgoingEvent Event = Event{
		Type:    EventReceiveMessage,
		Payload: data,
	}

	for client := range c.manager.clients {
		client.egress <- outgoingEvent
	}

	return nil
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
