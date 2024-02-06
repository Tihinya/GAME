package socket

// TODO:
// 1. Login +
// 1.1 save names in a state. names unique +
// 1.2 kick user from active on disconection +
// 2. Chat +
// 2.0 add all users; add a user on registretion +
// 2.1 Send message +
// 2.2 Receive all message +
// 3. Timer				Need to be tested
// 3.1 create three states +
// 3.2 1 player in lobby no timer	+
// 3.3 if 2 or 3 players in lobby then 20 sec waiting timer, then 10 sec countdown to start game +
// 3.4 if 4 players start 10 sec countdown to start game	+
// 3.5 if other state req are met switch to that state

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
