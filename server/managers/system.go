package managers

import (
	"time"
)

var fuseTime = time.Second * 500

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
var collisionManager = NewCollisionManager()
var powerUpManager = NewPowerUpManager()
var damageManager = NewDamageManager()
var healthManager = NewHealthManager()

// --------------------------------
// Createing Systems
// --------------------------------

var motionSystem = NewMotionSystem()
var inputSystem = NewInputSystem()
var powerUpSystem = NewPowerUpSystem()
var healthSystem = NewHealthSystem()

// var damageSystem = NewDamageSystem()
// var bombSystem = NewBombSystem()
// var explosionSystem = NewExplosionSystem()

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

		if DetectCollision(e) {
			pc.X -= mc.Velocity.X
			pc.Y -= mc.Velocity.Y

			mc.Velocity.X -= mc.Acceleration.X
			mc.Velocity.Y -= mc.Acceleration.Y
		}
	}
}

func (is *InputSystem) update(dt time.Time) {
	for e, ic := range is.manager.inputs {
		mc := motionManager.motions[e]
		bc, exists := bombManager.bombs[e]
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
		if exists && ic.Input[Space] && len(bc.PutedBomb) < bc.BombAmount {
			bomb := CreateBobm(e)
			if DetectCollision(bomb, e) {
				delete(positionManager.postions, bomb)
			}

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

func (ex *ExplosionSystem) update(dt time.Time) {
	for e := range bombManager.bombs {
		ec := explosionManager.explosions[e]
		tc := timerManager.timers[e]
		pc := positionManager.postions[e]
		if !tc.Time.Before(time.Now()) {
			continue
		}
		CreateExplosion(pc)
		for i := 1; i <= ec.Range; i++ {
			// var expolions []*PositionComponent
			// pcRight := &PositionComponent{pc.X + float64(i), pc.Y, pc.Size + float64(i)}
			// pcUp := &PositionComponent{pc.X, pc.Y + float64(i), pc.Size + float64(i)}
			// pcLeft := &PositionComponent{pc.X - float64(i), pc.Y, pc.Size + float64(i)}
			// pcDown := &PositionComponent{pc.X, pc.Y - float64(i), pc.Size + float64(i)}
			// expolionComponets := append(expolions, pcUp)
			// expolionComponets = append(expolions, pcDown)
			// expolionComponets = append(expolions, pcLeft)
			// expolionComponets = append(expolions, pcRight)
			// for i2, pc2 := range expolionComponets {
			// 	CreateExplosion(pc2)
			// 	for e2, hc := range collisionManager.collisions {
			// 		if DetectCollision(bomb, e) {
			// 			delete(positionManager.postions, bomb)
			// 		}

			// 	}
			// }

		}
	}
}

func DetectCollision(entity *Entity, ignoreList ...*Entity) bool {
	pc1 := positionManager.postions[entity]

	for _, e2 := range entityManager.entities {
		if entity == e2 && contains(ignoreList, entity) {
			continue
		}

		pc2 := positionManager.postions[e2]
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
