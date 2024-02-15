package gameloop

import (
	"time"
)

type GameLoop struct {
	onUpdate func(float64)
	tickRate int
	quit     chan bool
}

// Create new game loop
func New(tickRate int, onUpdate func(float64)) *GameLoop {
	return &GameLoop{
		onUpdate: onUpdate,
		tickRate: tickRate,
		quit:     make(chan bool),
	}
}

func (gl *GameLoop) startLoop() {

	tickInterval := time.Second / time.Duration(gl.tickRate)
	timeStart := time.Now()

	ticker := time.NewTicker(tickInterval)

	for {
		select {
		case t := <-ticker.C:
			gl.onUpdate(time.Since(timeStart).Seconds())
			timeStart = t

		case <-gl.quit:
			ticker.Stop()
			return
		}
	}
}

func (gl *GameLoop) GetTickRate() int {
	return gl.tickRate
}

// Set tickRate and restart game loop
func (gl *GameLoop) SetTickRate(tickRate int) {
	gl.tickRate = tickRate
	gl.Restart()
}

// Set onUpdate func
func (gl *GameLoop) SetOnUpdate(onUpdate func(float64)) {
	gl.onUpdate = onUpdate
}

// Start game loop
func (gl *GameLoop) Start() {
	gl.startLoop()
}

// Stop game loop
func (gl *GameLoop) Stop() {
	close(gl.quit)
}

// Restart game loop
func (gl *GameLoop) Restart() {
	gl.Stop()
	gl.Start()
}
