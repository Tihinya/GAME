package engine_test

import (
	"os"
	"testing"

	"bomberman-dom/engine"
	"bomberman-dom/socket"
)

func TestMain(m *testing.M) {
	socket.Instance = socket.NewManager()
	engine.SetBroadcaster(socket.Instance)

	code := m.Run()
	os.Exit(code)
}
