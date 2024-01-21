package internal

import "sync"

// --------------------------------
// Managers Structs
// --------------------------------

type EntityManager struct {
	entities []*Entity
	Id       int
	mutex    sync.RWMutex
}
type PositionManager struct {
	positions map[*Entity]*PositionComponent
	mutex     sync.RWMutex
}
type MotionManager struct {
	motions map[*Entity]*MotionComponent
}
type SpriteManager struct {
	sprites map[*Entity]*SpriteComponent
}
type CollisionManager struct {
	collisions map[*Entity]*CollisionComponent
}
type HealthManager struct {
	healths map[*Entity]*HealthComponent
}
type InputManager struct {
	inputs map[*Entity]*InputComponent
}
type TimerManager struct {
	timers map[*Entity]*TimerComponent
}
type PowerUpManager struct {
	powerUps map[*Entity]*PowerUpComponent
}
type DamageManager struct {
	damages map[*Entity]*DamageComponent
}
type BombManager struct {
	bombs map[*Entity]*BombComponent
}
type ExplosionManager struct {
	explosions map[*Entity]*ExplosionComponent
	mutex      sync.RWMutex
}

// --------------------------------
// Systems
// --------------------------------
type PositionSystem struct {
	manager *PositionManager
}

type MotionSystem struct {
	manager *MotionManager
}

type SpriteSystem struct {
	manager *SpriteManager
}

type CollisionSystem struct {
	manager *CollisionManager
}

type HealthSystem struct {
	manager *HealthManager
}

type InputSystem struct {
	manager *InputManager
}

type TimerSystem struct {
	manager *TimerManager
}

type PowerUpSystem struct {
	manager *PowerUpManager
}

type DamageSystem struct {
	manager *DamageManager
}

type BombSystem struct {
	manager *BombManager
}

type ExplosionSystem struct {
	manager *ExplosionManager
}

// --------------------------------
// Createing systems
// --------------------------------

func NewMotionSystem() *MotionSystem {
	return &MotionSystem{
		manager: motionManager,
	}
}

func NewHealthSystem() *HealthSystem {
	return &HealthSystem{
		manager: healthManager,
	}
}

func NewInputSystem() *InputSystem {
	return &InputSystem{
		manager: inputManager,
	}
}

func NewPowerUpSystem() *PowerUpSystem {
	return &PowerUpSystem{
		manager: powerUpManager,
	}
}

func NewDamageSystem() *DamageSystem {
	return &DamageSystem{
		manager: damageManager,
	}
}

func NewBombSystem() *BombSystem {
	return &BombSystem{
		manager: bombManager,
	}
}

func NewExplosionSystem() *ExplosionSystem {
	return &ExplosionSystem{
		manager: explosionManager,
	}
}

// --------------------------------
// Createing managers
// --------------------------------

func NewPositionManager() *PositionManager {
	return &PositionManager{
		positions: make(map[*Entity]*PositionComponent),
	}
}

func NewSpriteManager() *SpriteManager {
	return &SpriteManager{
		sprites: make(map[*Entity]*SpriteComponent),
	}
}

func NewMotionManager() *MotionManager {
	return &MotionManager{
		motions: make(map[*Entity]*MotionComponent),
	}
}

func NewCollisionManager() *CollisionManager {
	return &CollisionManager{
		collisions: make(map[*Entity]*CollisionComponent),
	}
}

func NewHealthManager() *HealthManager {
	return &HealthManager{
		healths: make(map[*Entity]*HealthComponent),
	}
}

func NewInputManager() *InputManager {
	return &InputManager{
		inputs: make(map[*Entity]*InputComponent),
	}
}

func NewTimerManager() *TimerManager {
	return &TimerManager{
		timers: make(map[*Entity]*TimerComponent),
	}
}

func NewPowerUpManager() *PowerUpManager {
	return &PowerUpManager{
		powerUps: make(map[*Entity]*PowerUpComponent),
	}
}

func NewDamageManager() *DamageManager {
	return &DamageManager{
		damages: make(map[*Entity]*DamageComponent),
	}
}

func NewBombManager() *BombManager {
	return &BombManager{
		bombs: make(map[*Entity]*BombComponent),
	}
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		entities: make([]*Entity, 0),
		Id:       1,
	}
}

func NewExplosionManager() *ExplosionManager {
	return &ExplosionManager{
		explosions: make(map[*Entity]*ExplosionComponent),
	}
}

func (positionManager *PositionManager) AddComponent(entity *Entity, component *PositionComponent) {
	positionManager.positions[entity] = component
}

func (spriteManager *SpriteManager) AddComponent(entity *Entity, component *SpriteComponent) {
	spriteManager.sprites[entity] = component
}

func (motionManager *MotionManager) AddComponent(entity *Entity, component *MotionComponent) {
	motionManager.motions[entity] = component
}

func (collisionManager *CollisionManager) AddComponent(entity *Entity, component *CollisionComponent) {
	collisionManager.collisions[entity] = component
}

func (healthManager *HealthManager) AddComponent(entity *Entity, component *HealthComponent) {
	healthManager.healths[entity] = component
}

func (inputManager *InputManager) AddComponent(entity *Entity, component *InputComponent) {
	inputManager.inputs[entity] = component
}

func (timerManager *TimerManager) AddComponent(entity *Entity, component *TimerComponent) {
	timerManager.timers[entity] = component
}

func (powerUpManager *PowerUpManager) AddComponent(entity *Entity, component *PowerUpComponent) {
	powerUpManager.powerUps[entity] = component
}

func (damageManager *DamageManager) AddComponent(entity *Entity, component *DamageComponent) {
	damageManager.damages[entity] = component
}

func (bombManager *BombManager) AddComponent(entity *Entity, component *BombComponent) {
	bombManager.bombs[entity] = component
}

func (explosionManager *ExplosionManager) AddComponent(entity *Entity, component *ExplosionComponent) {
	explosionManager.explosions[entity] = component
}

func (em *EntityManager) CreateEntity() *Entity {
	em.mutex.Lock()
	defer em.mutex.Unlock()
	entity := &Entity{Id: em.Id}
	em.entities = append(em.entities, entity)
	em.Id++
	return entity
}

// Use RLock and RUnlock for read access
func (m *PositionManager) GetPosition(entity *Entity) *PositionComponent {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.positions[entity]
}

// Use Lock and Unlock for write access
func (m *PositionManager) SetPosition(entity *Entity, position *PositionComponent) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.positions[entity] = position
}

func (m *ExplosionManager) GetExplosion(entity *Entity) *ExplosionComponent {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.explosions[entity]
}

func (m *ExplosionManager) SetExplosion(entity *Entity, explosion *ExplosionComponent) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.explosions[entity] = explosion
}

func DeleteAllEntityComponents(e *Entity) {
	delete(positionManager.positions, e)
	delete(motionManager.motions, e)
	delete(collisionManager.collisions, e)
	delete(healthManager.healths, e)
	delete(inputManager.inputs, e)
	delete(timerManager.timers, e)
	delete(powerUpManager.powerUps, e)
	delete(damageManager.damages, e)
	delete(bombManager.bombs, e)
	delete(explosionManager.explosions, e)
}
