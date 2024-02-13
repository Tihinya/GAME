package engine

import (
	"testing"
)

func TestMaxHealthSystem(t *testing.T) {
	player := entityManager.CreateEntity("player")
	defer func() {
		DeleteAllEntityComponents(player)
	}()

	playerPosition := &PositionComponent{X: 10, Y: 5, Size: 1}
	playerHealth := &HealthComponent{CurrentHealth: 3, MaxHealth: 3}

	healthManager.AddComponent(player, playerHealth)
	positionManager.AddComponent(player, playerPosition)

	hc := healthManager.healths[player]

	hc.CurrentHealth += 1

	CallHealthSystem.Update(0.1)

	if hc.CurrentHealth != 3 {
		t.Fatalf("Expected player health to be 1, got %v", hc.CurrentHealth)
	}
	if positionManager.GetPosition(player) == nil {
		t.Fatalf("Expected player position to sbe nil, got %v", positionManager.GetPosition(player))
	}
}

func TestMinHealthSystem(t *testing.T) {
	player := entityManager.CreateEntity("player")
	defer func() {
		DeleteAllEntityComponents(player)
	}()

	playerPosition := &PositionComponent{X: 10, Y: 5, Size: 1}
	playerHealth := &HealthComponent{CurrentHealth: 1, MaxHealth: 3}

	healthManager.AddComponent(player, playerHealth)
	positionManager.AddComponent(player, playerPosition)

	hc := healthManager.healths[player]
	hc.CurrentHealth -= 1

	CallHealthSystem.Update(0.1)

	if hc.CurrentHealth != 0 {
		t.Fatalf("Expected player health to be 0, got %v", hc.CurrentHealth)
	}

	if playerPosition := positionManager.GetPosition(player); playerPosition != nil {
		t.Fatalf("Expected player position to sbe nil, got %v", playerPosition)
	}
}
