package game

import (
	"time"

	"github.com/nsf/termbox-go"

	"github.com/NKTKLN/tetris/models"
)

// This function is responsible for controlling the falling figure with the keyboard
func (d *Data) ReadingKeyboard() {
	for {
		// Reading data from the keyboard
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			// Converting a pressed key to a command
			command, exists := models.ControlKeys[ev.Key]
			if !exists {
				command, exists = models.ControlCharacters[string(ev.Ch)]
				if !exists {	
					continue
				}
			}

			// Command execution
			switch command {
			case "quit":
				d.EndStats()
			default:
				if d.Move(command) {
					time.Sleep(100 * time.Microsecond)
					d.GameScreen()
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
