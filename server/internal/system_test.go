package internal

import (
	"fmt"
	"testing"
	"time"

	"bomberman-dom/gameloop"
)

func TestMovmentSystem(t *testing.T) {
	player := entityManager.CreateEntity()
	box := entityManager.CreateEntity()

	playerPosition := &PositionComponent{X: 10, Y: 5, Size: 1}
	boxPosition := &PositionComponent{X: 10, Y: 12, Size: 1}

	playerMotion := &MotionComponent{Velocity: Vec2{X: 0, Y: 0}, Acceleration: Vec2{X: 0, Y: 0}} // Example motion values
	playerInput := &InputComponent{Input: map[string]bool{"down": true}}

	positionManager.AddComponent(player, playerPosition)
	positionManager.AddComponent(box, boxPosition)

	motionManager.AddComponent(player, playerMotion)
	inputManager.AddComponent(player, playerInput)

	for i := 0; i < 10; i++ {
		// inputSystem.update(time.Now())
		// motionSystem.update(time.Now())
	}

	fmt.Println(playerPosition)
}

func TestBomb(t *testing.T) {
	player := CreatePlayer()
	playerPosition := positionManager.positions[player]

	fmt.Println("STAGE 1: Testing bomb placing")

	bomb := CreateBomb(player)
	fmt.Printf("STAGE 1: Player placed a bomb at (X: %v, Y: %v, Size: %v)\n",
		playerPosition.X, playerPosition.Y, playerPosition.Size)

	bombPos, exists := positionManager.positions[bomb]
	if !exists {
		t.Fatalf("STAGE 1 FAILED: Bomb position not found in PositionManager")
	}

	if bombPos.X != playerPosition.X || bombPos.Y != playerPosition.Y {
		t.Errorf("STAGE 1 FAILED: Bomb id %v placed incorrectly. Expected (%v, %v), got (%v, %v)",
			bomb.Id, playerPosition.X, playerPosition.Y, bombPos.X, bombPos.Y)
	}

	fmt.Printf("STAGE 1: Bomb id %v succcessfully found at (X: %v, Y: %v, Size: %v)\n",
		bomb.Id, bombPos.X, bombPos.Y, bombPos.Size)
	fmt.Println("STAGE 1: SUCCESS")

	fmt.Println("STAGE 2: Testing bomb explosion")
	fps := 10
	loop := gameloop.New(time.Duration(fps), func(dt float64) {
		explosionSystem.update(dt)
		fmt.Printf("STAGE 2: Time until bomb %v detonation: %v ms\n", bomb.Id, (time.Until(timerManager.timers[bomb].Time)))
	})
	fmt.Printf("STAGE 2: Gameloop started at a tickrate of %v, running for 600 milliseconds\n", fps)

	loop.Start()
	time.Sleep(time.Millisecond * 530)
	fmt.Println("STAGE 2: Checking if bomb hasn't detonated yet")

	for en, ec := range explosionManager.explosions {
		fmt.Println(*en, *ec)
	}
}
