package engine

import (
	"bomberman-dom/helpers"
	"bomberman-dom/models"
)

func broadcastBomb(X, Y float64, action string) {
	event := helpers.SerializeData("game_bomb", models.GameBomb{
		Action: action,
		X:      X,
		Y:      Y,
	})
	broadcaster.BroadcastAllClients(event)
}

func broadcastExplosion(X, Y float64, action string) {
	explosionEvent := helpers.SerializeData("game_explosion", models.GameExplosion{
		Action: action,
		X:      X,
		Y:      Y,
	})
	broadcaster.BroadcastAllClients(explosionEvent)
}

func broadcastObstacle(X, Y float64, obstacleType, action string) {
	obstacleEvent := helpers.SerializeData("game_obstacle", models.GameObstacle{
		Type:   obstacleType,
		Action: action,
		X:      X,
		Y:      Y,
	})
	broadcaster.BroadcastAllClients(obstacleEvent)
}

func broadcastPowerup(X, Y float64, powerupType int, action string) {
	powerupEvent := helpers.SerializeData("game_powerup", models.GamePowerup{
		Type:   powerupType,
		Action: action,
		X:      X,
		Y:      Y,
	})
	broadcaster.BroadcastAllClients(powerupEvent)
}

func broadcastMotion(X, Y float64, entity *Entity) {
	socketClientId := userEntityManager.GetUserIdByEntity(entity)
	motionEvent := helpers.SerializeData("game_motion", models.GamePlayer{
		ClientId: socketClientId,
		X:        X,
		Y:        Y,
	})
	broadcaster.BroadcastAllClients(motionEvent)
}

func broadcastDeleteExplosions(e *Entity) {
	pos := positionManager.GetPosition(e)
	broadcastExplosion(pos.X, pos.Y, "delete")
}
