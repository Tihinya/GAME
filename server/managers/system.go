package managers

import (
	"fmt"
	"time"
)

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

func movementSystem(position *PositionComponent, motion *MotionComponent) {
	position.X += motion.Velocity.X
	position.Y += motion.Velocity.Y

	motion.Velocity.X += motion.Acceleration.X
	motion.Velocity.Y += motion.Acceleration.Y
}

func inputSystem(input *InputComponent, motion *MotionComponent) {
	if input.Input[Up] {
		motion.Velocity.Y = -Speed
	}
	if input.Input[Down] {
		motion.Velocity.Y = Speed
	}
	if input.Input[Left] {
		motion.Velocity.X = -Speed
	}
	if input.Input[Right] {
		motion.Velocity.X = Speed
	}
}
func checkCollision(x1, y1, x2, y2, size1, size2 float64) bool {
	return x1 < x2+size2 && x1+size1 > x2 && y1 < y2+size2 && y1+size1 > y2
}

func collisionSystem(entityManager *EntityManager, entity1 *Entity, entity1Position *PositionComponent, entity1Collision *CollisionComponent, system *SystemManagers, powerUps []*PowerUpComponent, health *HealthComponent) {
	for _, entity2 := range entityManager.entities {
		if entity1 != entity2 {
			entity2Position := entity2.getPosition(system)
			entity2Collision := entity2.getCollision(system)
			entity2PowerUp := entity2.getPowerUp(system)
			entity2Damage := entity2.getDamage(system)

			if !checkCollision(entity1Position.X, entity1Position.Y, entity2Position.X, entity2Position.Y, entity1Collision.Size, entity2Collision.Size) {
				continue
			}
			if entity2Collision.Enabled && entity2PowerUp != nil {
				powerUps = append(powerUps, entity2PowerUp...)
			}
			if entity2Collision.Enabled && entity2Damage != nil {
				health.CurrentHealth -= entity2Damage.DamageAmount
			}
		}
	}
}

func renderSystem(input *InputComponent, sprite *SpriteComponent) {
	if input.Input[Up] {
		sprite.Texture = UpSprite
	}
	if input.Input[Down] {
		sprite.Texture = DownSprite
	}
	if input.Input[Left] {
		sprite.Texture = LeftSprite
	}
	if input.Input[Right] {
		sprite.Texture = RightSprite
	} else {
		sprite.Texture = UpSprite
	}
}

func healthSystem(health *HealthComponent, entity *Entity, system *SystemManagers, sprite *SpriteComponent, input *InputComponent, position *PositionComponent, collision *CollisionComponent) {
	if health.CurrentHealth <= 0 {
		system.HealthManager.DeleteComponet(entity, health)
		system.SpriteManager.DeleteComponet(entity, sprite)
		system.InputManager.DeleteComponet(entity, input)
		system.PositionManager.DeleteComponet(entity, position)
		system.CollisionManager.DeleteComponet(entity, collision)
	}
}

func powerUpSystem(position *PositionComponent, powerUps []*PowerUpComponent, motion *MotionComponent, health *HealthComponent, bomb *BombComponent) {
	motion.Acceleration.X = 0
	for _, powerUp := range powerUps {
		if powerUp.Name == "speed" {
			motion.Acceleration.X += Acceleration
			motion.Acceleration.Y += Acceleration
		}
		if powerUp.Name == "health" {
			health.CurrentHealth += Regeneration
		}
		if powerUp.Name == "bomb" {
			bomb.BombAmount += Bomb
		}
	}
}

func explosion(system *SystemManagers, entity1 *Entity, timer *TimerComponent, damage *DamageComponent, position *PositionComponent, collision *CollisionComponent) {
	if !timer.Time.Before(time.Now()) {
		return
	}
	for entity2 := range system.DamageManager.damages {
		if entity1 != entity2 {
			fmt.Println("df")
		}
	}
}

func (entityManager *EntityManager) update(dt int, system *SystemManagers) {
	for _, entity := range entityManager.entities {
		position := entity.getPosition(system)
		motion := entity.getMotion(system)
		input := entity.getInput(system)
		collision := entity.getCollision(system)
		sprite := entity.getSprite(system)
		health := entity.getHealth(system)
		powerUp := entity.getPowerUp(system)
		timer := entity.getTimer(system)
		damage := entity.getDamage(system)
		bomb := entity.getBomb(system)

		if health != nil {
			powerUpSystem(position, powerUp, motion, health, bomb)
		}

		if timer != nil {
			explosion(system, entity, timer, damage, position, collision)
		}

		if health != nil {
			healthSystem(health, entity, system, sprite, input, position, collision)
		}

		if motion != nil && input != nil {
			inputSystem(input, motion)
		}

		if motion != nil && position != nil && collision.Enabled {
			movementSystem(position, motion)
		}

		if position != nil && collision != nil && health != nil {
			collisionSystem(entityManager, entity, position, collision, system, powerUp, health)
		}

		if position != nil && sprite != nil {
			renderSystem(input, sprite)
		}
		fmt.Printf("Entity %d - Position: (%.2f, %.2f) - Velocity: (%.2f, %.2f)\n", entity.Id, position.X, position.Y, motion.Velocity.X, motion.Velocity.Y)
	}
}
