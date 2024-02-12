package engine

import (
	"testing"
)

func TestSpeedPowerUpSystem(t *testing.T) {
	player := entityManager.CreateEntity("player")
	powerUp := entityManager.CreateEntity("powerup")
	defer func() {
		DeleteAllEntityComponents(player)
		DeleteAllEntityComponents(powerUp)
	}()

	playerPosition := &PositionComponent{X: 10, Y: 5, Size: 1}
	SpeedPowerUpPosition := &PositionComponent{X: 10, Y: 6, Size: 1}
	positionManager.AddComponent(player, playerPosition)
	positionManager.AddComponent(powerUp, SpeedPowerUpPosition)

	playerCollision := &CollisionComponent{Enabled: false}
	PowerUpCollision := &CollisionComponent{Enabled: true}
	collisionManager.AddComponent(player, playerCollision)
	collisionManager.AddComponent(powerUp, PowerUpCollision)

	SpeedPowerUpName := &PowerUpComponent{Name: 1}
	powerUpManager.AddComponent(powerUp, SpeedPowerUpName)

	playerMotion := &MotionComponent{Speed: 1.0, Velocity: Vec2{X: 0, Y: 0}, Acceleration: Vec2{X: 0, Y: 0}}
	playerInput := &InputComponent{Input: map[string]bool{"down": true}}
	motionManager.AddComponent(player, playerMotion)
	inputManager.AddComponent(player, playerInput)
	playerPc := positionManager.positions[player]

	for i := 0; i < 3; i++ {
		CallInputSystem.Update(0.1)
		CallMotionSystem.Update(0.1)

	}

	SpeedPowerUpPc := positionManager.positions[powerUp]
	mc := motionManager.motions[player]
	if mc.Speed != 2 {
		t.Fatalf("Expected player speed to be 2, got %v", mc.Speed)
	}
	if SpeedPowerUpPc != nil {
		t.Fatalf("Expected player position to be nil, got %v", SpeedPowerUpPc)
	}
	if playerPc.Y != 10 {
		t.Fatalf("Expected player Y position to be 10, got %v", playerPc.Y)
	}
}

func TestHealthPowerUpSystem(t *testing.T) {
	player := entityManager.CreateEntity("player")
	healthPowerUp1 := entityManager.CreateEntity("powerup")
	healthPowerUp2 := entityManager.CreateEntity("powerup")

	defer func() {
		DeleteAllEntityComponents(player)
		DeleteAllEntityComponents(healthPowerUp1)
		DeleteAllEntityComponents(healthPowerUp2)
	}()

	playerPosition := &PositionComponent{X: 10, Y: 5, Size: 1}
	HealthPowerUp1Position := &PositionComponent{X: 10, Y: 6, Size: 1}
	HealthPowerUp2Position := &PositionComponent{X: 10, Y: 7, Size: 1}
	positionManager.AddComponent(player, playerPosition)
	positionManager.AddComponent(healthPowerUp1, HealthPowerUp1Position)
	positionManager.AddComponent(healthPowerUp2, HealthPowerUp2Position)

	playerCollision := &CollisionComponent{Enabled: false}
	HealthPowerUp1Collision := &CollisionComponent{Enabled: true}
	HealthPowerUp2Collision := &CollisionComponent{Enabled: true}
	collisionManager.AddComponent(player, playerCollision)
	collisionManager.AddComponent(healthPowerUp1, HealthPowerUp1Collision)
	collisionManager.AddComponent(healthPowerUp2, HealthPowerUp2Collision)

	playerHealth := &HealthComponent{CurrentHealth: 2, MaxHealth: 3}
	healthManager.AddComponent(player, playerHealth)

	HealthPowerUp1Name := &PowerUpComponent{Name: PowerUpHealth}
	HealthPowerUp2Name := &PowerUpComponent{Name: PowerUpHealth}
	powerUpManager.AddComponent(healthPowerUp1, HealthPowerUp1Name)
	powerUpManager.AddComponent(healthPowerUp2, HealthPowerUp2Name)

	playerMotion := &MotionComponent{Speed: 1.0, Velocity: Vec2{X: 0, Y: 0}, Acceleration: Vec2{X: 0, Y: 0}}
	playerInput := &InputComponent{Input: map[string]bool{"down": true}}
	motionManager.AddComponent(player, playerMotion)
	inputManager.AddComponent(player, playerInput)

	for i := 0; i < 3; i++ {
		CallInputSystem.Update(0.1)
		CallMotionSystem.Update(0.1)
		CallHealthSystem.Update(0.1)

	}
	playerPc := positionManager.positions[player]
	playerHc := healthManager.healths[player]
	powerUpPc := positionManager.positions[healthPowerUp1]

	if powerUpPc != nil {
		t.Fatalf("Expected player position to be nil, got %v", powerUpPc)
	}
	if playerHc.CurrentHealth != 3 {
		t.Fatalf("Expected player Y position to be 3, got %v", playerHc.CurrentHealth)
	}
	if playerPc.Y != 8 {
		t.Fatalf("Expected player Y position to be 10, got %v", playerPc.Y)
	}
}

