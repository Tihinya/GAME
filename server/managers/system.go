package managers

import (
	"time"
)

// --------------------------------
// Createing Managers
// --------------------------------

var entityManager = NewEntityManager()
var positionManager = NewPositionManager()
var motionManager = NewMotionManager()
var inputManager = NewInputManager()
var timerManager = NewTimerManager()
var bombManager = NewBombManager()
var explosionManager = NewExplosionManager()
var collectionManager = NewCollisionManager()
var powerUpManager = NewPowerUpManager()
var damageManager = NewDamageManager()
var healthManager = NewHealthManager()

// --------------------------------
// Createing Systems
// --------------------------------

var positionSystem = NewPositionSystem()
var motionSystem = NewMotionSystem()
var healthSystem = NewHealthSystem()
var inputSystem = NewInputSystem()
var timerSystem = NewTimerSystem()
var powerUpSystem = NewPowerUpSystem()
var damageSystem = NewDamageSystem()
var bombSystem = NewBombSystem()
var explosionSystem = NewExplosionSystem()

func (mv *MotionSystem) update(dt time.Time) {
	for e, mc := range mv.manager.motions {
		pc, exists := positionManager.postions[e]
		if !exists {
			return
		}
		pc.X += mc.Velocity.X
		pc.Y += mc.Velocity.Y

		mc.Velocity.X += mc.Acceleration.X
		mc.Velocity.Y += mc.Acceleration.Y

		if DetectCollisionSystem(e) {
			pc.X -= mc.Velocity.X
			pc.Y -= mc.Velocity.Y

			mc.Velocity.X -= mc.Acceleration.X
			mc.Velocity.Y -= mc.Acceleration.Y
		}
	}
}

func (is *InputSystem) update(dt time.Time) {
	for e, ic := range is.manager.inputs {
		mc, exists := motionManager.motions[e]
		if !exists {
			return
		}
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

	}
}

func (hs *HealthSystem) update(dt time.Time) {
	for e, hc := range hs.manager.healths {
		if hc.CurrentHealth <= 0 {
			DeleteAllEntityComponents(e)
		}
	}
}

func (pus *PowerUpSystem) update(dt time.Time) {
	for e, pum := range pus.manager.powerUps {
		mc := motionManager.motions[e]
		hc := healthManager.healths[e]
		bc := bombManager.bombs[e]
		for _, puc := range pum {
			switch puc.Name {
			case "speed":
				mc.Acceleration.X += Acceleration
				mc.Acceleration.Y += Acceleration
			case "health":
				hc.CurrentHealth += Regeneration
			case "bomb":
				bc.BombAmount += Bomb
			}
		}
	}

}

// func explosionSystem(entity *Entity, system *SystemManagers) {
// 	timer := entity.getTimer(system)
// 	if timer == nil {
// 		return
// 	}
// 	if !timer.Time.Before(time.Now()) {
// 		return
// 	}

// 	for entity2 := range system.DamageManager.damages {
// 		if entity != entity2 {
// 			fmt.Println("df")
// 		}
// 	}
// }

func DetectCollisionSystem(entity *Entity) bool {
	pc1 := positionManager.postions[entity]

	for _, e2 := range entityManager.entities {
		if entity == e2 {
			continue
		}
		pc2 := positionManager.postions[e2]

		// x1 < x2 + siz2 && x1 +siz1 > x2 && y1 <y2+siz2 && y1+siz1 >y2

		return pc1.X < pc2.X+pc2.Size && pc1.X+pc2.Size > pc2.X && pc1.Y < pc2.Y+pc2.Size && pc1.Y+pc1.Size > pc2.Y

	}
	return false
}
