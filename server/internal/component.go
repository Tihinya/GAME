package internal

import (
	"time"
)

type PositionComponent struct {
	X, Y float64
	Size float64
}
type Vec2 struct {
	X, Y float64
}
type MotionComponent struct {
	Velocity     Vec2
	Acceleration Vec2
}
type SpriteComponent struct {
	Texture string
}
type CollisionComponent struct {
	Enabled bool
}
type HealthComponent struct {
	CurrentHealth int
	MaxHealth     int
}
type InputComponent struct {
	Input map[string]bool
}
type TimerComponent struct {
	Time time.Time
}
type PowerUpComponent struct {
	ExtraBombs          int
	ExtraExplosionRange int
	ExtraSpeed          float64
}
type DamageComponent struct {
	DamageAmount int
}
type BombComponent struct {
	BlastRadius int
	IsActive    bool
	Owner       *Entity
}

// doesn't need any attributes,
// just for identifying explosions
type ExplosionComponent struct{}
type BoxComponent struct{}
