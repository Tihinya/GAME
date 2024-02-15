package engine

import (
	"time"

	"bomberman-dom/models"
)

var (
	fuseTime    = time.Millisecond * 3000
	explodeTime = time.Millisecond * 1500
)

// --------------------------------
// Createing Managers
// --------------------------------

var (
	entityManager           = NewEntityManager()
	positionManager         = NewPositionManager()
	motionManager           = NewMotionManager()
	inputManager            = NewInputManager()
	timerManager            = NewTimerManager()
	bombManager             = NewBombManager()
	explosionManager        = NewExplosionManager()
	collisionManager        = NewCollisionManager()
	powerUpManager          = NewPowerUpManager()
	damageManager           = NewDamageManager()
	healthManager           = NewHealthManager()
	explosionStopperManager = NewExplosionStopperManager()
	userEntityManager       = NewUserEntityManager()
)

// --------------------------------
// Createing Systems
// --------------------------------

var (
	CallMotionSystem    = NewMotionSystem()
	CallInputSystem     = NewInputSystem()
	CallHealthSystem    = NewHealthSystem()
	CallExplosionSystem = NewExplosionSystem()
	CallDamageSystem    = NewDamageSystem()
	CallPowerUpSystem   = NewPowerUpSystem()
)

const (
	SpeedPowerUp  = 0.2
	HealthPowerUp = 1
	BombPowerUp   = 1
	BlastPowerUp  = 1
)

func (pw *PowerUpSystem) Update(dt float64) {
	for powerUpEntity, powerUpComponent := range pw.manager.powerUps {
		powerUpPosition := positionManager.positions[powerUpEntity]
		for playerEntity, motionComponent := range motionManager.motions {
			playerPosition := positionManager.positions[playerEntity]
			if !isRectangleCollides(powerUpPosition, playerPosition) {
				continue
			}
			// log.Println(powerUpComponent.Name)
			if powerUpComponent.Name == "speedPowerUp" {
				motionComponent.SpeedMultiplier += SpeedPowerUp
				DeleteAllEntityComponents(powerUpEntity)
			} else if powerUpComponent.Name == "bombPowerUp" {
				bombComponent := bombManager.bombs[playerEntity]
				bombComponent.BombAmount += BombPowerUp
				DeleteAllEntityComponents(powerUpEntity)
			} else if powerUpComponent.Name == "exposionPowerup" {
				bombComponent := bombManager.bombs[playerEntity]
				bombComponent.BlastRadius += BlastPowerUp
				DeleteAllEntityComponents(powerUpEntity)
			}
		}
	}
}

func (hs *HealthSystem) Update(dt float64) {
	for damageEntity, damage := range damageManager.damages {
		damagePosition := positionManager.positions[damageEntity]
		for healthEntity, health := range healthManager.healths {
			healthPosition := positionManager.positions[healthEntity]

			if time.Since(health.lastTimeDamage) <= time.Second*2 || !isRectangleCollides(damagePosition, healthPosition) {
				continue
			}
			health.CurrentHealth -= damage.DamageAmount
			health.lastTimeDamage = time.Now()
			health.OnDestroy()
		}
	}
}

func (mv *MotionSystem) Update(dt float64) {
	for player, playerMotion := range mv.manager.motions {
		playerPosition := positionManager.positions[player]
		if playerMotion.Velocity.X == 0 && playerMotion.Velocity.Y == 0 {
			continue
		}
		playerPosition.X += playerMotion.Velocity.X * playerMotion.SpeedMultiplier
		playerPosition.Y += playerMotion.Velocity.Y * playerMotion.SpeedMultiplier

		// mc.Velocity.X += mc.Acceleration.X
		// mc.Velocity.Y += mc.Acceleration.Y

		if collistion := DetectCollision(player); collistion {
			playerPosition.X -= playerMotion.Velocity.X * playerMotion.SpeedMultiplier
			playerPosition.Y -= playerMotion.Velocity.Y * playerMotion.SpeedMultiplier

			// mc.Velocity.X -= mc.Acceleration.X
			// mc.Velocity.Y -= mc.Acceleration.Y
		}
	}
}

func (is *InputSystem) Update(dt float64) {
	for e, ic := range is.manager.inputs {
		mc := motionManager.motions[e]
		mc.Velocity.X = 0
		mc.Velocity.Y = 0
		if ic.Input[Up] {
			mc.Velocity.Y = -1
		}
		if ic.Input[Down] {
			mc.Velocity.Y = 1
		}
		if ic.Input[Left] {
			mc.Velocity.X = -1
		}
		if ic.Input[Right] {
			mc.Velocity.X = 1
		}
		if ic.Input[Space] {
			CreateBomb(e)
		}
	}
}

func (ex *ExplosionSystem) Update(dt float64) {
	for e, bombComponent := range bombManager.bombs {
		bombTimer := timerManager.GetTimer(e)
		if bombTimer == nil {
			continue // Skip if no timer is set for this bomb
		}
		if time.Now().After(bombTimer.Time) {
			player := bombComponent.Owner
			if player == nil {
				continue
			}

			playerBomb := bombManager.bombs[player]
			playerBomb.PlacedBombs -= 1
			SpreadExplosion(e)
			DeleteAllEntityComponents(e)
		}
	}

	for e2 := range explosionManager.explosions {
		explosionTimer := timerManager.GetTimer(e2)
		if explosionTimer == nil {
			continue
		}
		if time.Now().After(explosionTimer.Time) {
			DeleteAllEntityComponents(e2)
		}
	}
}

func HandleInput(input models.GameInput, playerId int) {
	player := userEntityManager.GetUserEntity(playerId)
	if e, exists := inputManager.inputs[player.entity]; exists {
		e.Input = input.Keys
	}
}

func DetectCollision(e1 *Entity, ignoreList ...*Entity) bool {
	currentCollider := positionManager.positions[e1]

	for e2, cc := range collisionManager.collisions {
		if e1 == e2 || !cc.Enabled || contains(ignoreList, e2) {
			continue
		}

		foundedCollider := positionManager.positions[e2]

		collides := isRectangleCollides(currentCollider, foundedCollider)

		if collides {
			return true
		}
	}
	return false
}

func isRectangleCollides(firstEntity, secondEntity *PositionComponent) bool {
	return (firstEntity.X < secondEntity.X+secondEntity.Size) &&
		(firstEntity.X+firstEntity.Size > secondEntity.X) &&
		(firstEntity.Y < secondEntity.Y+secondEntity.Size) &&
		(firstEntity.Y+firstEntity.Size > secondEntity.Y)
}

func contains[T comparable](itemArray []T, item T) bool {
	for _, it := range itemArray {
		if it == item {
			return true
		}
	}
	return false
}
