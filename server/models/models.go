package models

import (
	"time"
)

//-- Socket Events --\\

type SendMessageEvent struct {
	Message  string    `json:"message"`
	SenderID int       `json:"sender_id"`
	SentTime time.Time `json:"sent_time"`
}

type ConnectedUserListEvent struct {
	List map[int]string `json:"list"`
}

type ReceivedMessage struct {
	Message    string    `json:"message"`
	SenderID   int       `json:"sender_id"`
	ReceiverID int       `json:"receiver_id"`
	Sent       time.Time `json:"sent_time"`
}

type ClientInfo struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
}

type GameState struct {
	State string `json:"game_state"` // "STARTED", "PAUSED", "ENDED"
}

type GameInput struct {
	PlayerID int             `json:"player_id"`
	Keys     map[string]bool `json:"keys"`
}

type GameBomb struct {
	PlayerID int    `json:"player_id"`
	Action   string `json:"action"` // "place", "detonate"
	X        int    `json:"x"`
	Y        int    `json:"y"`
}

type GameError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//-- Miscellaneous --\\

// HTTP JSON response
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
