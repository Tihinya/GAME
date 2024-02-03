package socket

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

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
)

type Message struct {
	Name    string
	Time    time.Time
	Message string
}

type Manager struct {
	clients ClientList
	sync.RWMutex
	handlers map[string]EventHandler
	UserId   int
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
	m.handlers[EventSendMessage] = MessageHandler
	m.handlers[GameEventNotification] = GameNotificationHandler
	m.handlers[GameEventMovePlayer] = GameMoveHandler
	m.handlers[GameEventGameState] = GameStateHandler
	m.handlers[GameEventBomb] = GameBombHandler
	m.handlers[GameEventObstacle] = GameObstacleHandler
	m.handlers[GameEventPowerup] = GamePowerupHandler
	m.handlers[EventLoginHandler] = UsernameHandler
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
		client.connection.Close()
		delete(m.clients, client)
	}

	broadcastOnlineUserList(m)
}

func (m *Manager) GetConnectedClient(username string) Event {
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

	return SerializeData(EventClientInfoMessage, clientInfo)
}

func (m *Manager) GetConnectedClients() Event {
	var onlineUserList models.ConnectedUserListEvent
	onlineUserList.List = make(map[int]string)

	for client := range m.clients {
		if client.username != "" {
			onlineUserList.List[client.id] = client.username
		}
	}

	return SerializeData(EventOnlineUserList, onlineUserList)
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {
	// TODO: return error with websocket and close connection. This part wont work

	websocketUpgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := websocketUpgrader.Upgrade(w, r, nil)

	if err != nil {

		log.Println(err)
		// helpers.ReturnMessageJSON(w,
		// 	"Could not upgrade to websocket connection, internal error",
		// 	http.StatusInternalServerError, "Error")
		return
	}

	// Create New Client
	client := NewClient(conn, m)
	m.addClient(client)

	go client.readMessages()
	go client.writeMessages()

	// TODO: move broadcasts to lobby connection(where you parse username)
	//
	// broadcastClientInfo(m, username)
	// broadcastOnlineUserList(m)
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

func SerializeData(EventType string, data ...any) Event {
	if len(data) == 1 {
		jsonData, err := json.Marshal(data[0])
		if err != nil {
			log.Printf("failed to marshal online user list: %v", err)
		}

		var outgoingEvent Event
		outgoingEvent.Payload = jsonData
		outgoingEvent.Type = EventType

		return outgoingEvent
	}
	return Event{}
}
