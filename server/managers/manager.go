package managers

import (
	cm "bomberman-dom/components"
	en "bomberman-dom/entities"
)

// --------------------------------
// Managers Structs
// --------------------------------
type EntityManager struct {
	entities []*en.Entity
	Id       int
}
type PositionManager struct {
	postions map[*en.Entity]*cm.PositionComponent
}
type MotionManager struct {
	motions map[*en.Entity]*cm.MotionComponent
}
type SpriteManager struct {
	sprites map[*en.Entity]*cm.SpriteComponent
}
type CollisionManager struct {
	collisions map[*en.Entity]*cm.CollisionComponent
}
type HealthManager struct {
	healths map[*en.Entity]*cm.HealthComponent
}
type InputManager struct {
	inputs map[*en.Entity]*cm.InputComponent
}
type TimerManager struct {
	timers map[*en.Entity]*cm.TimerComponent
}
type PowerUpManager struct {
	powerUps map[*en.Entity]*cm.PowerUpComponent
}
type DamageManager struct {
	damages map[*en.Entity]*cm.DamageComponent
}

// --------------------------------
// Createing managers
// --------------------------------
func NewEntityManager() *EntityManager {
	return &EntityManager{
		entities: make([]*en.Entity, 0),
		Id:       1,
	}
}
func NewPositionManager() *PositionManager {
	return &PositionManager{
		postions: make(map[*en.Entity]*cm.PositionComponent),
	}
}
func NewSpriteManager() *SpriteManager {
	return &SpriteManager{
		sprites: make(map[*en.Entity]*cm.SpriteComponent),
	}
}
func NewMotionManager() *MotionManager {
	return &MotionManager{
		motions: make(map[*en.Entity]*cm.MotionComponent),
	}
}
func NewCollisionManager() *CollisionManager {
	return &CollisionManager{
		collisions: make(map[*en.Entity]*cm.CollisionComponent),
	}
}
func NewHealthManager() *HealthManager {
	return &HealthManager{
		healths: make(map[*en.Entity]*cm.HealthComponent),
	}
}
func NewInputManager() *InputManager {
	return &InputManager{
		inputs: make(map[*en.Entity]*cm.InputComponent),
	}
}
func NewTimerManager() *TimerManager {
	return &TimerManager{
		timers: make(map[*en.Entity]*cm.TimerComponent),
	}
}
func NewPowerUpManager() *PowerUpManager {
	return &PowerUpManager{
		powerUps: make(map[*en.Entity]*cm.PowerUpComponent),
	}
}
func NewDamageManager() *DamageManager {
	return &DamageManager{
		damages: make(map[*en.Entity]*cm.DamageComponent),
	}
}

// --------------------------------
// CAdd commponets
// --------------------------------

func (positionManager *PositionManager) AddComponet(entity *en.Entity, component *cm.PositionComponent) {
	positionManager.postions[entity] = component
}

func (spriteManager *SpriteManager) AddComponent(entity *en.Entity, component *cm.SpriteComponent) {
	spriteManager.sprites[entity] = component
}

func (motionManager *MotionManager) AddComponent(entity *en.Entity, component *cm.MotionComponent) {
	motionManager.motions[entity] = component
}

func (collisionManager *CollisionManager) AddComponent(entity *en.Entity, component *cm.CollisionComponent) {
	collisionManager.collisions[entity] = component
}

func (healthManager *HealthManager) AddComponent(entity *en.Entity, component *cm.HealthComponent) {
	healthManager.healths[entity] = component
}

func (inputManager *InputManager) AddComponet(entity *en.Entity, component *cm.InputComponent) {
	inputManager.inputs[entity] = component
}

func (timerManager *TimerManager) AddComponet(entity *en.Entity, component *cm.TimerComponent) {
	timerManager.timers[entity] = component
}

func (powerUpManager *PowerUpManager) AddComponet(entity *en.Entity, component *cm.PowerUpComponent) {
	powerUpManager.powerUps[entity] = component
}

func (damageManager *DamageManager) AddComponet(entity *en.Entity, component *cm.DamageComponent) {
	damageManager.damages[entity] = component
}
