package managers

import (
	"fmt"
	"time"
)

func (entityManager *EntityManager) update(dt int, system *SystemManagers) {
	for _, entity := range entityManager.entities {
		inputSystem(entity, system)
		movementSystem(entity, system, entityManager)
		// powerUpSystem(entity, system)
		// explosionSystem(entity, system)
		// healthSystem(entity, system)
		// renderSystem(entity, system)
		// collisionSystem(entity, system)

	}
}

func inputSystem(entity *Entity, system *SystemManagers) {
	motion := entity.getMotion(system)
	input := entity.getInput(system)
	if motion == nil && input == nil {
		return
	}

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

func collisionSystem(entity *Entity, system *SystemManagers, entityManager *EntityManager) bool {
	entity1Position := entity.getPosition(system)
	entity1Collision := entity.getCollision(system)
	health := entity.getHealth(system)
	powerUps := entity.getPowerUp(system)

	if entity1Position == nil && entity1Collision == nil && health == nil {
		return false
	}

	for _, entity2 := range entityManager.entities {
		if entity != entity2 {
			entity2Position := entity2.getPosition(system)
			entity2Collision := entity2.getCollision(system)
			entity2PowerUp := entity2.getPowerUp(system)
			entity2Damage := entity2.getDamage(system)

			if entity2Position == nil {
				return false
			}
			collision := checkCollision(entity1Position.X, entity1Position.Y, entity2Position.X, entity2Position.Y, entity1Position.Size, entity2Position.Size)

			if entity2PowerUp != nil && collision {
				powerUps = append(powerUps, entity2PowerUp...)
			}
			if entity2Damage != nil && collision {
				health.CurrentHealth -= entity2Damage.DamageAmount
			}
			if entity2Collision != nil {
				if !entity2Collision.Enabled {
					return true
				}
			}
			return collision
		}
	}
	return false
}

func renderSystem(entity *Entity, system *SystemManagers) {
	input := entity.getInput(system)
	sprite := entity.getSprite(system)

	if input == nil && sprite == nil {
		return
	}
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

func healthSystem(entity *Entity, system *SystemManagers) {
	position := entity.getPosition(system)
	input := entity.getInput(system)
	collision := entity.getCollision(system)
	sprite := entity.getSprite(system)
	health := entity.getHealth(system)

	if health == nil {
		return
	}

	if health.CurrentHealth <= 0 {
		system.HealthManager.DeleteComponet(entity, health)
		system.SpriteManager.DeleteComponet(entity, sprite)
		system.InputManager.DeleteComponet(entity, input)
		system.PositionManager.DeleteComponet(entity, position)
		system.CollisionManager.DeleteComponet(entity, collision)
	}
}

func explosionSystem(entity *Entity, system *SystemManagers) {
	timer := entity.getTimer(system)
	if timer == nil {
		return
	}
	if !timer.Time.Before(time.Now()) {
		return
	}

	for entity2 := range system.DamageManager.damages {
		if entity != entity2 {
			fmt.Println("df")
		}
	}
}

func powerUpSystem(entity *Entity, system *SystemManagers) {
	health := entity.getHealth(system)
	motion := entity.getMotion(system)
	powerUps := entity.getPowerUp(system)
	bomb := entity.getBomb(system)

	if health == nil {
		return
	}

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

func movementSystem(entity *Entity, system *SystemManagers, entityManager *EntityManager) {
	position := entity.getPosition(system)
	motion := entity.getMotion(system)

	if motion == nil {
		return
	}

	position.X += motion.Velocity.X
	position.Y += motion.Velocity.Y

	motion.Velocity.X += motion.Acceleration.X
	motion.Velocity.Y += motion.Acceleration.Y

	if collisionSystem(entity, system, entityManager) {
		position.X -= motion.Velocity.X
		position.Y -= motion.Velocity.Y

		motion.Velocity.X -= motion.Acceleration.X
		motion.Velocity.Y -= motion.Acceleration.Y
	}

}
