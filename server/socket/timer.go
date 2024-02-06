package socket

import (
	"bomberman-dom/models"
	"log"
	"time"
)

type State interface {
	startTimer(int)
	resetTimer()
}

type Timer struct {
	instance    *time.Ticker
	state       State
	currentTime int
	lobby       *Lobby
	C           chan models.LobbyState
	// description string
}

func (t *Timer) startTimer(countdown int) {
	t.state.startTimer(countdown)
}

func (t *Timer) resetTimer() {
	t.state.resetTimer()
}

func (t *Timer) changeState(state State) {
	t.state = state
}

type AwaitingPlayersState struct {
	timer *Timer
}

func newAwaitingPlayersState(timer *Timer) AwaitingPlayersState {
	return AwaitingPlayersState{
		timer,
	}
}

func (s AwaitingPlayersState) startTimer(countdown int) {
	s.timer.instance = time.NewTicker(time.Millisecond * 100)

	s.timer.C <- models.LobbyState{
		State: "awaiting_players_state",
	}

	for range s.timer.instance.C {
		playersAmount := s.timer.lobby.getAmountOfPlayers()
		if playersAmount > 1 && playersAmount < 4 {
			s.resetTimer()
			s.timer.changeState(newClosingLobbyState(s.timer))
			return
		} else if playersAmount == 4 {
			s.resetTimer()
			s.timer.changeState(newStartingGameState(s.timer))
			return
		}
	}
}

func (s AwaitingPlayersState) resetTimer() {
	close(s.timer.C)
	s.timer.instance.Stop()
}

type ClosingLobbyState struct {
	timer *Timer
}

func newClosingLobbyState(timer *Timer) ClosingLobbyState {
	return ClosingLobbyState{
		timer,
	}
}

func (s ClosingLobbyState) startTimer(countdown int) {
	s.timer.instance = time.NewTicker(time.Second * 1)
	s.timer.currentTime = 21

	for range s.timer.instance.C {
		playersAmount := s.timer.lobby.getAmountOfPlayers()

		if playersAmount < 2 {
			s.resetTimer()
			s.timer.changeState(newAwaitingPlayersState(s.timer))
			return
		} else if playersAmount == 4 {
			s.resetTimer()
			s.timer.changeState(newStartingGameState(s.timer))
			return
		} else if s.timer.currentTime == 0 {
			s.resetTimer()
			s.timer.changeState(newStartingGameState(s.timer))
			return
		}

		s.timer.currentTime--
		s.timer.C <- models.LobbyState{
			CurrentTime: s.timer.currentTime,
			State:       "closing_lobby_state",
		}
	}
}

func (s ClosingLobbyState) resetTimer() {
	s.timer.instance.Stop()
	s.timer.currentTime = 0
	close(s.timer.C)
}

type StartingGameState struct {
	timer *Timer
}

func newStartingGameState(timer *Timer) StartingGameState {
	return StartingGameState{
		timer,
	}
}

func (s StartingGameState) startTimer(countdown int) {
	s.timer.instance = time.NewTicker(time.Second * 1)
	s.timer.currentTime = 11

	for range s.timer.instance.C {
		playersAmount := s.timer.lobby.getAmountOfPlayers()

		if playersAmount < 2 {
			s.resetTimer()
			s.timer.changeState(newAwaitingPlayersState(s.timer))
			return
		} else if playersAmount == 4 {
			s.resetTimer()
			s.timer.changeState(newStartingGameState(s.timer))
			return
		} else if s.timer.currentTime == 0 {
			s.resetTimer()
			// s.timer.changeState(newStartingGameState(s.timer))
			// move to game
			s.timer.lobby.userList = make(OnlineList)
			log.Println("You won my cock!")
			return
		}

		s.timer.currentTime--
		s.timer.C <- models.LobbyState{
			CurrentTime: s.timer.currentTime,
			State:       "starting_game_state",
		}
	}
}

func (s StartingGameState) resetTimer() {
	s.timer.instance.Stop()
	s.timer.currentTime = 0
	close(s.timer.C)
}
