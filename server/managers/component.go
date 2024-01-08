package managers

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
	Id   int
	Name string
}
type DamageComponent struct {
	DamageAmount int
	DamageType   string
}
type BombComponent struct {
	BombAmount int
}
type ExplosionComponent struct {
	Rang int
}
