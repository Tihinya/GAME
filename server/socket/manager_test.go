package socket_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"bomberman-dom/socket"

	"github.com/gorilla/websocket"
)

func TestSecureWebSocketConnection(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(socket.NewManager().ServeWS))
	defer server.Close()

	url := "ws" + server.URL[4:]

	user := connectSocket(t, url, "user1")
	defer user.Close()
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
		t.Fatalf("ERROR: Could not open a ws connection for user: %v\n %v\n", username, err)
	}

	fmt.Printf("SUCCESS: successfully connected an user account to websockets\n")

	return user
}
