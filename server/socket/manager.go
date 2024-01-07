package socket

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"bomberman-dom/models"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	ErrEventNotSupported = errors.New("this event type is not supported")
	Instance             *Manager
	UserIdCounter        int
)

type Manager struct {
	clients ClientList
	sync.RWMutex
	handlers map[string]EventHandler
}

func NewManager() *Manager {
	m := &Manager{
		clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
	}
	m.setupEventHandlers()

	return m
}

func (m *Manager) setupEventHandlers() {
	m.handlers[EventSendMessage] = SendMessageHandler
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return ErrEventNotSupported
	}
}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()
	m.clients[client] = true
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		m.Lock()
		defer m.Unlock()
		client.connection.Close()
		delete(m.clients, client)
	}
}

func (m *Manager) GetConnectedClient(username string) Event {
	return Event{}
}

func (m *Manager) GetConnectedClients() Event {
	var onlineUserList models.ConnectedUserListEvent

	for client := range m.clients {
		onlineUserList.List[client.id] = client.username
	}

	return SerializeData(EventOnlineUserList, onlineUserList)
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	// ws://localhost:8080/ws?username=exampleUser
	username := r.URL.Query().Get("username")
	if username == "" || m.usernameInClients(username) {
		log.Printf("Failed to set username, username is empty or already taken")
		return
	}

	websocketUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Create New Client
	client := NewClient(conn, m, username, idCounter())
	m.addClient(client)

	go client.readMessages()
	go client.writeMessages()

	broadcastClientInfo(m, username)
	broadcastOnlineUserList(m)
}

func (m *Manager) usernameInClients(username string) bool {
	for client := range m.clients {
		if client.username == username {
			return true
		}
	}
	return false
}

func (m *Manager) GetClientById(id int) *Client {
	for client := range m.clients {
		if client.id == id {
			return client
		}
	}
	return nil
}

func (m *Manager) GetClientByUsername(username string) *Client {
	for client := range m.clients {
		if client.username == username {
			return client
		}
	}
	return nil
}

func idCounter() int {
	UserIdCounter++
	return UserIdCounter
}

func SerializeData(EventType string, data ...any) Event {
	fmt.Println(data)
	if len(data) == 1 {
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Printf("failed to marshal online user list: %v", err)
		}

		var outgoingEvent Event
		outgoingEvent.Payload = jsonData
		outgoingEvent.Type = EventOnlineUserList

		return outgoingEvent
	}
	return Event{}
}
