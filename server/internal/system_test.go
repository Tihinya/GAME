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

func TestBombPlacingAndDetonation(t *testing.T) {
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

	fmt.Printf("STAGE 2: Checking if bomb hasn't detonated in 100ms (detonation is %v)\n", fuseTime)
	initialCheckTime := time.Millisecond * 100
	time.Sleep(initialCheckTime)

	bombTimer := timerManager.timers[bomb]
	if bombTimer != nil && time.Now().Before(bombTimer.Time) {
		fmt.Printf("STAGE 2: Bomb %v has not exploded in 100ms\n", bomb.Id)
	} else {
		t.Errorf("STAGE 2: Bomb %v has exploded", bomb.Id)
	}

	fmt.Printf("STAGE 2: Checking if bomb hasn't detonated in 530ms (detonation is %v)\n", fuseTime)
	time.Sleep(fuseTime - initialCheckTime)

	if bombTimer != nil && time.Now().Before(bombTimer.Time) {
		t.Errorf("STAGE 2 FAILED: Bomb %v should have exploded", bomb.Id)
	} else {
		fmt.Printf("STAGE 2: Bomb %v has exploded after %v\n", bomb.Id, fuseTime)
	}

	spreadPositions := getExplosionSpreadPositions(bomb)
	for _, pos := range spreadPositions {
		//if !explosionExistsAt(pos) {
		//	t.Errorf("STAGE 2 FAILED: No explosion found at (X: %v, Y: %v)", pos.X, pos.Y)
		//}
		fmt.Println(pos)
	}
}

func getExplosionSpreadPositions(e *Entity) []PositionComponent {
	var positions []PositionComponent
	bc := bombManager.bombs[e]
	bombPos := positionManager.positions[e]

	// Spread the explosion in each direction and collect the positions
	for i := 1; i <= bc.BlastRadius; i++ {
		positions = append(positions, PositionComponent{X: bombPos.X + float64(i), Y: bombPos.Y}) // Right
		positions = append(positions, PositionComponent{X: bombPos.X - float64(i), Y: bombPos.Y}) // Left
		positions = append(positions, PositionComponent{X: bombPos.X, Y: bombPos.Y + float64(i)}) // Up
		positions = append(positions, PositionComponent{X: bombPos.X, Y: bombPos.Y - float64(i)}) // Down
	}

	return positions
}
