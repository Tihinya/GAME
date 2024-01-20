package internal

import (
	"fmt"
	"time"
)

var (
	fuseTime    = time.Millisecond * 500
	explodeTime = time.Millisecond * 150
)

// --------------------------------
// Createing Managers
// --------------------------------

var (
	entityManager    = NewEntityManager()
	positionManager  = NewPositionManager()
	motionManager    = NewMotionManager()
	inputManager     = NewInputManager()
	timerManager     = NewTimerManager()
	bombManager      = NewBombManager()
	explosionManager = NewExplosionManager()
	collisionManager = NewCollisionManager()
	powerUpManager   = NewPowerUpManager()
	damageManager    = NewDamageManager()
	healthManager    = NewHealthManager()
)

// --------------------------------
// Createing Systems
// --------------------------------

var (
	motionSystem    = NewMotionSystem()
	inputSystem     = NewInputSystem()
	powerUpSystem   = NewPowerUpSystem()
	healthSystem    = NewHealthSystem()
	explosionSystem = NewExplosionSystem()
)

// var damageSystem = NewDamageSystem()
// var bombSystem = NewBombSystem()
// var explosionSystem = NewExplosionSystem()

func (mv *MotionSystem) update(dt float64) {
	for e, mc := range mv.manager.motions {
		pc, exists := positionManager.positions[e]
		if !exists {
			return
		}
		pc.X += mc.Velocity.X
		pc.Y += mc.Velocity.Y

		mc.Velocity.X += mc.Acceleration.X
		mc.Velocity.Y += mc.Acceleration.Y

		if DetectCollision(e) {
			pc.X -= mc.Velocity.X
			pc.Y -= mc.Velocity.Y

			mc.Velocity.X -= mc.Acceleration.X
			mc.Velocity.Y -= mc.Acceleration.Y
		}
	}
}

func (is *InputSystem) update(dt float64) {
	for e, ic := range is.manager.inputs {
		mc := motionManager.motions[e]
		if ic.Input[Up] {
			mc.Velocity.Y = -Speed
		}
		if ic.Input[Down] {
			mc.Velocity.Y = Speed
		}
		if ic.Input[Left] {
			mc.Velocity.X = -Speed
		}
		if ic.Input[Right] {
			mc.Velocity.X = Speed
		}
		if ic.Input[Space] {
			bomb := CreateBomb(e)
			if DetectCollision(bomb, e) {
				delete(positionManager.positions, bomb)
			}

		}
	}
}

func (hs *HealthSystem) update(dt float64) {
	for e, hc := range hs.manager.healths {
		if hc.CurrentHealth <= 0 {
			DeleteAllEntityComponents(e)
		}
	}
}

func (ex *ExplosionSystem) update(dt float64) {
	for e := range bombManager.bombs {
		bombTimer, exists := timerManager.timers[e]
		if !exists || bombTimer == nil {
			fmt.Println("No timer found for bomb", e.Id)
			continue // Skip if no timer is set for this bomb
		}
		if time.Now().After(bombTimer.Time) {
			puc := powerUpManager.powerUps[e]
			pc := positionManager.positions[e]
			SpreadExplosion(pc, puc)
		}
	}
}

func DetectCollision(entity *Entity, ignoreList ...*Entity) bool {
	pc1 := positionManager.positions[entity]

	for _, e2 := range entityManager.entities {
		if entity == e2 && contains(ignoreList, entity) {
			continue
		}

		pc2 := positionManager.positions[e2]
		collides := pc1.X < pc2.X+pc2.Size && pc1.X+pc2.Size > pc2.X && pc1.Y < pc2.Y+pc2.Size && pc1.Y+pc1.Size > pc2.Y
		return collides
	}
	return false
}

func contains[T comparable](itemArray []T, item T) bool {
	for _, it := range itemArray {
		if it == item {
			return true
		}
	}
	return false
}
