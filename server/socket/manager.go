package socket

import (
	"encoding/json"
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
)

type Manager struct {
	clients ClientList
	sync.RWMutex
	handlers    map[string]EventHandler
	UserId      int
	Lobby       *Lobby
	Broadcaster helpers.Broadcaster
}

func NewManager() *Manager {
	lobby := &Lobby{userList: make(OnlineList)}
	timer := &Timer{lobby: lobby, C: make(chan models.LobbyState)}
	timer.state = newAwaitingPlayersState(timer)
	lobby.timer = timer

	m := &Manager{
		clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
		Lobby:    lobby,
	}
	lobby.manager = m
	m.setupEventHandlers()

	return m
}

func (m *Manager) setupEventHandlers() {
	m.handlers[EventSendMessage] = MessageHandler
	m.handlers[GameEventGameState] = GameStateHandler
	m.handlers[EventLoginHandler] = UsernameHandler
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
	m.clients[client] = true
}

func (m *Manager) removeClient(client *Client) {
	if client.lobby != nil {
		client.lobby.removePlayer(client.username)
		client.lobby = nil
	}

	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
	}

	broadcastOnlineUserList(m)
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
		if client.username != "" {
			onlineUserList.List[client.id] = client.username
		}
	}

	return helpers.SerializeData(EventOnlineUserList, onlineUserList)
}

func (m *Manager) ServeWS(w http.ResponseWriter, r *http.Request) {

	websocketUpgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := websocketUpgrader.Upgrade(w, r, nil)

	if err != nil {

		log.Println(err)
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

func SerializeData(EventType string, data ...any) models.Event {
	if len(data) < 1 {
		return models.Event{}
	}

	jsonData, err := json.Marshal(data[0])
	if err != nil {
		log.Printf("failed to marshal online user list: %v", err)
	}

	var outgoingEvent models.Event
	outgoingEvent.Payload = jsonData
	outgoingEvent.Type = EventType

	return outgoingEvent
}
