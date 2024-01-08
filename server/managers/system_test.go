package managers

import (
	"fmt"
	"testing"
)

var positionManager = NewPositionManager()
var motionManager = NewMotionManager()
var entityManager = NewEntityManager()
var inputManager = NewInputManager()
var collectionManager = NewCollisionManager()
var powerUpManager = NewPowerUpManager()
var damageManager = NewDamageManager()
var healthManager = NewHealthManager()

func TestMovmentSystem(t *testing.T) {

	player := entityManager.CreateEntity()
	box := entityManager.CreateEntity()

	playerPosition := &PositionComponent{X: 10, Y: 5, Size: 1}
	boxPosition := &PositionComponent{X: 10, Y: 6, Size: 1}

	playerMotion := &MotionComponent{Velocity: Vec2{X: 0, Y: 0}, Acceleration: Vec2{X: 0, Y: 0}} // Example motion values
	playerInput := &InputComponent{Input: map[string]bool{"down": true}}

	positionManager.AddComponet(player, playerPosition)
	positionManager.AddComponet(box, boxPosition)

	motionManager.AddComponent(player, playerMotion)
	inputManager.AddComponet(player, playerInput)

	systems := &SystemManagers{
		PositionManager:  positionManager,
		MotionManager:    motionManager,
		InputManager:     inputManager,
		CollisionManager: collectionManager,
		PowerUpManager:   powerUpManager,
		DamageManager:    damageManager,
		HealthManager:    healthManager,
	}

	entityManager.update(1, systems)

	// expectedPositionX := 11.0
	// expectedPositionY := 5.0
	// expectedMotionVelocityX := 1.1
	// expectedMotionVelocityY := 0.0

	fmt.Println(player.getPosition(systems))

	// if player.getPosition(systems).X != expectedPositionX || player.getPosition(systems).Y != expectedPositionY {
	// 	t.Errorf("Unexpected player position. Got (%.2f, %.2f), expected (%.2f, %.2f)", player.getPosition(systems).X, player.getPosition(systems).Y, expectedPositionX, expectedPositionY)
	// }

	// if player.getMotion(systems).Velocity.X != expectedMotionVelocityX || player.getMotion(systems).Velocity.Y != expectedMotionVelocityY {
	// 	t.Errorf("Unexpected player motion. Got (%.2f, %.2f), expected (%.2f, %.2f)", player.getMotion(systems).Velocity.X, player.getMotion(systems).Velocity.Y, expectedMotionVelocityX, expectedMotionVelocityY)
	// }

}
