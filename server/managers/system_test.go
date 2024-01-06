package managers

import (
	"testing"
)

var spriteManager = NewSpriteManager()
var positionManager = NewPositionManager()
var motionManager = NewMotionManager() // Add motion manager
var entityManager = NewEntityManager()

func TestMovmentSystem(t *testing.T) {

	player := entityManager.CreateEntity()
	box := entityManager.CreateEntity()

	playerSprite := &SpriteComponent{Texture: "player.png"}
	boxSprite := &SpriteComponent{Texture: "enemy.png"}

	playerPosition := &PositionComponent{X: 10, Y: 5}
	boxPosition := &PositionComponent{X: 11, Y: 5}

	playerMotion := &MotionComponent{Velocity: Vec2{X: 1, Y: 0}, Acceleration: Vec2{X: 0.1, Y: 0}} // Example motion values

	positionManager.AddComponet(player, playerPosition)
	positionManager.AddComponet(box, boxPosition)

	spriteManager.AddComponent(player, playerSprite)
	spriteManager.AddComponent(box, boxSprite)

	motionManager.AddComponent(player, playerMotion)

	systems := &SystemManagers{
		PositionManager: positionManager,
		MotionManager:   motionManager,
	}

	entityManager.update(1, systems)

	expectedPositionX := 11.0
	expectedPositionY := 5.0
	expectedMotionVelocityX := 1.1
	expectedMotionVelocityY := 0.0

	if player.getPosition(systems).X != expectedPositionX || player.getPosition(systems).Y != expectedPositionY {
		t.Errorf("Unexpected player position. Got (%.2f, %.2f), expected (%.2f, %.2f)", player.getPosition(systems).X, player.getPosition(systems).Y, expectedPositionX, expectedPositionY)
	}

	if player.getMotion(systems).Velocity.X != expectedMotionVelocityX || player.getMotion(systems).Velocity.Y != expectedMotionVelocityY {
		t.Errorf("Unexpected player motion. Got (%.2f, %.2f), expected (%.2f, %.2f)", player.getMotion(systems).Velocity.X, player.getMotion(systems).Velocity.Y, expectedMotionVelocityX, expectedMotionVelocityY)
	}

}

func TestInputSystem(t *testing.T) {
	// Add your input system tests here.
}
