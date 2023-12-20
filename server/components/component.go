package components

import "time"

type PositionComponent struct {
	X, Y float64
}
type Vec2 struct {
	X, Y int
}
type MotionComponent struct {
	Velocity     Vec2
	Acceleration Vec2
}
type SpriteComponent struct {
	Texture string
}
type CollisionComponent struct {
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
}
type DamageComponent struct {
	DamageAmount int
	DamageType   string
}
