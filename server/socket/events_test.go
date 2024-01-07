package socket_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"bomberman-dom/models"
	"bomberman-dom/socket"

	"github.com/gorilla/websocket"
)

func TestWebSocketChat(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(socket.NewManager().ServeWS))
	defer server.Close()

	url := "ws" + server.URL[4:]

	user := connectSocket(t, url, "user1")
	defer user.Close()

	user1 := connectSocket(t, url, "user2")
	defer user1.Close()

	message := "Hello World!"
	sendMessage(t, user1, message, 1) // send message from user id to user id
}

func sendMessage(t *testing.T, connection *websocket.Conn, message string, senderId int) {
	sendData := &models.SendMessageEvent{
		ReceiveMessageEvent: models.ReceiveMessageEvent{
			Message: message,
		},
		SenderID: senderId,
		SentTime: time.Now(),
	}

	jsonData, err := json.Marshal(sendData)
	if err != nil {
		t.Fatal("Failed marshalling to JSON", err)
	}

	var payload json.RawMessage = jsonData

	eventData := socket.Event{
		Type:    "send_message",
		Payload: payload,
	}

	eventJSON, err := json.Marshal(eventData)
	if err != nil {
		t.Fatal("Failed marshalling to JSON", err)
	}

	err = connection.WriteMessage(websocket.TextMessage, eventJSON)
	if err != nil {
		t.Fatalf("Could not send message from user 1 to user 2 %v", err)
	}
	fmt.Printf("Sent message\n")
}
