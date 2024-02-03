package engine

import (
	"time"
)

const (
	Speed          = 1.0
	Acceleration   = 1.0
	Regeneration   = 1
	Bomb           = 1
	ExplosionRange = 1
)

const (
	Up               = "up"
	Down             = "down"
	Left             = "left"
	Right            = "right"
	Space            = "space"
	PowerUpSpeed     = 1
	PowerUpBomb      = 2
	PowerUpHealth    = 3
	PowerUpExplosion = 4
)

const (
	componentSize         = 1.0
	playerMaxHealth       = 3
	boxHealth             = 1
	defaultExplosionRange = 2
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

func CreatePlayer(socketId int) *Entity {
	player := entityManager.CreateEntity()

	playerUser := &UserEntityComponent{entity: player}
	playerPosition := &PositionComponent{X: player1SpawnX, Y: player1Spawny, Size: componentSize}
	playerMotion := &MotionComponent{Velocity: Vec2{X: 0, Y: 0}, Acceleration: Vec2{X: 0, Y: 0}}
	playerInput := &InputComponent{Input: map[string]bool{}}
	playerHealth := &HealthComponent{CurrentHealth: playerMaxHealth, MaxHealth: playerMaxHealth}
	playerPowerUps := &PowerUpComponent{ExtraBombs: 0, ExtraExplosionRange: 0, ExtraSpeed: 0}

	userEntityManager.AddComponent(socketId, playerUser)
	positionManager.AddComponent(player, playerPosition)
	motionManager.AddComponent(player, playerMotion)
	inputManager.AddComponent(player, playerInput)
	healthManager.AddComponent(player, playerHealth)
	powerUpManager.AddComponent(player, playerPowerUps)

	return player
}

func CreateBomb(player *Entity) *Entity {
	var playerActiveBombs int

	pc := positionManager.GetPosition(player)
	puc := powerUpManager.powerUps[player]

	bombManager.mutex.RLock()
	for _, bc := range bombManager.bombs {
		bombManager.mutex.RUnlock()
		if bc.Owner == player {
			playerActiveBombs++
		}
		bombManager.mutex.RLock()
	}
	bombManager.mutex.RUnlock()

	if !(playerActiveBombs < (1 + puc.ExtraBombs)) {
		return nil
	}

	bomb := entityManager.CreateEntity()

	bombComponent := &BombComponent{
		BlastRadius: defaultExplosionRange + puc.ExtraExplosionRange,
		IsActive:    true,
		Owner:       player,
	}
	bombPosition := &PositionComponent{X: pc.X, Y: pc.Y, Size: pc.Size}
	bombCollision := &CollisionComponent{Enabled: false}
	bombTimer := &TimerComponent{time.Now().Add(fuseTime)}

	collisionManager.AddComponent(bomb, bombCollision)
	timerManager.AddComponent(bomb, bombTimer)
	positionManager.AddComponent(bomb, bombPosition)
	bombManager.AddComponent(bomb, bombComponent)

	broadcastBomb(pc.X, pc.Y, "create")

	return bomb
}

func SpreadExplosion(e *Entity) {
	pc := positionManager.GetPosition(e)
	bc := bombManager.GetBomb(e)

	if pc == nil || bc == nil {
		return
	}

	// Create an explosion at the bomb's position
	createExplosionAtPosition(&PositionComponent{X: pc.X, Y: pc.Y, Size: 1})

	directions := []struct {
		dx, dy float64
	}{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}

	for _, dir := range directions {
		for i := 1; i < bc.BlastRadius; i++ {
			newPos := &PositionComponent{X: pc.X + float64(i)*dir.dx, Y: pc.Y + float64(i)*dir.dy}

			if IsWallAtPosition(newPos) {
				break // Wall is blocking this direction
			}
			if IsBoxAtPosition(newPos) {
				createExplosionAtPosition(newPos)
				break // Explode at wall and block all next explosions
			}
			createExplosionAtPosition(newPos)
		}
	}
}

func createExplosionAtPosition(pos *PositionComponent) {
	ExplodeBox(pos)
	CreateExplosion(pos)
}

func CreateExplosion(positionComponent *PositionComponent) {
	explosion := entityManager.CreateEntity()

	explosionComponent := &ExplosionComponent{}
	explosionPosition := positionComponent
	explosionTimer := &TimerComponent{time.Now().Add(explodeTime)}
	explosionDamage := &DamageComponent{DamageAmount: explosionDamage}

	timerManager.AddComponent(explosion, explosionTimer)
	positionManager.AddComponent(explosion, explosionPosition)
	damageManager.AddComponent(explosion, explosionDamage)
	explosionManager.AddComponent(explosion, explosionComponent)

	broadcastExplosion(positionComponent.X, positionComponent.Y, "create")
}

func CreatePowerUp(powerUpName int) *Entity {
	powerUp := entityManager.CreateEntity()

	powerUpPosition := &PositionComponent{}
	powerUpProperty := &PowerUpComponent{Name: powerUpName}

	positionManager.AddComponent(powerUp, powerUpPosition)
	powerUpManager.AddComponent(powerUp, powerUpProperty)

	broadcastPowerup(powerUpPosition.X, powerUpPosition.Y, powerUpName, "create")

	return powerUp
}

func CreateWall(X, Y float64) *Entity {
	wall := entityManager.CreateEntity()

	playerPosition := &PositionComponent{X: X, Y: Y, Size: 1}
	wallIdentifier := &WallComponent{}

	positionManager.AddComponent(wall, playerPosition)
	wallManager.AddComponent(wall, wallIdentifier)

	broadcastObstacle(X, Y, "wall", "create")

	return wall
}

func CreateBox(X, Y float64) *Entity {
	box := entityManager.CreateEntity()

	playerPosition := &PositionComponent{X: X, Y: Y, Size: 1}
	playerHealth := &HealthComponent{CurrentHealth: boxHealth, MaxHealth: boxHealth}
	boxIdentifier := &BoxComponent{}

	positionManager.AddComponent(box, playerPosition)
	healthManager.AddComponent(box, playerHealth)
	boxManager.AddComponent(box, boxIdentifier)

	broadcastObstacle(X, Y, "box", "create")

	return box
}

func IsWallAtPosition(pos *PositionComponent) bool {
	// Iterate over all entities to check for a wall at the given position
	for _, e := range entityManager.entities {
		if wallComp := wallManager.GetWall(e); wallComp != nil {
			wallPos := positionManager.GetPosition(e)
			if wallPos != nil && wallPos.X == pos.X && wallPos.Y == pos.Y {
				return true // Found a wall at the position
			}
		}
	}
	return false
}

func IsBoxAtPosition(pos *PositionComponent) bool {
	// Iterate over all entities to check for a wall at the given position
	for _, e := range entityManager.entities {
		if boxComp := boxManager.GetBox(e); boxComp != nil {
			boxPos := positionManager.GetPosition(e)
			if boxPos != nil && boxPos.X == pos.X && boxPos.Y == pos.Y {
				return true // Found a wall at the position
			}
		}
	}
	return false
}