func TestBombPowerUpSystem(t *testing.T) {
	player := entityManager.CreateEntity("powerup")
	bombPowerUp := entityManager.CreateEntity("bomb")
	defer func() {
		DeleteAllEntityComponents(player)
		DeleteAllEntityComponents(bombPowerUp)
	}()

	playerPosition := &PositionComponent{X: 10, Y: 5, Size: 1}
	BombPowerUpPosition := &PositionComponent{X: 10, Y: 6, Size: 1}
	positionManager.AddComponent(player, playerPosition)
	positionManager.AddComponent(bombPowerUp, BombPowerUpPosition)

	playerCollision := &CollisionComponent{Enabled: false}
	BombPowerUpCollision := &CollisionComponent{Enabled: true}
	collisionManager.AddComponent(player, playerCollision)
	collisionManager.AddComponent(bombPowerUp, BombPowerUpCollision)

	playerPowerUps := &PowerUpComponent{ExtraBombs: 1, ExtraExplosionRange: 1, ExtraSpeed: 1}
	BombPowerUpName := &PowerUpComponent{Name: PowerUpBomb}
	powerUpManager.AddComponent(player, playerPowerUps)
	powerUpManager.AddComponent(bombPowerUp, BombPowerUpName)

	playerMotion := &MotionComponent{Speed: 1.0, Velocity: Vec2{X: 0, Y: 0}, Acceleration: Vec2{X: 0, Y: 0}}
	playerInput := &InputComponent{Input: map[string]bool{"down": true}}
	motionManager.AddComponent(player, playerMotion)
	inputManager.AddComponent(player, playerInput)

	for i := 0; i < 3; i++ {
		CallInputSystem.Update(0.1)
		CallMotionSystem.Update(0.1)

	}

	playerPc := positionManager.positions[player]
	bombPowerUpPc := powerUpManager.powerUps[player]

	powerUpPc := positionManager.positions[bombPowerUp]
	if powerUpPc != nil {
		t.Fatalf("Expected player position to be nil, got %v", powerUpPc)
	}
	if bombPowerUpPc.ExtraBombs != 2 {
		t.Fatalf("Expected player bomb amout to be 2, got %v", bombPowerUpPc.ExtraBombs)
	}
	if playerPc.Y != 8 {
		t.Fatalf("Expected player Y position to be 10, got %v", playerPc.Y)
	}
}

func TestExplosionPowerUpSystem(t *testing.T) {
	player := entityManager.CreateEntity("player")
	ExplosionPowerUp := entityManager.CreateEntity("explosion")
	defer func() {
		DeleteAllEntityComponents(player)
		DeleteAllEntityComponents(ExplosionPowerUp)
	}()

	playerPosition := &PositionComponent{X: 10, Y: 5, Size: 1}
	PowerUpPosition := &PositionComponent{X: 10, Y: 6, Size: 1}
	positionManager.AddComponent(player, playerPosition)
	positionManager.AddComponent(ExplosionPowerUp, PowerUpPosition)

	playerCollision := &CollisionComponent{Enabled: false}
	ExplosionPowerUpCollision := &CollisionComponent{Enabled: true}
	collisionManager.AddComponent(player, playerCollision)
	collisionManager.AddComponent(ExplosionPowerUp, ExplosionPowerUpCollision)

	playerPowerUps := &PowerUpComponent{ExtraBombs: 1, ExtraExplosionRange: 1, ExtraSpeed: 1}
	powerUpName := &PowerUpComponent{Name: PowerUpExplosion}
	powerUpManager.AddComponent(ExplosionPowerUp, powerUpName)
	powerUpManager.AddComponent(player, playerPowerUps)

	playerMotion := &MotionComponent{Speed: 1.0, Velocity: Vec2{X: 0, Y: 0}, Acceleration: Vec2{X: 0, Y: 0}}
	playerInput := &InputComponent{Input: map[string]bool{"down": true}}
	motionManager.AddComponent(player, playerMotion)
	inputManager.AddComponent(player, playerInput)

	for i := 0; i < 3; i++ {
		CallInputSystem.Update(0.1)
		CallMotionSystem.Update(0.1)

	}

	playerPc := positionManager.positions[player]
	playerPuc := powerUpManager.powerUps[player]
	explosionPowerUpPc := positionManager.positions[ExplosionPowerUp]
	if playerPuc.ExtraExplosionRange != 2 {
		t.Fatalf("Expected player explosion range to be 2, got %v", explosionPowerUpPc)
	}
	if explosionPowerUpPc != nil {
		t.Fatalf("Expected explosion power up position to be nil, got %v", explosionPowerUpPc)
	}
	if playerPc.Y != 8 {
		t.Fatalf("Expected player Y position to be 10, got %v", playerPc.Y)
	}

}
