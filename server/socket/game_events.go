package socket

import (
	"bomberman-dom/models"
	"encoding/json"
	"fmt"
	"strings"
)

// TODO:
// 1. Login +
// 1.1 save names in a state. names unique +
// 1.2 kick user from active on disconection
// 2. Chat
// 2.0 add all users; add a user on registretion
// 2.1 Send message
// 2.2 Receive all message
// 3. Timer

func GameStateHandler(event Event, c *Client) error {
	return nil
}

func GameMoveHandler(event Event, c *Client) error {
	return nil
}

func GameBombHandler(event Event, c *Client) error {
	return nil
}

func GameObstacleHandler(event Event, c *Client) error {
	return nil
}

func GamePowerupHandler(event Event, c *Client) error {
	return nil
}

func GameNotificationHandler(event Event, c *Client) error {
	return nil
}

func UsernameHandler(event Event, c *Client) error {
	var adduserEvent models.AddUsernameEvent
	if err := json.Unmarshal(event.Payload, &adduserEvent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}
	adduserEvent.UserName = strings.TrimSpace(adduserEvent.UserName)

	if adduserEvent.UserName == "" || c.manager.usernameInClients(adduserEvent.UserName) {
		c.egress <- SerializeData(GameEventGameState, models.Response{Status: "Индус", Message: "ПОШЛИ НАХУЙ"})
		return fmt.Errorf("username is empty or username already taken")
	}

	if len(c.manager.clients) > 4 {
		c.egress <- SerializeData(GameEventGameState, models.Response{Status: "Индус", Message: "СЕЛ НАХУЙ"})
		return fmt.Errorf("too many users")
	}

	c.username = adduserEvent.UserName

	c.egress <- SerializeData(GameEventGameState, models.ChangeState{State: "lobby"})
	broadcastClientInfo(c.manager, c)
	broadcastOnlineUserList(c.manager)
	return nil
}
