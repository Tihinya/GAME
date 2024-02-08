package socket

import (
	"errors"
	"log"
	"net/http"
	"sync"

	"bomberman-dom/helpers"
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
	handlers    map[string]EventHandler
	Broadcaster helpers.Broadcaster
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
	m.handlers[GameEventInput] = GameInputHandler
}

func (m *Manager) routeEvent(event models.Event, c *Client) error {
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
		client.connection.Close()
		delete(m.clients, client)
	}
}

func (m *Manager) GetConnectedClient(username string) models.Event {
	var clientInfo models.ClientInfo

	for client := range m.clients {
		if client.username == username {
			clientInfo = models.ClientInfo{
				Username: username,
				Id:       client.id,
			}
			break
		}
	}

	return helpers.SerializeData(EventClientInfoMessage, clientInfo)
}

func (m *Manager) GetConnectedClients() models.Event {
	var onlineUserList models.ConnectedUserListEvent
	onlineUserList.List = make(map[int]string)

	for client := range m.clients {
		onlineUserList.List[client.id] = client.username
	}

	return helpers.SerializeData(EventOnlineUserList, onlineUserList)
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	// ws://localhost:8080/ws?username=exampleUser
	username := r.URL.Query().Get("username")
	if username == "" || m.usernameInClients(username) {
		helpers.ReturnMessageJSON(w,
			"Username is empty or already taken",
			http.StatusBadRequest, "Error")
		return
	}

	websocketUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		helpers.ReturnMessageJSON(w,
			"Could not upgrade to websocket connection, internal error",
			http.StatusInternalServerError, "Error")
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
