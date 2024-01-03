package statemanager

type GameState interface {
	Init()
	Update(dt float64)
	Cleanup()
}

type StateManager struct {
	currentState GameState
	states       map[string]GameState
}

func (sm *StateManager) SetState(name string) {
	if sm.currentState != nil {
		sm.currentState.Cleanup()
	}
	sm.currentState = sm.states[name]
	sm.currentState.Init()
}

func (sm *StateManager) Update(dt float64) {
	if sm.currentState != nil {
		sm.currentState.Update(dt)
	}
}
