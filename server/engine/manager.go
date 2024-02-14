package engine

// --------------------------------
// Managers Structs
// --------------------------------

type EntityManager struct {
	entities []*Entity
	Id       int
}
type PositionManager struct {
	positions map[*Entity]*PositionComponent
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
}
type BoxManager struct {
	boxes map[*Entity]*BoxComponent
}

type WallManager struct {
	walls map[*Entity]*WallComponent
}
type UserEntityManager struct {
	users map[int]*UserEntityComponent
}

// --------------------------------
// Systems
// --------------------------------

type MotionSystem struct {
	manager *MotionManager
}

type HealthSystem struct {
	manager *HealthManager
}

type InputSystem struct {
	manager *InputManager
}

type ExplosionSystem struct {
	manager *ExplosionManager
}

type DamageSystem struct {
	manager *DamageManager
}
type PowerUpSystem struct {
	manager *PowerUpManager
}

// --------------------------------
// Createing systems
// --------------------------------

func NewDamageSystem() *DamageSystem {
	return &DamageSystem{
		manager: damageManager,
	}
}
func NewPowerUpSystem() *PowerUpSystem {
	return &PowerUpSystem{
		manager: powerUpManager,
	}
}

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

func NewUserEntityManager() *UserEntityManager {
	return &UserEntityManager{
		users: make(map[int]*UserEntityComponent),
	}
}

// --------------------------------
// Add Component functions
// --------------------------------

func (positionManager *PositionManager) AddComponent(entity *Entity, component *PositionComponent) {
	// positionManager.mutex.Lock()
	// defer positionManager.mutex.Unlock()
	positionManager.positions[entity] = component
}

func (spriteManager *SpriteManager) AddComponent(entity *Entity, component *SpriteComponent) {
	// spriteManager.mutex.Lock()
	// defer spriteManager.mutex.Unlock()
	spriteManager.sprites[entity] = component
}

func (motionManager *MotionManager) AddComponent(entity *Entity, component *MotionComponent) {
	// motionManager.mutex.Lock()
	// defer motionManager.mutex.Unlock()
	motionManager.motions[entity] = component
}

func (collisionManager *CollisionManager) AddComponent(entity *Entity, component *CollisionComponent) {
	// collisionManager.mutex.Lock()
	// defer collisionManager.mutex.Unlock()
	collisionManager.collisions[entity] = component
}

func (healthManager *HealthManager) AddComponent(entity *Entity, component *HealthComponent) {
	// healthManager.mutex.Lock()
	// defer healthManager.mutex.Unlock()
	healthManager.healths[entity] = component
}

func (inputManager *InputManager) AddComponent(entity *Entity, component *InputComponent) {
	// inputManager.mutex.Lock()
	// defer inputManager.mutex.Unlock()
	inputManager.inputs[entity] = component
}

func (timerManager *TimerManager) AddComponent(entity *Entity, component *TimerComponent) {
	// timerManager.mutex.Lock()
	// defer timerManager.mutex.Unlock()
	timerManager.timers[entity] = component
}

func (powerUpManager *PowerUpManager) AddComponent(entity *Entity, component *PowerUpComponent) {
	// powerUpManager.mutex.Lock()
	// defer powerUpManager.mutex.Unlock()
	powerUpManager.powerUps[entity] = component
}

func (damageManager *DamageManager) AddComponent(entity *Entity, component *DamageComponent) {
	// damageManager.mutex.Lock()
	// defer damageManager.mutex.Unlock()
	damageManager.damages[entity] = component
}

func (bombManager *BombManager) AddComponent(entity *Entity, component *BombComponent) {
	// bombManager.mutex.Lock()
	// defer bombManager.mutex.Unlock()
	bombManager.bombs[entity] = component
}

func (explosionManager *ExplosionManager) AddComponent(entity *Entity, component *ExplosionComponent) {
	// explosionManager.mutex.Lock()
	// defer explosionManager.mutex.Unlock()
	explosionManager.explosions[entity] = component
}

func (boxManager *BoxManager) AddComponent(entity *Entity, component *BoxComponent) {
	// boxManager.mutex.Lock()
	// defer boxManager.mutex.Unlock()
	boxManager.boxes[entity] = component
}

func (wallManager *WallManager) AddComponent(entity *Entity, component *WallComponent) {
	// wallManager.mutex.Lock()
	// defer wallManager.mutex.Unlock()
	wallManager.walls[entity] = component
}

