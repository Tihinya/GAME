package internal

import "time"

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
	componentSize         = 1.0
	playerMaxHealth       = 3
	boxHealth             = 1
	defaultExplosionRange = 6
	explosionDamage       = 1
)

const (
	LeftSprite  = "leftSprite.png"
	RightSprite = "rightSprite.png"
	UpSprite    = "upSprite.png"
	DownSprite  = "downSprite.png"
)

var (
	player1SpawnX = 10.0
	player1Spawny = 5.0
)

func CreatePlayer() *Entity {
	player := entityManager.CreateEntity()

	playerPosition := &PositionComponent{X: player1SpawnX, Y: player1Spawny, Size: componentSize}
	playerMotion := &MotionComponent{Velocity: Vec2{X: 0, Y: 0}, Acceleration: Vec2{X: 0, Y: 0}}
	playerInput := &InputComponent{Input: map[string]bool{}}
	playerHealth := &HealthComponent{CurrentHealth: playerMaxHealth, MaxHealth: playerMaxHealth}
	playerExposion := &ExplosionComponent{Range: defaultExplosionRange}

	positionManager.AddComponet(player, playerPosition)
	motionManager.AddComponent(player, playerMotion)
	inputManager.AddComponet(player, playerInput)
	healthManager.AddComponent(player, playerHealth)
	explosionManager.AddComponet(player, playerExposion)

	return player
}

func CreateBobm(player *Entity) *Entity {
	pc := positionManager.postions[player]
	ec := explosionManager.explosions[player]

	bomb := entityManager.CreateEntity()

	bombPosition := &PositionComponent{X: pc.X, Y: pc.Y, Size: pc.Size}
	bombCollision := &CollisionComponent{Enabled: false}
	bombTimer := &TimerComponent{time.Now().Add(fuseTime)}
	bombExplosion := &ExplosionComponent{Range: ec.Range}

	collisionManager.AddComponent(bomb, bombCollision)
	timerManager.AddComponet(bomb, bombTimer)
	positionManager.AddComponet(bomb, bombPosition)
	explosionManager.AddComponet(bomb, bombExplosion)

	return bomb
}

func CreateExplosion(positionComponent *PositionComponent) *Entity {
	explosion := entityManager.CreateEntity()

	explosionPosition := positionComponent
	explosionTimer := &TimerComponent{time.Now().Add(fuseTime)}
	explosionDamage := &DamageComponent{DamageAmount: explosionDamage}
	timerManager.AddComponet(explosion, explosionTimer)
	positionManager.AddComponet(explosion, explosionPosition)
	damageManager.AddComponet(explosion, explosionDamage)

	return explosion
}

func CreatePowerUp(powerUpName string) *Entity {
	powerUp := entityManager.CreateEntity()

	powerUpPosition := &PositionComponent{}
	powerUpProperty := &PowerUpComponent{Name: powerUpName}
	positionManager.AddComponet(powerUp, powerUpPosition)
	powerUpManager.AddComponet(powerUp, powerUpProperty)

	return powerUp

}

func CreateWall() *Entity {
	wall := entityManager.CreateEntity()

	playerPosition := &PositionComponent{}

	positionManager.AddComponet(wall, playerPosition)
	return wall
}

func CreateBox() *Entity {
	box := entityManager.CreateEntity()

	playerPosition := &PositionComponent{}
	playerHealth := &HealthComponent{CurrentHealth: boxHealth, MaxHealth: boxHealth}

	positionManager.AddComponet(box, playerPosition)
	healthManager.AddComponent(box, playerHealth)

	return box
}
