package engine

import (
	"fmt"
	"time"

	"bomberman-dom/models"
)

var (
	fuseTime    = time.Millisecond * 2000
	explodeTime = time.Millisecond * 1500
)

// --------------------------------
// Createing Managers
// --------------------------------

var (
	entityManager     = NewEntityManager()
	positionManager   = NewPositionManager()
	motionManager     = NewMotionManager()
	inputManager      = NewInputManager()
	timerManager      = NewTimerManager()
	bombManager       = NewBombManager()
	explosionManager  = NewExplosionManager()
	collisionManager  = NewCollisionManager()
	powerUpManager    = NewPowerUpManager()
	damageManager     = NewDamageManager()
	healthManager     = NewHealthManager()
	boxManager        = NewBoxManager()
	wallManager       = NewWallManager()
	userEntityManager = NewUserEntityManager()
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

// var damageSystem = NewDamageSystem()
// var bombSystem = NewBombSystem()

// Solving import cycles the interface way!!
// var broadcaster helpers.Broadcaster

// func SetBroadcaster(b helpers.Broadcaster) {
// 	broadcaster = b
// }

func (dm *DamageSystem) Update(dt float64) {
	for e, damge := range dm.manager.damages {
		if collistion, entity := DetectCollision(e); collistion {
			entityHealth := healthManager.healths[entity]
			entityHealth.CurrentHealth -= damge.DamageAmount
			if entityHealth.CurrentHealth < 1 {
				DeleteAllEntityComponents(e)
			}
		}
	}
}

const SpeedPowerUp = 10
const HealthPowerUp = 1

func (pw *PowerUpSystem) Update(dt float64) {
	for powerup, powerUpComponent := range pw.manager.powerUps {
		if collistion, player := DetectCollision(powerup); !collistion {
			if powerUpComponent.Name == "speedPowerUp" {
				mc := motionManager.motions[player]
				mc.Speed += SpeedPowerUp
				DeleteAllEntityComponents(powerup)
			} else if powerUpComponent.Name == "bombPowerUp" {
				DeleteAllEntityComponents(powerup)
			} else if powerUpComponent.Name == "healthPowerUp" {
				hl := healthManager.healths[player]
				hl.MaxHealth += HealthPowerUp
				DeleteAllEntityComponents(powerup)
			} else if powerUpComponent.Name == "exposionPowerup" {
				DeleteAllEntityComponents(powerup)
			}
		}
	}
}

func (mv *MotionSystem) Update(dt float64) {
	for e, mc := range mv.manager.motions {
		pc, exists := positionManager.positions[e]
		if !exists || (mc.Velocity.X == 0 && mc.Velocity.Y == 0) {
			return
		}
		pc.X += mc.Velocity.X
		pc.Y += mc.Velocity.Y

		mc.Velocity.X += mc.Acceleration.X
		mc.Velocity.Y += mc.Acceleration.Y

		if collistion, _ := DetectCollision(e); collistion {
			pc.X -= mc.Velocity.X
			pc.Y -= mc.Velocity.Y

			mc.Velocity.X -= mc.Acceleration.X
			mc.Velocity.Y -= mc.Acceleration.Y
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
			bomb := CreateBomb(e)
			fmt.Println(bomb)
			// if collistion, _ := DetectCollision(bomb, e); collistion {
			// 	delete(positionManager.positions, bomb)
			// }
		}
	}
}

func (hs *HealthSystem) Update(dt float64) {
	for e, hc := range hs.manager.healths {
		if hc.CurrentHealth <= 0 {
			DeleteAllEntityComponents(e)
		}
		if hc.CurrentHealth > hc.MaxHealth {
			hc.CurrentHealth = hc.MaxHealth
		}
	}
}

func (ex *ExplosionSystem) Update(dt float64) {
	for e := range bombManager.bombs {
		bombTimer := timerManager.GetTimer(e)
		if bombTimer == nil {
			fmt.Println("No timer found for bomb", e.Id)
			continue // Skip if no timer is set for this bomb
		}
		if time.Now().After(bombTimer.Time) {
			SpreadExplosion(e)
			DeleteAllEntityComponents(e)
		}
	}

	for e2 := range explosionManager.explosions {
		explosionTimer := timerManager.GetTimer(e2)
		if explosionTimer == nil {
			fmt.Println("No timer found for explosion", e2.Id)
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

func DetectCollision(e1 *Entity, ignoreList ...*Entity) (bool, *Entity) {
	currentCollider := positionManager.positions[e1]
	// mc, mcExists := motionManager.motions[e1]

	// collisionManager.mutex.RLock()

	for e2, cc := range collisionManager.collisions {
		if e1 == e2 || !cc.Enabled || contains(ignoreList, e2) {
			continue
		}

		foundedCollider := positionManager.positions[e2]
		// powerUp, isPowerUp := powerUpManager.powerUps[e2]

		collides := (currentCollider.X < foundedCollider.X+foundedCollider.Size) &&
			(currentCollider.X+currentCollider.Size > foundedCollider.X) &&
			(currentCollider.Y < foundedCollider.Y+foundedCollider.Size) &&
			(currentCollider.Y+currentCollider.Size > foundedCollider.Y)

		// log.Print(cc.Enabled && collides, currentCollider, foundedCollider)
		if cc.Enabled && collides {
			return true, e2
		}
		// if isPowerUp && cc.Disabled && collides {
		// 	playerPowerUps, exists := powerUpManager.powerUps[e1]
		// 	if !exists {
		// 		return collides
		// 	}

		// 	switch powerUp.Name {
		// 	case PowerUpSpeed:
		// 		if mcExists {
		// 			mc.Speed += Speed
		// 			DeleteAllEntityComponents(e2)
		// 		}
		// 	case PowerUpHealth:
		// 		if hc, exists := healthManager.healths[e1]; exists {
		// 			hc.CurrentHealth += Regeneration
		// 			DeleteAllEntityComponents(e2)
		// 		}
		// 	case PowerUpBomb:
		// 		playerPowerUps.ExtraBombs += Bomb
		// 		DeleteAllEntityComponents(e2)
		// 	case PowerUpExplosion:
		// 		playerPowerUps.ExtraExplosionRange += ExplosionRange
		// 		DeleteAllEntityComponents(e2)
		// 	}
		// 	return false
		// }
	}
	// collisionManager.mutex.RUnlock()
	return false, nil
}

func contains[T comparable](itemArray []T, item T) bool {
	for _, it := range itemArray {
		if it == item {
			return true
		}
	}
	return false
}
