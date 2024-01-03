package statemanager

import (
	"fmt"
	"testing"
)

type mockGameState struct {
	inited, updated, cleaned bool
}

func (m *mockGameState) Init() {
	// For main menu it could be intializing menu, loading menu assets...
	m.inited = true
}

func (m *mockGameState) Update(dt float64) {
	// Could have check function that waits for 'Start game' button press
	// And change state manager state to statemanager.SetState("gameplay")
	m.updated = true
}

func (m *mockGameState) Cleanup() {
	// Cleanup resources of main menu
	m.cleaned = true
}

func TestStateManager(t *testing.T) {
	stateManager := StateManager{states: make(map[string]GameState)}
	testState := &mockGameState{}
	stateManager.states["test"] = testState

	// e.g. mainmenu gets initalized and all other mockGameStates get changed
	// given state, like "main-menu" or "test" below
	stateManager.SetState("test")
	if !testState.inited {
		t.Errorf("Expected Init to be called")
	}
	fmt.Println("State manager state set successfully")

	// Updates current game state with time delta (for a time-based game)
	stateManager.Update(0.016)
	if !testState.updated {
		t.Errorf("Expected Update to be called")
	}
	fmt.Println("State manager updated successfully")

	stateManager.SetState("test") // Re-setting to call Cleanup
	if !testState.cleaned {
		t.Errorf("Expected Cleanup to be called")
	}
	fmt.Println("State manager cleaned up and re-set successfully")
}
