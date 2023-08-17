package game

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"

	"github.com/NKTKLN/tetris/models"
)

// Displaying the playing field
func (d *Data) GameScreen() {
	// Receiving terminal center
	terminalWidth, terminalHeight := termbox.Size()
	xStart := terminalWidth/2 - 9
	yStart := terminalHeight/2 - 7

	// Cleaning the terminal
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}

	// Displaying the game screen
	for yCord, line := range models.GameScreen {
		printText(xStart-1, yCord+yStart-1, line,
			termbox.ColorWhite)
	}

	// Displaying the playing field
	for yCord, line := range d.Field[4:] {
		for xCord, color := range line {
			if color != 0 {
				termbox.SetCell(xCord+xStart, yCord+yStart, []rune("#")[0],
					termbox.Attribute(color), termbox.ColorDefault)
			}
		}
	}

	// Displaying the next figure
	for yCord, line := range d.Figure.Type {
		if d.Figure.YCord+yCord >= 4 {
			for xCord, cell := range line {
				if cell == 1 {
					termbox.SetCell(d.Figure.XCord+xCord+xStart, d.Figure.YCord+yCord+yStart-4, []rune("#")[0],
						termbox.Attribute(d.Figure.Color), termbox.ColorDefault)
				}
			}
		}
		
	}

	// Displaying a falling figure
	for yCord, line := range models.FigureTypes[d.FigureType] {
		for xCord, cell := range line {
			if cell == 1 {
				termbox.SetCell(xCord+xStart+13, yCord+yStart, []rune("#")[0],
					termbox.Attribute(d.FigureType+2), termbox.ColorDefault)
			}
		}
	}

	// Displaying additional information
	printText(xStart+12, yStart+6, timeFormatting(d.StartTime),
		termbox.ColorCyan)

	printText(xStart+12, yStart+4, fmt.Sprintf("%05d", d.PlayerScore),
		termbox.ColorMagenta)

	termbox.Flush()
}

// Printing of the player's score along with the exit from the game
func (d *Data) EndStats() {
	termbox.Close()
	fmt.Printf(models.StatisticalMessage, d.PlayerScore, timeFormatting(d.StartTime))
	os.Exit(0)
}

// Adding text to the screen
func printText(xStart, yStart int, text string, color termbox.Attribute) {
	for xCord, char := range text {
		termbox.SetCell(xCord+xStart, yStart, char,
			color, termbox.ColorDefault)
	}
}

// Formatting the time
func timeFormatting(startTime time.Time) string {
	elapsedTime := time.Since(startTime)
	minutes := int(elapsedTime.Minutes())
	seconds := int(elapsedTime.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
