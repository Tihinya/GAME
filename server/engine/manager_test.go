package engine

import (
	"testing"
	"time"
)

func TestValidComponentManager(t *testing.T) {
	timerManager := NewTimerManager()
	spriteManager := NewSpriteManager()

	entity := entityManager.CreateEntity("player")
	defer func() {
		DeleteAllEntityComponents(entity)
	}()

	positionComponent := &PositionComponent{X: 10, Y: 20}
	positionManager.AddComponent(entity, positionComponent)

	spriteComponent := &SpriteComponent{Texture: "player.png"}
	spriteManager.AddComponent(entity, spriteComponent)

	timerComponent := &TimerComponent{Time: time.Now()}
	timerManager.AddComponent(entity, timerComponent)

	if retrievedComponent, ok := positionManager.positions[entity]; !ok || retrievedComponent != positionComponent {
		t.Fatal("AddComponent to PositionManager failed")
	}
	if retrievedComponent, ok := spriteManager.sprites[entity]; !ok || retrievedComponent != spriteComponent {
		t.Fatal("AddComponent to SpriteManager failed")
	}
	if retrievedComponent, ok := timerManager.timers[entity]; !ok || retrievedComponent != timerComponent {
		t.Fatal("AddComponent to TimerManager failed")
	}
}
