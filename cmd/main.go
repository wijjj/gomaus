package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	// Get screen dimensions
	screenWidth, screenHeight := robotgo.GetScreenSize()

	// Create a new local random generator
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("Press ESC to quit")

	// Start listening for keyboard events in a goroutine
	evtChan := hook.Start()
	defer hook.End()

	// Use a channel to listen for an exit signal
	exitChan := make(chan struct{})

	go func() {
		for ev := range evtChan {
			// @FIXME: key hook not working yet!
			if ev.Kind == hook.KeyHold && ev.Keychar == 0xff1b {
				fmt.Println("ESC pressed, exiting...")
				exitChan <- struct{}{}
				return
			}
		}
	}()

	for {
		select {
		case <-exitChan:
			return
		default:
			// Generate a random position within the screen bounds using the local generator
			x := randGen.Intn(screenWidth)
			y := randGen.Intn(screenHeight)

			// Smoothly move the mouse to the random position
			smoothMoveMouse(x, y, randGen)

			// Wait for a random interval between 1 to 6 seconds
			time.Sleep(time.Duration(randGen.Intn(5)+1) * time.Second)
		}
	}
}

func smoothMoveMouse(targetX, targetY int, randGen *rand.Rand) {
	currentX, currentY := robotgo.GetMousePos()
	steps := 100

	for i := 0; i < steps; i++ {
		t := float64(i) / float64(steps)
		t = t * t * (3 - 2*t) // Smoothstep easing function

		newX := int(float64(currentX) + t*float64(targetX-currentX))
		newY := int(float64(currentY) + t*float64(targetY-currentY))
		robotgo.Move(newX, newY)

		time.Sleep(time.Duration(randGen.Intn(10)+10) * time.Millisecond) // Small delay for smoother movement
	}
}
