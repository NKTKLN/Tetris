package game

import (
	"time"
	"math/rand"
)

func (d *Data) Game() {
	go d.ReadingKeyboard()
	
	// Main game loop
	for {
		d.Create()

		// Randomization of the following figure
		d.FigureType = rand.Intn(7)

		for {
			d.GameScreen()

			// Calculation of the drop delay relative to the player's score
			fallDelay := 20 * 24
			if d.PlayerScore/50 < 25 {
				fallDelay = 20 * (d.PlayerScore / 50)
			}
			time.Sleep(time.Duration(500-fallDelay) * time.Millisecond)

			if !d.Move("down") {
				break
			}
		}

		d.PlayerScore++

		d.Save()
		d.CheckFilledLines()

		// Checking for overrun of the glass
		for _, line := range d.Field[:4] {
			for _, color := range line {
				if color != 0 {
					d.EndStats()
				}
			}
		}
	}
}
