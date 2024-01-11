package models

import "time"

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

//-- Miscellaneous --\\

// HTTP JSON response
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