func (UserEntityManager *UserEntityManager) AddComponent(userId int, component *UserEntityComponent) {
	// UserEntityManager.mutex.Lock()
	// defer UserEntityManager.mutex.Unlock()
	UserEntityManager.users[userId] = component
}

// --------------------------------
// Delete Component functions
// --------------------------------

func (positionManager *PositionManager) DeleteComponent(entity *Entity) {
	delete(positionManager.positions, entity)
}

func (spriteManager *SpriteManager) DeleteComponent(entity *Entity) {
	delete(spriteManager.sprites, entity)
}

func (motionManager *MotionManager) DeleteComponent(entity *Entity) {
	delete(motionManager.motions, entity)
}

func (collisionManager *CollisionManager) DeleteComponent(entity *Entity) {
	delete(collisionManager.collisions, entity)
}

func (healthManager *HealthManager) DeleteComponent(entity *Entity) {
	delete(healthManager.healths, entity)
}

func (inputManager *InputManager) DeleteComponent(entity *Entity) {
	delete(inputManager.inputs, entity)
}

func (timerManager *TimerManager) DeleteComponent(entity *Entity) {
	delete(timerManager.timers, entity)
}

func (powerUpManager *PowerUpManager) DeleteComponent(entity *Entity) {
	delete(powerUpManager.powerUps, entity)
}

func (damageManager *DamageManager) DeleteComponent(entity *Entity) {
	delete(damageManager.damages, entity)
}

func (bombManager *BombManager) DeleteComponent(entity *Entity) {
	delete(bombManager.bombs, entity)
}

func (explosionManager *ExplosionManager) DeleteComponent(entity *Entity) {
	delete(explosionManager.explosions, entity)
}

func (boxManager *BoxManager) DeleteComponent(entity *Entity) {
	delete(boxManager.boxes, entity)
}

func (wallManager *WallManager) DeleteComponent(entity *Entity) {
	delete(wallManager.walls, entity)
}

func (UserEntityManager *UserEntityManager) DeleteComponent(userId int) {
	delete(UserEntityManager.users, userId)
}

// --------------------------------
// Get manager components
// --------------------------------

// Use RLock and RUnlock for read access
func (m *PositionManager) GetPosition(entity *Entity) *PositionComponent {
	return m.positions[entity]
}

func (m *TimerManager) GetTimer(entity *Entity) *TimerComponent {
	return m.timers[entity]
}

func (m *ExplosionManager) GetExplosion(entity *Entity) *ExplosionComponent {
	return m.explosions[entity]
}

func (m *BombManager) GetBomb(entity *Entity) *BombComponent {
	return m.bombs[entity]
}

func (m *BoxManager) GetBox(entity *Entity) *BoxComponent {
	return m.boxes[entity]
}

func (m *WallManager) GetWall(entity *Entity) *WallComponent {
	return m.walls[entity]
}

func (m *InputManager) GetInputs(entity *Entity) *InputComponent {
	return m.inputs[entity]
}

func (m *UserEntityManager) GetUserEntity(userId int) *UserEntityComponent {
	return m.users[userId]
}

// Not sure about this function, since reverse map lookups are inefficient,
// need to check in real-time if this lags the game or not
func (m *UserEntityManager) GetUserIdByEntity(entity *Entity) int {
	for userId, ent := range m.users {
		if ent.entity == entity {
			return userId
		}
	}
	return 0
}

// --------------------------------
// Set manager components
// --------------------------------

func (m *PositionManager) SetPosition(entity *Entity, position *PositionComponent) {
	m.positions[entity] = position
}

func (m *ExplosionManager) SetExplosion(entity *Entity, explosion *ExplosionComponent) {
	m.explosions[entity] = explosion
}

func (m *BombManager) SetBomb(entity *Entity, Bomb *BombComponent) {
	m.bombs[entity] = Bomb
}

func (m *BoxManager) SetBox(entity *Entity, Box *BoxComponent) {
	m.boxes[entity] = Box
}

func (m *InputManager) SetInputs(entity *Entity, inputs *InputComponent) {
	m.inputs[entity] = inputs
}

func (m *UserEntityManager) SetUserEntity(userId int, user *UserEntityComponent) {
	m.users[userId] = user
}

// --------------------------------
// Other
// --------------------------------

func (em *EntityManager) CreateEntity(name string) *Entity {
	entity := &Entity{Id: em.Id, Name: name}

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
