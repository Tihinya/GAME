package managers

// --------------------------------
// Managers Structs
// --------------------------------

type SystemManagers struct {
	PositionManager  *PositionManager
	MotionManager    *MotionManager
	InputManager     *InputManager
	CollisionManager *CollisionManager
	SpriteManager    *SpriteManager
	HealthManager    *HealthManager
	PowerUpManager   *PowerUpManager
	TimerManager     *TimerManager
	DamageManager    *DamageManager
	BombManager      *BombManager
}

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

// --------------------------------
// Position
// --------------------------------

func (positionManager *PositionManager) AddComponet(entity *Entity, component *PositionComponent) {
	positionManager.postions[entity] = component
}

func (positionManager *PositionManager) DeleteComponet(entity *Entity, component *PositionComponent) {
	delete(positionManager.postions, entity)
}

// --------------------------------
// Sprite
// --------------------------------

func (spriteManager *SpriteManager) AddComponent(entity *Entity, component *SpriteComponent) {
	spriteManager.sprites[entity] = component
}

func (spriteManager *SpriteManager) DeleteComponet(entity *Entity, component *SpriteComponent) {
	delete(spriteManager.sprites, entity)
}

// --------------------------------
// Motion
// --------------------------------

func (motionManager *MotionManager) AddComponent(entity *Entity, component *MotionComponent) {
	motionManager.motions[entity] = component
}

func (motionManager *MotionManager) DeleteComponet(entity *Entity, component *MotionComponent) {
	delete(motionManager.motions, entity)
}

// --------------------------------
// Collision
// --------------------------------

func (collisionManager *CollisionManager) AddComponent(entity *Entity, component *CollisionComponent) {
	collisionManager.collisions[entity] = component
}

func (collisionManager *CollisionManager) DeleteComponet(entity *Entity, component *CollisionComponent) {
	delete(collisionManager.collisions, entity)
}

// --------------------------------
// Health
// --------------------------------

func (healthManager *HealthManager) AddComponent(entity *Entity, component *HealthComponent) {
	healthManager.healths[entity] = component
}

func (healthManager *HealthManager) DeleteComponet(entity *Entity, component *HealthComponent) {
	delete(healthManager.healths, entity)
}

// --------------------------------
// Input
// --------------------------------

func (inputManager *InputManager) AddComponet(entity *Entity, component *InputComponent) {
	inputManager.inputs[entity] = component
}

func (inputManager *InputManager) DeleteComponet(entity *Entity, component *InputComponent) {
	delete(inputManager.inputs, entity)
}

// --------------------------------
// Timer
// --------------------------------

func (timerManager *TimerManager) AddComponet(entity *Entity, component *TimerComponent) {
	timerManager.timers[entity] = component
}

func (timerManager *TimerManager) DeleteComponet(entity *Entity, component *TimerComponent) {
	delete(timerManager.timers, entity)
}

// --------------------------------
// PowerUp
// --------------------------------

func (powerUpManager *PowerUpManager) AddComponet(entity *Entity, component *PowerUpComponent) {
	powerUpManager.powerUps[entity] = append(powerUpManager.powerUps[entity], component)
}

func (powerUpManager *PowerUpManager) DeleteComponet(entity *Entity, component *PowerUpComponent) {
	delete(powerUpManager.powerUps, entity)
}

// --------------------------------
// Damage
// --------------------------------

func (damageManager *DamageManager) AddComponet(entity *Entity, component *DamageComponent) {
	damageManager.damages[entity] = component
}

func (damageManager *DamageManager) DeleteComponet(entity *Entity, component *DamageComponent) {
	delete(damageManager.damages, entity)
}

// --------------------------------
// Bomb
// --------------------------------

func (bombManager *BombManager) AddComponet(entity *Entity, component *BombComponent) {
	bombManager.bombs[entity] = component
}

func (bombManager *BombManager) DeleteComponet(entity *Entity, component *BombComponent) {
	delete(bombManager.bombs, entity)
}

func (em *EntityManager) CreateEntity() *Entity {
	entity := &Entity{Id: em.Id}
	em.entities = append(em.entities, entity)
	em.Id++
	return entity
}
