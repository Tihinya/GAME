package managers

import (
	"testing"
	"time"
)

func TestValidComponentManager(t *testing.T) {
	entityManager := NewEntityManager()
	entity := entityManager.CreateEntity()
	positionManager := NewPositionManager()
	timerManager := NewTimerManager()
	spriteManager := NewSpriteManager()

	positionComponent := &PositionComponent{X: 10, Y: 20}
	positionManager.AddComponet(entity, positionComponent)

	spriteComponent := &SpriteComponent{Texture: "player.png"}
	spriteManager.AddComponent(entity, spriteComponent)

	timerComponent := &TimerComponent{Time: time.Now()}
	timerManager.AddComponet(entity, timerComponent)

	if retrievedComponent, ok := positionManager.postions[entity]; !ok || retrievedComponent != positionComponent {
		t.Fatal("AddComponent to PositionManager failed")
	}
	if retrievedComponent, ok := spriteManager.sprites[entity]; !ok || retrievedComponent != spriteComponent {
		t.Fatal("AddComponent to SpriteManager failed")
	}
	if retrievedComponent, ok := timerManager.timers[entity]; !ok || retrievedComponent != timerComponent {
		t.Fatal("AddComponet to TimerManager failed")
	}
}
