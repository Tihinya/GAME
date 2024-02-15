package socket

import (
	"bomberman-dom/engine"
	"bomberman-dom/models"
	"encoding/json"
	"fmt"
)

func GameStateHandler(event models.Event, c *Client) error {
	return nil
}

func GameInputHandler(event models.Event, c *Client) error {
	var gameInput models.GameInput
	if err := json.Unmarshal(event.Payload, &gameInput); err != nil {
		return fmt.Errorf("GameInputHandler - bad payload in request: %v", err)
	}

	engine.HandleInput(gameInput, c.id)

	return nil
}
