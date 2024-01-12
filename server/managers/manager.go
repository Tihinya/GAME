package managers

// --------------------------------
// Managers Structs
// --------------------------------

type EntityManager struct {
	entities []*Entity
	Id       int
}
type PositionManager struct {
	postions map[*Entity]*PositionComponent
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
	powerUps map[*Entity][]*PowerUpComponent
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

func NewPositionSystem() *PositionSystem {
	return &PositionSystem{
		manager: positionManager,
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

func NewTimerSystem() *TimerSystem {
	return &TimerSystem{
		manager: timerManager,
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
		postions: make(map[*Entity]*PositionComponent),
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
		powerUps: make(map[*Entity][]*PowerUpComponent),
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

func (positionManager *PositionManager) AddComponet(entity *Entity, component *PositionComponent) {
	positionManager.postions[entity] = component
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

func (inputManager *InputManager) AddComponet(entity *Entity, component *InputComponent) {
	inputManager.inputs[entity] = component
}

func (timerManager *TimerManager) AddComponet(entity *Entity, component *TimerComponent) {
	timerManager.timers[entity] = component
}

func (powerUpManager *PowerUpManager) AddComponet(entity *Entity, component *PowerUpComponent) {
	powerUpManager.powerUps[entity] = append(powerUpManager.powerUps[entity], component)
}

func (damageManager *DamageManager) AddComponet(entity *Entity, component *DamageComponent) {
	damageManager.damages[entity] = component
}

func (bombManager *BombManager) AddComponet(entity *Entity, component *BombComponent) {
	bombManager.bombs[entity] = component
}

func (explosionManager *ExplosionManager) AddComponet(entity *Entity, component *ExplosionComponent) {
	explosionManager.explosions[entity] = component
}

func (em *EntityManager) CreateEntity() *Entity {
	entity := &Entity{Id: em.Id}
	em.entities = append(em.entities, entity)
	em.Id++
	return entity
}

func DeleteAllEntityComponents(e *Entity) {
	if _, exists := positionManager.postions[e]; exists {
		delete(positionManager.postions, e)
	}

	if _, exists := motionManager.motions[e]; exists {
		delete(motionManager.motions, e)
	}

	if _, exists := collectionManager.collisions[e]; exists {
		delete(collectionManager.collisions, e)
	}

	if _, exists := healthManager.healths[e]; exists {
		delete(healthManager.healths, e)
	}

	if _, exists := inputManager.inputs[e]; exists {
		delete(inputManager.inputs, e)
	}

	if _, exists := timerManager.timers[e]; exists {
		delete(timerManager.timers, e)
	}

	if _, exists := powerUpManager.powerUps[e]; exists {
		delete(powerUpManager.powerUps, e)
	}

	if _, exists := damageManager.damages[e]; exists {
		delete(damageManager.damages, e)
	}

	if _, exists := bombManager.bombs[e]; exists {
		delete(bombManager.bombs, e)
	}

	if _, exists := explosionManager.explosions[e]; exists {
		delete(explosionManager.explosions, e)
	}
}
