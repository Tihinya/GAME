package engine

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
	mutex     sync.RWMutex
	positions map[*Entity]*PositionComponent
}
type MotionManager struct {
	mutex   sync.RWMutex
	motions map[*Entity]*MotionComponent
}
type SpriteManager struct {
	mutex   sync.RWMutex
	sprites map[*Entity]*SpriteComponent
}
type CollisionManager struct {
	mutex      sync.RWMutex
	collisions map[*Entity]*CollisionComponent
}
type HealthManager struct {
	mutex   sync.RWMutex
	healths map[*Entity]*HealthComponent
}
type InputManager struct {
	mutex  sync.RWMutex
	inputs map[*Entity]*InputComponent
}
type TimerManager struct {
	mutex  sync.RWMutex
	timers map[*Entity]*TimerComponent
}
type PowerUpManager struct {
	mutex    sync.RWMutex
	powerUps map[*Entity]*PowerUpComponent
}
type DamageManager struct {
	mutex   sync.RWMutex
	damages map[*Entity]*DamageComponent
}
type BombManager struct {
	mutex sync.RWMutex
	bombs map[*Entity]*BombComponent
}
type ExplosionManager struct {
	mutex      sync.RWMutex
	explosions map[*Entity]*ExplosionComponent
}
type BoxManager struct {
	mutex sync.RWMutex
	boxes map[*Entity]*BoxComponent
}

