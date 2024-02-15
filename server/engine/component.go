package engine

import (
	"time"
)

type PositionComponent struct {
	X, Y  float64
	Size  float64
	Angle float64
}

type Vec2 struct {
	X, Y float64
}

type MotionComponent struct {
	SpeedMultiplier float64
	Velocity        Vec2
	Acceleration    Vec2
}

type CollisionComponent struct {
	Enabled bool
}

type HealthComponent struct {
	lastTimeDamage time.Time
	CurrentHealth  int
	MaxHealth      int
	OnDestroy      func()
}

type InputComponent struct {
	Input map[string]bool
}

type TimerComponent struct {
	Time time.Time
}

type PowerUpComponent struct {
	Name string
}

type DamageComponent struct {
	DamageAmount int
}

// type BM struct {
// 	BlastRadius int
// 	BombAmount  int
// 	PlacedBombs int
// }

type BombComponent struct {
	BlastRadius int
	BombAmount  int
	PlacedBombs int
	Owner       *Entity
}
type UserEntityComponent struct {
	entity *Entity
}

type ExplosionStopperComponent struct {
	passable bool
}

// doesn't need any attributes,
// just for identifying explosions/boxes/walls
type (
	ExplosionComponent struct{}
)
