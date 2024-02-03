package models

import (
	"encoding/json"
	"time"
)

//-- Socket Events --\\

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

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

type GameState struct { // game_state
	State string `json:"game_state"` // "STARTED", "PAUSED", "ENDED"
}

type GameInput struct { // game_input
	Keys map[string]bool `json:"keys"`
}

type GameBomb struct { // game_bomb
	Action string  `json:"action"` // create, delete
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
}

type GameExplosion struct { // game_explosion
	Action string  `json:"action"` // create, delete
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
}

type GameError struct { // game_error
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type GameObstacle struct { // game_obstacle
	Type   string  `json:"type"`   // box, wall, powerup
	Action string  `json:"action"` // create, delete
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
}

type GamePowerup struct { // game_powerup
	Type   int     `json:"type"`   // speed(1), bomb(2), health(3), explosion(4)
	Action string  `json:"action"` // create, delete
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
}

type GamePlayer struct { // game_player
	ClientId int     `json:"clientId"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}

//-- Miscellaneous --\\

// HTTP JSON response
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
