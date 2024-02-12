package gameloop

import (
	"fmt"
	"testing"
	"time"
)

func TestGameLoopAt60FPS(t *testing.T) {
	const fps = 60
	const expectedInterval = fps
	const tolerance = time.Millisecond * 2 // Adjust this based on acceptable tolerance (ms)

	var lastUpdate time.Time
	updateIntervals := make([]time.Duration, 0, fps)

	onUpdate := func(dt float64) {
		now := time.Now()
		if !lastUpdate.IsZero() {
			interval := now.Sub(lastUpdate)
			updateIntervals = append(updateIntervals, interval)
		}
		lastUpdate = now
	}

	loop := New(expectedInterval, onUpdate)
	go loop.Start()
	defer loop.Stop()

	// Let the game loop run for a short duration
	time.Sleep(time.Millisecond * 1000)

	// Analyze the intervals
	for _, interval := range updateIntervals {
		if interval < interval-expectedInterval-tolerance || interval > interval+expectedInterval+tolerance {
			t.Errorf("Interval %v is outside the expected range of %v ± %v", interval, expectedInterval, tolerance)
		}
	}

	fmt.Printf("Measured %d intervals, expected approximately %d intervals for 60fps\n", len(updateIntervals), fps)
}

func TestGameLoopOtherFunctionalities(t *testing.T) {
	var loopUpdated bool
	var fps = 30 // 30 fps aka 33.33ms
	loop := New(fps, func(dt float64) {
		loopUpdated = true
	})

	// Start function
	go loop.Start()
	fmt.Printf("Gameloop started at a tickrate of %v tps, testing if loop is updated\n", fps)
	// Wait for 100ms to check if loop updates it within its tickrate
	time.Sleep(time.Millisecond * 100)

	if !loopUpdated {
		t.Error("Gameloop did not update successfully, is it running?")
	}
	fmt.Println("Gameloop successfully updated")

	loopUpdated = false
	loop.Stop()
	fmt.Println("Gameloop stopped")

	// Stop function
	fmt.Println("Testing if loop is not updated while stopped")
	time.Sleep(time.Millisecond * 150)
	if loopUpdated {
		t.Error("Gameloop updated while it's supposed to be stopped")
	}
	fmt.Println("Gameloop successfully not updated while stopped")

	// Restart function
	fmt.Println("Testing restarting function:")
	go loop.Restart()
	fmt.Println("Gameloop restarted, checking if loop is updated")
	time.Sleep(time.Millisecond * 100)

	if !loopUpdated {
		t.Error("Gameloop did not update successfully, is it running?")
	}
	fmt.Println("Gameloop successfully updated")
}
