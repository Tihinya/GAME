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
	explodeBoxEvent := helpers.SerializeData("game_obstacle", models.GameObstacle{
		Type:   obstacleType,
		Action: action,
		X:      X,
		Y:      Y,
	})
	broadcaster.BroadcastAllClients(explodeBoxEvent)
}
