package socket

import (
	"bomberman-dom/engine"
	"bomberman-dom/models"
	"encoding/json"
	"fmt"
)

// TODO:
// 1. Game engine
// 1.1 Reduce amount of sent events
// 1.2 Compress explosionsystem broacasts to 1 event per tick, rather than one event per explosion per tick
// 1.3 check for bugs and game logic integrity 8==================D <- good size ( Y . Y ) smol
// 2. Front end
// 2.1 Recieve events
// 2.2 Display assets
// 2.3 Send events: awsd <space/>
// 3. Show game end page ( Optional )

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
