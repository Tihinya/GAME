package socket_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"bomberman-dom/models"
	"bomberman-dom/socket"

	"github.com/gorilla/websocket"
)

func TestWebSocket(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(socket.NewManager().ServeWS))
	defer server.Close()

	url := "ws" + server.URL[4:]

	user1 := connectSocket(t, url, "user1")
	defer user1.Close()

	user2 := connectSocket(t, url, "user2")
	defer user2.Close()

	message := "Hello World!"
	sendMessage(t, user1, message, 1)
	// Test receiving chat message events
	receiveMessage(t, user2, "chat_message", message)

	user3 := connectSocket(t, url, "user3")
	defer user3.Close()

	// Test receiving client info event
	receiveMessage(t, user3, "client_info", "user3", "3")

	user4 := connectSocket(t, url, "user4")
	defer user4.Close()
	receiveMessage(t, user4, "online_list", "4")
}

func connectSocket(t *testing.T, baseURL string, username string) *websocket.Conn {
	dialer := websocket.DefaultDialer
	header := http.Header{}

	// Append the username query parameter to the URL
	fullURL := fmt.Sprintf("%s/ws?username=%s", baseURL, username)

	fmt.Printf("TRYING: Testing websocket connection with user: %v\n", username)

	// Connect user
	user, _, err := dialer.Dial(fullURL, header)
	if err != nil {
		t.Fatalf(
			"ERROR: Could not open a ws connection for user: %v\n %v\n",
			username, err)
	}

	fmt.Printf("SUCCESS: successfully connected an user account to websockets\n")

	return user
}

func sendMessage(t *testing.T, connection *websocket.Conn, message string, senderId int) {
	sendData := &models.MessageEvent{
		Message: message,
		Time:    time.Now(),
	}

	jsonData, err := json.Marshal(sendData)
	if err != nil {
		t.Fatal("ERROR: Failed marshalling to JSON", err)
	}

	var payload json.RawMessage = jsonData

	eventData := models.Event{
		Type:    "send_message",
		Payload: payload,
	}

	eventJSON, err := json.Marshal(eventData)
	if err != nil {
		t.Fatal("ERROR: Failed marshalling to JSON", err)
	}

	err = connection.WriteMessage(websocket.TextMessage, eventJSON)
	if err != nil {
		t.Fatalf("Could not send message from user 1 to user 2 %v", err)
	}
	fmt.Printf("SUCCESS: Sent message '%v'\n", message)
}

// receieveType types are ["chat_message", "online_list", "client_info"]
// if receive type is "chat_message", parameter for data is a string message
// if receive_type is "client_info", parameter for data is a string "username, id"
// if receive type is "online_list", parameter for data is the connected user's string id
func receiveMessage(t *testing.T, connection *websocket.Conn, receiveType string, data ...any) {
	timeout := time.After(1 * time.Second)
	messageChan := make(chan []byte)
	errChan := make(chan error)

	// Start a goroutine to read messages
	go func() {
		for {
			_, msg, err := connection.ReadMessage()
			if err != nil {
				errChan <- err
				return
			}
			messageChan <- msg
		}
	}()

	fmt.Println("TRYING: receiving a message...")

	for {
		select {
		case <-timeout:
			fmt.Printf("TIMEOUT: no expected message received within 1 second\n")
			return

		case err := <-errChan:
			t.Fatalf("ERROR: Could not read message: %v", err)

		case receivedPayload := <-messageChan:
			var event models.Event
			if err := json.Unmarshal(receivedPayload, &event); err != nil {
				t.Fatalf("ERROR: bad payload in request: %v", err)
			}

			switch event.Type {
			case socket.EventReceiveMessage:
				if receiveType == "chat_message" {
					var message models.MessageEvent
					if err := json.Unmarshal(event.Payload, &message); err != nil {
						t.Fatalf("ERROR: unmarshaling message: %v\n", err)
					}
					if data[0] == message.Message {
						fmt.Printf("SUCCESS: Received message %v\n", message.Message)
						return
					}
				}

			case socket.EventOnlineUserList:
				if receiveType == "online_list" {
					var onlineUsersList models.ConnectedUserListEvent
					if err := json.Unmarshal(event.Payload, &onlineUsersList); err != nil {
						t.Fatalf(
							"ERROR: unmarshaling online users list: %v\n",
							err)
					}
					if !contains(onlineUsersList.List, data[0]) {
						t.Fatalf(
							"ERROR: Expected user id %v to be in online user list: %v\n",
							data, onlineUsersList.List)
					}
					fmt.Printf(
						"SUCCESS: Found user id %v in online users list %v\n",
						data, onlineUsersList.List)
					return
				}

			case socket.EventClientInfoMessage:
				if receiveType == "client_info" {
					var clientInfo models.ClientInfo
					if err := json.Unmarshal(event.Payload, &clientInfo); err != nil {
						t.Fatalf("ERROR: unmarshaling client info: %v\n", err)
					}

					if strconv.Itoa(clientInfo.Id) != data[1] {
						fmt.Printf(
							"ERROR: did not find an user with id %v and username %v in clientInfo",
							data[1], data[0])
					}

					fmt.Printf(
						"SUCCESS: found user with id %v and username %v in clientInfo: %v\n",
						data[1], data[0], clientInfo)
					return
				}

			default:
				t.Fatalf("ERROR: Unexpected message type: %v", event.Type)
			}
		}
	}
}

func contains(arr map[int]string, userid any) bool {
	for id := range arr {
		if userid == strconv.Itoa(id) {
			return true
		}
	}
	return false
}