type WallManager struct {
	mutex sync.RWMutex
	walls map[*Entity]*WallComponent
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

func NewBoxManager() *BoxManager {
	return &BoxManager{
		boxes: make(map[*Entity]*BoxComponent),
	}
}

func NewWallManager() *WallManager {
	return &WallManager{
		walls: make(map[*Entity]*WallComponent),
	}
}

// --------------------------------
// Add Component functions
// --------------------------------

func (positionManager *PositionManager) AddComponent(entity *Entity, component *PositionComponent) {
	positionManager.mutex.Lock()
	defer positionManager.mutex.Unlock()
	positionManager.positions[entity] = component
}

func (spriteManager *SpriteManager) AddComponent(entity *Entity, component *SpriteComponent) {
	spriteManager.mutex.Lock()
	defer spriteManager.mutex.Unlock()
	spriteManager.sprites[entity] = component
}

func (motionManager *MotionManager) AddComponent(entity *Entity, component *MotionComponent) {
	motionManager.mutex.Lock()
	defer motionManager.mutex.Unlock()
	motionManager.motions[entity] = component
}

func (collisionManager *CollisionManager) AddComponent(entity *Entity, component *CollisionComponent) {
	collisionManager.mutex.Lock()
	defer collisionManager.mutex.Unlock()
	collisionManager.collisions[entity] = component
}

func (healthManager *HealthManager) AddComponent(entity *Entity, component *HealthComponent) {
	healthManager.mutex.Lock()
	defer healthManager.mutex.Unlock()
	healthManager.healths[entity] = component
}

func (inputManager *InputManager) AddComponent(entity *Entity, component *InputComponent) {
	inputManager.mutex.Lock()
	defer inputManager.mutex.Unlock()
	inputManager.inputs[entity] = component
}

func (timerManager *TimerManager) AddComponent(entity *Entity, component *TimerComponent) {
	timerManager.mutex.Lock()
	defer timerManager.mutex.Unlock()
	timerManager.timers[entity] = component
}

func (powerUpManager *PowerUpManager) AddComponent(entity *Entity, component *PowerUpComponent) {
	powerUpManager.mutex.Lock()
	defer powerUpManager.mutex.Unlock()
	powerUpManager.powerUps[entity] = component
}

func (damageManager *DamageManager) AddComponent(entity *Entity, component *DamageComponent) {
	damageManager.mutex.Lock()
	defer damageManager.mutex.Unlock()
	damageManager.damages[entity] = component
}

func (bombManager *BombManager) AddComponent(entity *Entity, component *BombComponent) {
	bombManager.mutex.Lock()
	defer bombManager.mutex.Unlock()
	bombManager.bombs[entity] = component
}

func (explosionManager *ExplosionManager) AddComponent(entity *Entity, component *ExplosionComponent) {
	explosionManager.mutex.Lock()
	defer explosionManager.mutex.Unlock()
	explosionManager.explosions[entity] = component
}

func (boxManager *BoxManager) AddComponent(entity *Entity, component *BoxComponent) {
	boxManager.mutex.Lock()
	defer boxManager.mutex.Unlock()
	boxManager.boxes[entity] = component
}

func (wallManager *WallManager) AddComponent(entity *Entity, component *WallComponent) {
	wallManager.mutex.Lock()
	defer wallManager.mutex.Unlock()
	wallManager.walls[entity] = component
}

// --------------------------------
// Delete Component functions
// --------------------------------

func (positionManager *PositionManager) DeleteComponent(entity *Entity) {
	positionManager.mutex.Lock()
	defer positionManager.mutex.Unlock()
	delete(positionManager.positions, entity)
}

func (spriteManager *SpriteManager) DeleteComponent(entity *Entity) {
	spriteManager.mutex.Lock()
	defer spriteManager.mutex.Unlock()
	delete(spriteManager.sprites, entity)
}

func (motionManager *MotionManager) DeleteComponent(entity *Entity) {
	motionManager.mutex.Lock()
	defer motionManager.mutex.Unlock()
	delete(motionManager.motions, entity)
}

func (collisionManager *CollisionManager) DeleteComponent(entity *Entity) {
	collisionManager.mutex.Lock()
	defer collisionManager.mutex.Unlock()
	delete(collisionManager.collisions, entity)
}

func (healthManager *HealthManager) DeleteComponent(entity *Entity) {
	healthManager.mutex.Lock()
	defer healthManager.mutex.Unlock()
	delete(healthManager.healths, entity)
}

func (inputManager *InputManager) DeleteComponent(entity *Entity) {
	inputManager.mutex.Lock()
	defer inputManager.mutex.Unlock()
	delete(inputManager.inputs, entity)
}

func (timerManager *TimerManager) DeleteComponent(entity *Entity) {
	timerManager.mutex.Lock()
	defer timerManager.mutex.Unlock()
	delete(timerManager.timers, entity)
}

func (powerUpManager *PowerUpManager) DeleteComponent(entity *Entity) {
	powerUpManager.mutex.Lock()
	defer powerUpManager.mutex.Unlock()
	delete(powerUpManager.powerUps, entity)
}

func (damageManager *DamageManager) DeleteComponent(entity *Entity) {
	damageManager.mutex.Lock()
	defer damageManager.mutex.Unlock()
	delete(damageManager.damages, entity)
}

func (bombManager *BombManager) DeleteComponent(entity *Entity) {
	bombManager.mutex.Lock()
	defer bombManager.mutex.Unlock()
	delete(bombManager.bombs, entity)
}

func (explosionManager *ExplosionManager) DeleteComponent(entity *Entity) {
	explosionManager.mutex.Lock()
	defer explosionManager.mutex.Unlock()
	delete(explosionManager.explosions, entity)
}

func (boxManager *BoxManager) DeleteComponent(entity *Entity) {
	boxManager.mutex.Lock()
	defer boxManager.mutex.Unlock()
	delete(boxManager.boxes, entity)
}

func (wallManager *WallManager) DeleteComponent(entity *Entity) {
	wallManager.mutex.Lock()
	defer boxManager.mutex.Unlock()
	delete(wallManager.walls, entity)
}

// --------------------------------
// Get manager components
// --------------------------------

// Use RLock and RUnlock for read access
func (m *PositionManager) GetPosition(entity *Entity) *PositionComponent {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.positions[entity]
}

func (m *TimerManager) GetTimer(entity *Entity) *TimerComponent {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.timers[entity]
}

func (m *ExplosionManager) GetExplosion(entity *Entity) *ExplosionComponent {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.explosions[entity]
}

func (m *BombManager) GetBomb(entity *Entity) *BombComponent {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.bombs[entity]
}

func (m *BoxManager) GetBox(entity *Entity) *BoxComponent {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.boxes[entity]
}

func (m *WallManager) GetWall(entity *Entity) *WallComponent {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.walls[entity]
}

// --------------------------------
// Set manager components
// --------------------------------

func (m *PositionManager) SetPosition(entity *Entity, position *PositionComponent) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.positions[entity] = position
}

func (m *ExplosionManager) SetExplosion(entity *Entity, explosion *ExplosionComponent) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.explosions[entity] = explosion
}

func (m *BombManager) SetBomb(entity *Entity, Bomb *BombComponent) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.bombs[entity] = Bomb
}

func (m *BoxManager) SetBox(entity *Entity, Box *BoxComponent) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.boxes[entity] = Box
}

// --------------------------------
// Other
// --------------------------------

func (em *EntityManager) CreateEntity() *Entity {
	em.mutex.Lock()
	defer em.mutex.Unlock()
	entity := &Entity{Id: em.Id}
	em.entities = append(em.entities, entity)
	em.Id++
	return entity
}

func DeleteAllEntityComponents(e *Entity) {
	positionManager.DeleteComponent(e)
	motionManager.DeleteComponent(e)
	collisionManager.DeleteComponent(e)
	healthManager.DeleteComponent(e)
	inputManager.DeleteComponent(e)
	timerManager.DeleteComponent(e)
	powerUpManager.DeleteComponent(e)
	damageManager.DeleteComponent(e)
	bombManager.DeleteComponent(e)
	explosionManager.DeleteComponent(e)
	boxManager.DeleteComponent(e)
}
