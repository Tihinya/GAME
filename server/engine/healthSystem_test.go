package engine

import (
	"testing"
	"time"
)

func TestMaxHealthSystem(t *testing.T) {
	player := entityManager.CreateEntity()

	playerPosition := &PositionComponent{X: 10, Y: 5, Size: 1}
	playerHealth := &HealthComponent{CurrentHealth: 3, MaxHealth: 3}

	healthManager.AddComponent(player, playerHealth)
	positionManager.AddComponet(player, playerPosition)

	hc := healthManager.healths[player]

	hc.CurrentHealth += 1

	healthSystem.update(time.Now())

	if hc.CurrentHealth != 3 {
		t.Fatalf("Expected player health to be 1, got %v", hc.CurrentHealth)
	}
	if positionManager.postions[player] == nil {
		t.Fatalf("Expected player position to sbe nil, got %v", positionManager.postions[player])
	}
}

func TestMinHealthSystem(t *testing.T) {
	player := entityManager.CreateEntity()

	playerPosition := &PositionComponent{X: 10, Y: 5, Size: 1}
	playerHealth := &HealthComponent{CurrentHealth: 1, MaxHealth: 3}

	healthManager.AddComponent(player, playerHealth)
	positionManager.AddComponet(player, playerPosition)

	hc := healthManager.healths[player]
	hc.CurrentHealth -= 1

	healthSystem.update(time.Now())

	if hc.CurrentHealth != 0 {
		t.Fatalf("Expected player health to be 0, got %v", hc.CurrentHealth)
	}

	if positionManager.postions[player] != nil {
		t.Fatalf("Expected player position to sbe nil, got %v", positionManager.postions[player])
	}

}
