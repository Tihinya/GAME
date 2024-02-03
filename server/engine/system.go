package engine

import (
	"fmt"
	"time"

	"bomberman-dom/helpers"
	"bomberman-dom/models"
)

var (
	fuseTime    = time.Millisecond * 500
	explodeTime = time.Millisecond * 150
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
	motionSystem    = NewMotionSystem()
	inputSystem     = NewInputSystem()
	powerUpSystem   = NewPowerUpSystem()
	healthSystem    = NewHealthSystem()
	explosionSystem = NewExplosionSystem()
)

// var damageSystem = NewDamageSystem()
// var bombSystem = NewBombSystem()

// Solving import cycles the interface way!!
var broadcaster helpers.Broadcaster

func SetBroadcaster(b helpers.Broadcaster) {
	broadcaster = b
}

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
		broadcastMotion(pc.X, pc.Y, e)
	}
}

func (is *InputSystem) update(dt float64) {
	for e, ic := range is.manager.inputs {
		mc := motionManager.motions[e]
		if ic.Input[Up] {
			mc.Velocity.Y = -mc.Speed
		}
		if ic.Input[Down] {
			mc.Velocity.Y = mc.Speed
		}
		if ic.Input[Left] {
			mc.Velocity.X = -mc.Speed
		}
		if ic.Input[Right] {
			mc.Velocity.X = mc.Speed
		}
		if ic.Input[Space] {
			CreateBomb(e)
		}
	}
}

func (hs *HealthSystem) update(dt float64) {
	for e, hc := range hs.manager.healths {
		if hc.CurrentHealth <= 0 {
			socketId := userEntityManager.GetUserIdByEntity(e)
			broadcastPlayerHealth(socketId, hc.CurrentHealth)
			DeleteAllEntityComponents(e)
		}
		if hc.CurrentHealth > hc.MaxHealth {
			socketId := userEntityManager.GetUserIdByEntity(e)
			hc.CurrentHealth = hc.MaxHealth
			broadcastPlayerHealth(socketId, hc.CurrentHealth)
		}
	}
}

func (ex *ExplosionSystem) update(dt float64) {
	bombManager.mutex.RLock()
	for e := range bombManager.bombs {
		bombManager.mutex.RUnlock()

		bombTimer := timerManager.GetTimer(e)
		if bombTimer == nil {
			fmt.Println("No timer found for bomb", e.Id)
			continue // Skip if no timer is set for this bomb
		}
		if time.Now().After(bombTimer.Time) {
			SpreadExplosion(e)
			DeleteAllEntityComponents(e)
		}
		bombManager.mutex.RLock()
	}
	bombManager.mutex.RUnlock()

	explosionManager.mutex.RLock()
	for e2 := range explosionManager.explosions {
		explosionManager.mutex.RUnlock()

		explosionTimer := timerManager.GetTimer(e2)
		if explosionTimer == nil {
			fmt.Println("No timer found for explosion", e2.Id)
			continue
		}
		if time.Now().After(explosionTimer.Time) {
			broadcastDeleteExplosions(e2)
			DeleteAllEntityComponents(e2)
		}
		explosionManager.mutex.RLock()
	}
	explosionManager.mutex.RUnlock()
}

func ExplodeBox(pos *PositionComponent) {
	entityManager.mutex.RLock()
	defer entityManager.mutex.RUnlock()
	for _, e := range entityManager.entities {
		pc := positionManager.GetPosition(e)
		if pc != nil && boxManager.GetBox(e) != nil && ((pc.X == pos.X) && (pc.Y == pos.Y)) {
			broadcastObstacle(pos.X, pos.Y, "box", "delete")
			DeleteAllEntityComponents(e)
		}
	}
}

func HandleInput(input models.GameInput, playerId int) {
	ic := &InputComponent{Input: input.Keys}
	player := userEntityManager.GetUserEntity(playerId)
	inputManager.SetInputs(player.entity, ic)
}

func DetectCollision(e1 *Entity, ignoreList ...*Entity) bool {
	pc1 := positionManager.positions[e1]

	// collisionManager.mutex.RLock()

	for e2, cc := range collisionManager.collisions {
		if e1 == e2 || contains(ignoreList, e1) {
			continue
		}

		mc := motionManager.motions[e1]
		pc2 := positionManager.positions[e2]
		puc := powerUpManager.powerUps[e2]
		puc1 := powerUpManager.powerUps[e1]
		hc := healthManager.healths[e1]

		collides := pc1.X < pc2.X+pc2.Size && pc1.X+pc2.Size > pc2.X && pc1.Y < pc2.Y+pc2.Size && pc1.Y+pc1.Size > pc2.Y

		if cc.Enabled && collides {
			switch puc.Name {
			case PowerUpSpeed:
				mc.Speed += Speed
				DeleteAllEntityComponents(e2)
				broadcastPowerup(pc2.X, pc2.Y, 1, "delete")
			case PowerUpHealth:
				hc.CurrentHealth += Regeneration
				DeleteAllEntityComponents(e2)
				broadcastPowerup(pc2.X, pc2.Y, 3, "delete")
			case PowerUpBomb:
				puc1.ExtraBombs += Bomb
				DeleteAllEntityComponents(e2)
				broadcastPowerup(pc2.X, pc2.Y, 2, "delete")
			case PowerUpExplosion:
				puc1.ExtraExplosionRange += ExplosionRange
				DeleteAllEntityComponents(e2)
				broadcastPowerup(pc2.X, pc2.Y, 4, "delete")
			}
			return false
		}
		return collides
	}
	// collisionManager.mutex.RUnlock()
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
