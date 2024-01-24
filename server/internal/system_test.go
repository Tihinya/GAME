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
	spreadPositions := getExplosionSpreadPositions(bomb)
	fmt.Printf("STAGE 1: Player placed a bomb at (X: %v, Y: %v, Size: %v)\n",
		playerPosition.X, playerPosition.Y, playerPosition.Size)

	bombPos := positionManager.GetPosition(bomb)
	if bombPos == nil {
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
	})
	fmt.Printf("STAGE 2: Gameloop started at a tickrate of %v, running for 600 milliseconds\n", fps)

	loop.Start()

	fmt.Printf("STAGE 2: Checking if bomb hasn't detonated in 100ms (detonation is %v)\n", fuseTime)
	initialCheckTime := time.Millisecond * 100
	time.Sleep(initialCheckTime)

	bombTimer := timerManager.GetTimer(bomb)
	if bombTimer != nil && time.Now().Before(bombTimer.Time) {
		fmt.Printf("STAGE 2: Bomb %v has not exploded (after %v)\n", bomb.Id, initialCheckTime)
	} else {
		t.Errorf("STAGE 2: Bomb %v has exploded", bomb.Id)
	}

	fmt.Printf("STAGE 2: Checking if bomb hasn't detonated in 530ms (detonation is %v)\n", fuseTime)
	time.Sleep(fuseTime - initialCheckTime)

	if bombTimer != nil && time.Now().Before(bombTimer.Time) {
		t.Errorf("STAGE 2 FAILED: Bomb %v should have exploded", bomb.Id)
	} else {
		fmt.Printf("STAGE 2: Bomb %v has exploded (after %v)\n", bomb.Id, fuseTime)
	}

	fmt.Println("STAGE 2: Testing for bomb explosion positions")

	for _, pos := range spreadPositions {
		if !explosionExistsAt(pos) {
			t.Errorf("STAGE 2 FAILED: No explosion found at (X: %v, Y: %v)", pos.X, pos.Y)
		}
		fmt.Printf("Explosion exists at (X: %v, Y: %v)\n", pos.X, pos.Y)
	}
	fmt.Println("STAGE 2: Checking if explosions disappear prematurely (after 50ms)")
	time.Sleep(explodeTime - (100 * time.Millisecond))

	for _, pos := range spreadPositions {
		if !explosionExistsAt(pos) {
			t.Errorf("STAGE 2 FAILED: No explosion found at (X: %v, Y: %v)", pos.X, pos.Y)
		}
	}

	fmt.Println("STAGE 2: Previous explosions exist at the same positions, checking if explosions disappear (after 100 ms)")
	time.Sleep(150 * time.Millisecond)
	for _, pos := range spreadPositions {
		if explosionExistsAt(pos) {
			t.Errorf("STAGE 2 ERROR: Explosions found at previous positions")
			return
		}
	}
	fmt.Printf("STAGE 2: No explosions found at previous positions\n")
	fmt.Printf("STAGE 2: SUCCESS\n")
	loop.Stop()
}

func TestBoxExplosion(t *testing.T) {
	fps := 30
	loop := gameloop.New(time.Duration(fps), func(dt float64) {
		explosionSystem.update(dt)
	})
	loop.Start()
	fmt.Printf("Gameloop started at a tickrate of %v, running for 600 milliseconds\n", fps)

	player := CreatePlayer() // X: 10 Y: 5

	bomb := CreateBomb(player)

	box := CreateBox(10, 4) // Box one unit above player

	fmt.Printf("Created player id %v, bomb id %v and box id %v\n", player.Id, bomb.Id, box.Id)
	fmt.Printf("Checking if box exists in given position\n")

	entityManager.mutex.RLock()
	for _, e := range entityManager.entities {
		if boxManager.GetBox(e) != nil {
			pc := positionManager.GetPosition(e)
			fmt.Printf("Box exists in (X: %v, Y: %v)\n", pc.X, pc.Y)
		}
	}
	entityManager.mutex.RUnlock()

	fmt.Printf("Waiting for %v for bomb to explode\n", explodeTime+fuseTime)
	time.Sleep(fuseTime + (explodeTime / 2))

	fmt.Printf("Checking if box was hit by explosion\n")
	entityManager.mutex.RLock()
	for _, e := range entityManager.entities {
		if boxManager.GetBox(e) != nil {
			pc := positionManager.GetPosition(e)
			if explosionExistsAt(*pc) {
				fmt.Printf("Found explosion at box's location\n")
			} else {
				fmt.Printf("ERROR: did not find explosion at location X: (%v, Y: %v)\n", pc.X, pc.Y)
			}
		}
	}
	entityManager.mutex.RUnlock()

	fmt.Printf("Checking if box was destroyed by explosion\n")
	time.Sleep(time.Millisecond * 200)

	if boxManager.GetBox(box) != nil {
		fmt.Printf("ERROR: Box is still alive >:(\n")
	}
	fmt.Printf("SUCCESS: box has been assassinated\n")
}

func getExplosionSpreadPositions(e *Entity) []PositionComponent {
	var positions []PositionComponent
	bc := bombManager.bombs[e]
	bombPos := positionManager.GetPosition(e)

	positions = append(positions, *bombPos)

	// Spread the explosion in each direction and collect the positions
	for i := 1; i < bc.BlastRadius; i++ {
		positions = append(positions, PositionComponent{
			X: bombPos.X + float64(i), Y: bombPos.Y, Size: 1,
		}) // Right
		positions = append(positions, PositionComponent{
			X: bombPos.X - float64(i), Y: bombPos.Y, Size: 1,
		}) // Left
		positions = append(positions, PositionComponent{
			X: bombPos.X, Y: bombPos.Y + float64(i), Size: 1,
		}) // Up
		positions = append(positions, PositionComponent{
			X: bombPos.X, Y: bombPos.Y - float64(i), Size: 1,
		}) // Down
	}

	return positions
}

func explosionExistsAt(pos PositionComponent) bool {
	entityManager.mutex.RLock()
	defer entityManager.mutex.RUnlock()
	for _, e := range entityManager.entities {
		if explosionManager.GetExplosion(e) != nil {
			pc := positionManager.GetPosition(e)
			if pc.X == pos.X && pc.Y == pos.Y {
				return true
			}
		}
	}
	return false
}
