package socket

import (
	"encoding/json"
	"fmt"

	"bomberman-dom/models"
)

func GameInputHandler(event models.Event, c *Client) error {
	var gameInput models.GameInput
	if err := json.Unmarshal(event.Payload, &gameInput); err != nil {
		return fmt.Errorf("GameInputHandler - bad payload in request: %v", err)
	}

	return nil
}
