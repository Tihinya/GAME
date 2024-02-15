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
type ExplosionStopperManager struct {
	explosionStoppers map[*Entity]*ExplosionStopperComponent
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

func NewExplosionStopperManager() *ExplosionStopperManager {
	return &ExplosionStopperManager{
		explosionStoppers: make(map[*Entity]*ExplosionStopperComponent),
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
	positionManager.positions[entity] = component
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

func (unbreakableManager *ExplosionStopperManager) AddComponent(entity *Entity, component *ExplosionStopperComponent) {
	unbreakableManager.explosionStoppers[entity] = component
}

func (UserEntityManager *UserEntityManager) AddComponent(userId int, component *UserEntityComponent) {
	UserEntityManager.users[userId] = component
}

// --------------------------------
// Delete Component functions
// --------------------------------

func (positionManager *PositionManager) DeleteComponent(entity *Entity) {
	delete(positionManager.positions, entity)
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

func (unbreakableManager *ExplosionStopperManager) DeleteComponent(entity *Entity) {
	delete(unbreakableManager.explosionStoppers, entity)
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

func (m *InputManager) GetInputs(entity *Entity) *InputComponent {
	return m.inputs[entity]
}

func (m *UserEntityManager) GetUserEntity(userId int) *UserEntityComponent {
	return m.users[userId]
}

// Not sure about this function, since reverse map lookups are inefficient,
// need to check in real-time if this lags the game or not

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
	explosionStopperManager.DeleteComponent(e)
}
