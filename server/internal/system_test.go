package internal

import (
	"fmt"
	"testing"
	"time"
)

func TestMovmentSystem(t *testing.T) {

	player := entityManager.CreateEntity()
	box := entityManager.CreateEntity()

	playerPosition := &PositionComponent{X: 10, Y: 5, Size: 1}
	boxPosition := &PositionComponent{X: 10, Y: 12, Size: 1}

	playerMotion := &MotionComponent{Velocity: Vec2{X: 0, Y: 0}, Acceleration: Vec2{X: 0, Y: 0}} // Example motion values
	playerInput := &InputComponent{Input: map[string]bool{"down": true}}

	positionManager.AddComponet(player, playerPosition)
	positionManager.AddComponet(box, boxPosition)

	motionManager.AddComponent(player, playerMotion)
	inputManager.AddComponet(player, playerInput)

	for i := 0; i < 10; i++ {
		inputSystem.update(time.Now())
		motionSystem.update(time.Now())

	}

	fmt.Println(playerPosition)
}
