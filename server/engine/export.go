package engine

import "math/rand"

type GameState struct {
	Map      [][]*Tile `json:"map"`
	Players  []*Tile   `json:"players"`
	Powerups []*Tile   `json:"powerups"`
}

type Tile struct {
	Name string  `json:"name"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Size float64 `json:"size"`
}

// 4. generate initial game
// 5. create map with static objects y13*x31
// 6. Update map every 1/60 of a second

// takes initial map from config. 1 - spawn point, 2 - floor, 3 - spawn point protection, 4 - walls
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
				if rand.Intn(100) < 20 {
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
			gs.Players = append(gs.Players, createTile(position.X, position.Y, ""))
		} else if entity.Name == "powerup" {
			gs.Powerups = append(gs.Powerups, createTile(position.X, position.Y, ""))
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

func createTile(x float64, y float64, name string, size ...float64) *Tile {
	tile := &Tile{
		X:    x,
		Y:    y,
		Name: name,
	}

	if len(size) > 0 {
		tile.Size = size[0]
	}

	return tile
}
