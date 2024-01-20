package internal

import (
	"time"
)

var (
	Speed        = 1.0
	Acceleration = 1.0
	Regeneration = 1
	Bomb         = 1
)

const (
	Up    = "up"
	Down  = "down"
	Left  = "left"
	Right = "right"
	Space = "space"
)

const (
	componentSize         = 1.0
	playerMaxHealth       = 3
	boxHealth             = 1
	defaultExplosionRange = 6
	explosionDamage       = 1
	defaultBombAmount     = 1
)

const (
	LeftSprite  = "leftSprite.png"
	RightSprite = "rightSprite.png"
	UpSprite    = "upSprite.png"
	DownSprite  = "downSprite.png"
)

var (
	player1SpawnX = 10.0
	player1Spawny = 5.0
)

func CreatePlayer() *Entity {
	player := entityManager.CreateEntity()

	playerPosition := &PositionComponent{X: player1SpawnX, Y: player1Spawny, Size: componentSize}
	playerMotion := &MotionComponent{Velocity: Vec2{X: 0, Y: 0}, Acceleration: Vec2{X: 0, Y: 0}}
	playerInput := &InputComponent{Input: map[string]bool{}}
	playerHealth := &HealthComponent{CurrentHealth: playerMaxHealth, MaxHealth: playerMaxHealth}
	playerExposion := &ExplosionComponent{Range: defaultExplosionRange}
	playerPowerUps := &PowerUpComponent{ExtraBombs: 0, ExtraExplosionRange: 0, ExtraSpeed: 0}

	positionManager.AddComponent(player, playerPosition)
	motionManager.AddComponent(player, playerMotion)
	inputManager.AddComponent(player, playerInput)
	healthManager.AddComponent(player, playerHealth)
	explosionManager.AddComponent(player, playerExposion)
	powerUpManager.AddComponent(player, playerPowerUps)

	return player
}

func CreateBomb(player *Entity) *Entity {
	var playerActiveBombs int

	pc := positionManager.positions[player]
	ec := explosionManager.explosions[player]
	puc := powerUpManager.powerUps[player]
	for _, bc := range bombManager.bombs {
		if bc.Owner == player {
			playerActiveBombs++
		}
	}

	if !(playerActiveBombs < (1 + puc.ExtraBombs)) {
		return nil
	}

	bomb := entityManager.CreateEntity()

	bombComponent := &BombComponent{
		BlastRadius: ec.Range + puc.ExtraExplosionRange,
		IsActive:    true,
		Owner:       player,
	}
	bombPosition := &PositionComponent{X: pc.X, Y: pc.Y, Size: pc.Size}
	bombCollision := &CollisionComponent{Enabled: false}
	bombTimer := &TimerComponent{time.Now().Add(fuseTime)}
	bombExplosion := &ExplosionComponent{Range: ec.Range}

	collisionManager.AddComponent(bomb, bombCollision)
	timerManager.AddComponent(bomb, bombTimer)
	positionManager.AddComponent(bomb, bombPosition)
	explosionManager.AddComponent(bomb, bombExplosion)
	bombManager.AddComponent(bomb, bombComponent)

	return bomb
}

func SpreadExplosion(pc *PositionComponent, puc *PowerUpComponent) {
	// Create an explosion at the bomb's position
	createExplosionAtPosition(pc.X, pc.Y)

	// Spread the explosion in each direction
	for i := 1; i <= (defaultExplosionRange + puc.ExtraExplosionRange); i++ {
		createExplosionAtPosition(pc.X+float64(i), pc.Y) // Right
		createExplosionAtPosition(pc.X-float64(i), pc.Y) // Left
		createExplosionAtPosition(pc.X, pc.Y+float64(i)) // Up
		createExplosionAtPosition(pc.X, pc.Y-float64(i)) // Down
	}
}

func createExplosionAtPosition(X, Y float64) {
	CreateExplosion(&PositionComponent{X: X, Y: Y})
}

func CreateExplosion(positionComponent *PositionComponent) {
	// Collision check here idk how tf to do it

	explosion := entityManager.CreateEntity()

	explosionPosition := positionComponent
	explosionTimer := &TimerComponent{time.Now().Add(explodeTime)}
	explosionDamage := &DamageComponent{DamageAmount: explosionDamage}

	timerManager.AddComponent(explosion, explosionTimer)
	positionManager.AddComponent(explosion, explosionPosition)
	damageManager.AddComponent(explosion, explosionDamage)
}

func CreatePowerUp(powerUpName string) *Entity {
	powerUp := entityManager.CreateEntity()

	powerUpPosition := &PositionComponent{}
	// powerUpProperty := &PowerUpComponent{Name: powerUpName}
	positionManager.AddComponent(powerUp, powerUpPosition)
	// powerUpManager.AddComponent(powerUp, powerUpProperty)

	return powerUp
}

func CreateWall() *Entity {
	wall := entityManager.CreateEntity()

	playerPosition := &PositionComponent{}

	positionManager.AddComponent(wall, playerPosition)
	return wall
}

func CreateBox() *Entity {
	box := entityManager.CreateEntity()

	playerPosition := &PositionComponent{}
	playerHealth := &HealthComponent{CurrentHealth: boxHealth, MaxHealth: boxHealth}

	positionManager.AddComponent(box, playerPosition)
	healthManager.AddComponent(box, playerHealth)

	return box
}
