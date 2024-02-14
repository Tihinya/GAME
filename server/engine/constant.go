package engine

import (
	"math"
	"time"
)

const (
	Up               = "KeyW"
	Down             = "KeyS"
	Left             = "KeyA"
	Right            = "KeyD"
	Space            = "Space"
	PowerUpSpeed     = 1
	PowerUpBomb      = 2
	PowerUpHealth    = 3
	PowerUpExplosion = 4
)

const (
	componentSize         = 40.0
	playerMaxHealth       = 3
	boxHealth             = 1
	defaultExplosionRange = 2
	explosionDamage       = 1
	defaultBombAmount     = 1
)

func CreatePlayer(socketId int, x, y float64) *Entity {
	player := entityManager.CreateEntity("player")

	playerUser := &UserEntityComponent{entity: player}
	playerCollision := &CollisionComponent{Enabled: true}
	playerPosition := &PositionComponent{X: x, Y: y, Size: componentSize}
	playerMotion := &MotionComponent{Velocity: Vec2{X: 0, Y: 0}, Acceleration: Vec2{X: 0, Y: 0}}
	playerInput := &InputComponent{Input: map[string]bool{}}
	playerHealth := &HealthComponent{CurrentHealth: playerMaxHealth, MaxHealth: playerMaxHealth}
	// playerPowerUps := &PowerUpComponent{ExtraBombs: 0, ExtraExplosionRange: 0, ExtraSpeed: 0}

	userEntityManager.AddComponent(socketId, playerUser)
	positionManager.AddComponent(player, playerPosition)
	motionManager.AddComponent(player, playerMotion)
	inputManager.AddComponent(player, playerInput)
	healthManager.AddComponent(player, playerHealth)
	// powerUpManager.AddComponent(player, playerPowerUps)
	collisionManager.AddComponent(player, playerCollision)

	return player
}

func CreateBomb(player *Entity) *Entity {
	// var playerActiveBombs int

	playerPosition := positionManager.GetPosition(player)
	// playerPowerUps := powerUpManager.powerUps[player]

	// for _, bc := range bombManager.bombs {
	// 	if bc.Owner == player {
	// 		playerActiveBombs++
	// 	}
	// }

	// if player.PlacedBombs > player.BombAmount {
	// 	return nil
	// }

	bomb := entityManager.CreateEntity("bomb")

	bombComponent := &BombComponent{
		BlastRadius: defaultExplosionRange,
		Owner:       player,
	}
	bombPosition := &PositionComponent{X: roundBase(playerPosition.X, componentSize), Y: roundBase(playerPosition.Y, componentSize), Size: componentSize}
	bombCollision := &CollisionComponent{Enabled: false}
	bombTimer := &TimerComponent{time.Now().Add(fuseTime)}

	collisionManager.AddComponent(bomb, bombCollision)
	timerManager.AddComponent(bomb, bombTimer)
	positionManager.AddComponent(bomb, bombPosition)
	bombManager.AddComponent(bomb, bombComponent)

	return bomb
}

func roundBase(num float64, base float64) float64 {
	modulo := math.Mod(num, base)

	num -= modulo

	if modulo < base/2 {
		return num
	} else {
		return num + base
	}
}

func SpreadExplosion(e *Entity) {
	pc := positionManager.GetPosition(e)
	bc := bombManager.GetBomb(e)

	if pc == nil || bc == nil {
		return
	}

	CreateExplosion(&PositionComponent{X: pc.X, Y: pc.Y, Size: componentSize})

	directions := []struct {
		dx, dy float64
	}{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}

	for _, dir := range directions {
		for i := 1; i < bc.BlastRadius; i++ {
			currentRadius := float64(i) * componentSize
			newPos := &PositionComponent{X: pc.X + currentRadius*dir.dx, Y: pc.Y + currentRadius*dir.dy}

			if IsWallAtPosition(newPos) {
				break
			}

			CreateExplosion(newPos)
			if box := GetBoxAtPosition(newPos); box != nil {
				DeleteAllEntityComponents(box)
				break
			}
		}
	}
}

func CreateExplosion(positionComponent *PositionComponent) {
	explosion := entityManager.CreateEntity("explosion")

	explosionComponent := &ExplosionComponent{}
	explosionPosition := positionComponent
	explosionTimer := &TimerComponent{time.Now().Add(explodeTime)}
	explosionDamage := &DamageComponent{DamageAmount: explosionDamage}

	timerManager.AddComponent(explosion, explosionTimer)
	positionManager.AddComponent(explosion, explosionPosition)
	damageManager.AddComponent(explosion, explosionDamage)
	explosionManager.AddComponent(explosion, explosionComponent)
}

func CreatePowerUp(powerUpName string) *Entity {
	powerUp := entityManager.CreateEntity("powerup")

	powerUpPosition := &PositionComponent{}
	powerUpProperty := &PowerUpComponent{Name: powerUpName}

	positionManager.AddComponent(powerUp, powerUpPosition)
	powerUpManager.AddComponent(powerUp, powerUpProperty)

	return powerUp
}

func CreateWall(X, Y float64) *Entity {
	wall := entityManager.CreateEntity("wall")
	wallCollision := &CollisionComponent{Enabled: true}
	wallPosition := &PositionComponent{X: X, Y: Y, Size: componentSize}
	wallIdentifier := &WallComponent{}

	collisionManager.AddComponent(wall, wallCollision)
	positionManager.AddComponent(wall, wallPosition)
	wallManager.AddComponent(wall, wallIdentifier)

	return wall
}

func CreateBox(X, Y float64) *Entity {
	box := entityManager.CreateEntity("box")

	boxCollision := &CollisionComponent{Enabled: true}
	playerPosition := &PositionComponent{X: X, Y: Y, Size: componentSize}
	playerHealth := &HealthComponent{CurrentHealth: boxHealth, MaxHealth: boxHealth}
	boxIdentifier := &BoxComponent{}

	collisionManager.AddComponent(box, boxCollision)
	positionManager.AddComponent(box, playerPosition)
	healthManager.AddComponent(box, playerHealth)
	boxManager.AddComponent(box, boxIdentifier)

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

func GetBoxAtPosition(pos *PositionComponent) *Entity {
	// Iterate over all entities to check for a wall at the given position
	for _, e := range entityManager.entities {
		if boxComp := boxManager.GetBox(e); boxComp != nil {
			boxPos := positionManager.GetPosition(e)
			if boxPos.X == pos.X && boxPos.Y == pos.Y {
				return e // Found a wall at the position
			}
		}
	}
	return nil
}
