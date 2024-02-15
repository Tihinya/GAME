package engine_test

import (
	"os"
	"testing"

	"bomberman-dom/socket"
)

func TestMain(m *testing.M) {
	socket.Instance = socket.NewManager()
	// engine.SetBroadcaster(socket.Instance.Lobby)

	code := m.Run()
	os.Exit(code)
}
