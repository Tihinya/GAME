package engine

import (
	"math"
	"math/rand"
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
	PowerUpExplosion = 3
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
	playerPosition := &PositionComponent{X: x, Y: y, Size: 30}
	playerMotion := &MotionComponent{Velocity: Vec2{X: 0, Y: 0}, Acceleration: Vec2{X: 0, Y: 0}, SpeedMultiplier: 1}
	playerInput := &InputComponent{Input: map[string]bool{}}
	playerBomb := &BombComponent{BlastRadius: 2, BombAmount: 1}
	playerHealth := &HealthComponent{CurrentHealth: playerMaxHealth, MaxHealth: playerMaxHealth}

	playerHealth.OnDestroy = func() {
		if playerHealth.CurrentHealth > 0 {
			playerPosition.X = x
			playerPosition.Y = y
		} else {
			DeleteAllEntityComponents(player)
		}
	}

	bombManager.AddComponent(player, playerBomb)
	userEntityManager.AddComponent(socketId, playerUser)
	positionManager.AddComponent(player, playerPosition)
	motionManager.AddComponent(player, playerMotion)
	inputManager.AddComponent(player, playerInput)
	healthManager.AddComponent(player, playerHealth)
	collisionManager.AddComponent(player, playerCollision)

	return player
}

func CreateBomb(player *Entity) *Entity {
	playerPosition := positionManager.GetPosition(player)
	bombPosition := &PositionComponent{X: roundBase(playerPosition.X, componentSize), Y: roundBase(playerPosition.Y, componentSize), Size: componentSize}

	for e := range bombManager.bombs {
		existingBombPosition := positionManager.positions[e]
		if bombPosition.X == existingBombPosition.X && bombPosition.Y == existingBombPosition.Y {
			return nil
		}
	}

	playerBomb := bombManager.bombs[player]
	if playerBomb.PlacedBombs >= playerBomb.BombAmount {
		return nil
	}
	playerBomb.PlacedBombs += 1
	bomb := entityManager.CreateEntity("bomb")

	Bomb := &BombComponent{
		BlastRadius: playerBomb.BlastRadius,
		Owner:       player,
	}
	bombCollision := &CollisionComponent{Enabled: false}
	bombTimer := &TimerComponent{time.Now().Add(fuseTime)}
	explosionStopper := &ExplosionStopperComponent{passable: false}

	collisionManager.AddComponent(bomb, bombCollision)
	timerManager.AddComponent(bomb, bombTimer)
	positionManager.AddComponent(bomb, bombPosition)
	bombManager.AddComponent(bomb, Bomb)
	explosionStopperManager.AddComponent(bomb, explosionStopper)

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
			newPos := &PositionComponent{X: pc.X + currentRadius*dir.dx, Y: pc.Y + currentRadius*dir.dy, Size: componentSize}

			stopper := isExplosionBlocked(newPos)
			if stopper == nil {
				CreateExplosion(newPos)
				continue
			}

			if stopper.passable {
				CreateExplosion(newPos)
			}
			break
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

func CreatePowerUp(powerUpName string, x, y float64) *Entity {
	powerUp := entityManager.CreateEntity("powerup")

	powerUpPosition := &PositionComponent{X: x, Y: y, Size: componentSize}
	powerUpProperty := &PowerUpComponent{Name: powerUpName}

	positionManager.AddComponent(powerUp, powerUpPosition)
	powerUpManager.AddComponent(powerUp, powerUpProperty)

	return powerUp
}

func CreateWall(X, Y float64) *Entity {
	wall := entityManager.CreateEntity("wall")
	wallCollision := &CollisionComponent{Enabled: true}
	wallPosition := &PositionComponent{X: X, Y: Y, Size: componentSize}
	explosionStopper := &ExplosionStopperComponent{passable: false}

	collisionManager.AddComponent(wall, wallCollision)
	positionManager.AddComponent(wall, wallPosition)
	explosionStopperManager.AddComponent(wall, explosionStopper)

	return wall
}

func CreateBox(X, Y float64) *Entity {
	box := entityManager.CreateEntity("box")

	boxCollision := &CollisionComponent{Enabled: true}
	boxPosition := &PositionComponent{X: X, Y: Y, Size: componentSize}
	boxHealth := &HealthComponent{CurrentHealth: boxHealth, MaxHealth: boxHealth}
	explosionStopper := &ExplosionStopperComponent{passable: true}

	boxHealth.OnDestroy = func() {
		DeleteAllEntityComponents(box)
		if rand.Intn(101) <= 60 {
			switch rand.Intn(3) + 1 {
			case PowerUpSpeed:
				CreatePowerUp("speedPowerUp", X, Y)
			case PowerUpBomb:
				CreatePowerUp("bombPowerUp", X, Y)
			case PowerUpExplosion:
				CreatePowerUp("exposionPowerup", X, Y)
			}
		}
	}

	collisionManager.AddComponent(box, boxCollision)
	positionManager.AddComponent(box, boxPosition)
	healthManager.AddComponent(box, boxHealth)
	explosionStopperManager.AddComponent(box, explosionStopper)

	return box
}

func isExplosionBlocked(pos *PositionComponent) *ExplosionStopperComponent {
	for e, stopper := range explosionStopperManager.explosionStoppers {
		unbreakablePosition := positionManager.positions[e]
		if pos.X == unbreakablePosition.X && pos.Y == unbreakablePosition.Y {
			return stopper
		}
	}
	return nil
}
