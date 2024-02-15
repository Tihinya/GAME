package engine

import (
	"math/rand"
)

type GameState struct {
	Map      [][]*Tile `json:"map"`
	Players  []*Tile   `json:"players"`
	Powerups []*Tile   `json:"powerups"`
}

type Tile struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Size float64 `json:"size"`
}

func CreateGame(initialMap [][]int, players []int) {
	currentPlayer := 0
	maxPlayers := len(players)

	for y, rows := range initialMap {
		for x, blockType := range rows {
			xCoord := float64(x * 40)
			yCoord := float64(y * 40)
			switch blockType {
			case 1:
				if currentPlayer < maxPlayers {
					CreatePlayer(players[currentPlayer], xCoord, yCoord)
					currentPlayer++
				}
			case 2:
				if rand.Intn(100) < 30 {
					CreateBox(xCoord, yCoord)
				}
			case 4:
				CreateWall(xCoord, yCoord)
			}
		}
	}
}

func CreateMap() *GameState {
	gs := &GameState{
		Map:      make([][]*Tile, 13),
		Players:  make([]*Tile, 0),
		Powerups: make([]*Tile, 0),
	}

	for i := range gs.Map {
		gs.Map[i] = make([]*Tile, 31)
	}

	for entity, position := range positionManager.positions {
		if entity.Name != "player" && entity.Name != "powerup" {
			gs.Map[int(position.Y)/40][int(position.X)/40] = createTile(position.X, position.Y, entity.Name)
			continue
		}

		if entity.Name == "player" {
			gs.Players = append(gs.Players, createTile(position.X, position.Y, "", entity.Id))
		} else if entity.Name == "powerup" {
			entityPowerUP, exist := powerUpManager.powerUps[entity]
			if !exist {
				continue
			}
			gs.Powerups = append(gs.Powerups, createTile(position.X, position.Y, entityPowerUP.Name))
		}
	}

	return gs
}

func RemoveMap() {
	for _, e := range entityManager.entities {
		DeleteAllEntityComponents(e)
	}
	entityManager = NewEntityManager()

}

func createTile(x float64, y float64, name string, id ...int) *Tile {
	tile := &Tile{
		X:    x,
		Y:    y,
		Name: name,
	}

	if len(id) > 0 {
		tile.Id = id[0]
	}

	return tile
}
