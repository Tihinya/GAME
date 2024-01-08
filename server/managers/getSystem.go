package managers

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
	LeftSprite  = "leftSprite.png"
	RightSprite = "rightSprite.png"
	UpSprite    = "upSprite.png"
	DownSprite  = "downSprite.png"
)

func (entity *Entity) getPosition(systemMangers *SystemManagers) *PositionComponent {
	return systemMangers.PositionManager.postions[entity]
}
func (entity *Entity) getMotion(systemMangers *SystemManagers) *MotionComponent {
	return systemMangers.MotionManager.motions[entity]
}
func (entity *Entity) getInput(systemMangers *SystemManagers) *InputComponent {
	return systemMangers.InputManager.inputs[entity]
}
func (entity *Entity) getCollision(systemMangers *SystemManagers) *CollisionComponent {
	return systemMangers.CollisionManager.collisions[entity]
}
func (entity *Entity) getSprite(systemMangers *SystemManagers) *SpriteComponent {
	return systemMangers.SpriteManager.sprites[entity]
}
func (entity *Entity) getHealth(systemMangers *SystemManagers) *HealthComponent {
	return systemMangers.HealthManager.healths[entity]
}
func (entity *Entity) getDamage(systemMangers *SystemManagers) *DamageComponent {
	return systemMangers.DamageManager.damages[entity]
}
func (entity *Entity) getTimer(systemMangers *SystemManagers) *TimerComponent {
	return systemMangers.TimerManager.timers[entity]
}
func (entity *Entity) getPowerUp(systemMangers *SystemManagers) []*PowerUpComponent {
	return systemMangers.PowerUpManager.powerUps[entity]
}
func (entity *Entity) getBomb(systemMangers *SystemManagers) *BombComponent {
	return systemMangers.BombManager.bombs[entity]
}
func (entity *Entity) getExplosion(systemMangers *SystemManagers) *ExplosionComponent {
	return systemMangers.ExplosionManager.explosions[entity]
}
