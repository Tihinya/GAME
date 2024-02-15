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

type MessageEvent struct {
	Name    string    `json:"name"`
	Time    time.Time `json:"send_time"`
	Message string    `json:"message"`
}

type ConnectedUserListEvent struct {
	List map[int]string `json:"list"`
}

type ClientInfo struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
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

type GamePlayer struct { // game_player_creation & game_player_position
	ClientId int     `json:"clientId"`
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
}

type GamePlayerHealth struct {
	ClientId int `json:"clientId"`
	Health   int `json:"health"`
}

//-- Miscellaneous --\\

// HTTP JSON response
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

type AddUsernameEvent struct {
	UserName string `json:"username"`
}

type CurrentUsers struct {
	UserList []string `json:"user_list"`
}

type ChangeState struct {
	State string `json:"state"`
}

type LobbyState struct {
	CurrentTime int    `json:"currentTime"`
	State       string `json:"state"`
}

type Position struct {
	X, Y float64
	Size float64
}

type GameStateTransmission struct {
	players    []Position
	bombs      []Position
	explosions []Position
	walls      []Position
	boxes      []Position
	powerups   []Position
}
