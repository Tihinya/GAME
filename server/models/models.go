package models

import "time"

//-- Socket Events --\\

type MessageEvent struct {
	Message  string    `json:"message"`
	SentTime time.Time `json:"sent_time"`
	Name     string    `json:"name"`
}

type ConnectedUserListEvent struct {
	List map[int]string `json:"list"`
}

type ClientInfo struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
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
