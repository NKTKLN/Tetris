package game

import (
	"math/rand"

	"github.com/NKTKLN/tetris/models"
)

// This function is responsible for creating a new shape with random coordinates on the x-axis and random rotation
func (d *Data) Create() {
	d.Figure = models.Figure{
		YCord: 0,
		XCord: 0,
		Color: d.FigureType + 2,
		Type:  models.FigureTypes[d.FigureType],
	}
	
	// Randomization of x coordinate and rotation position
	orientation := rand.Intn(4)
	for ori := 0; ori <= orientation; ori++ {
		d.Move("rotate")
	}

	d.Figure.XCord = rand.Intn(10 - len(d.Figure.Type[0]))
	d.Figure.YCord = 4 - len(d.Figure.Type)

	// Lowering the figure to the edge of the glass
lines:
	for index := len(d.Figure.Type) - 1; index >= 0; index-- {
		for _, block := range d.Figure.Type[index] {
			if block == 1 {
				break lines
			}
		}

		d.Figure.YCord++
	}
}

// This function is responsible for correct movement of the figure on the field
func (d *Data) Move(command string) bool {
	// Creating a copy of a falling figure with a new move
	newFigure := d.Figure
	switch command {
	case "left":
		newFigure.XCord--
	case "right":
		newFigure.XCord++
	case "rotate":
		rows := len(d.Figure.Type)
		cols := len(d.Figure.Type[0])
		newFigure.Type = make([][]int, cols)

		for index := range newFigure.Type {
			newFigure.Type[index] = make([]int, rows)
		}

		for row, line := range d.Figure.Type {
			for col, block := range line {
				newFigure.Type[col][rows-row-1] = block
			}
		}
	case "down":
		newFigure.YCord++
	case "bottom":
		for d.Move("down") {
		}
		return true
	default:
		return false
	}

	// Checking for correctness of a new move
	for row, line := range newFigure.Type {
		for col, block := range line {
			if block == 1 {
				newXCord := newFigure.XCord + col
				newYCord := newFigure.YCord + row

				if (newXCord < 0 || newXCord >= 10 || newYCord < 0 || newYCord >= 19) ||
					d.Field[newYCord][newXCord] != 0 {
					return false
				}
			}
		}
	}

	d.Figure = newFigure
	return true
}

// Adding a fallen figure to the field
func (d *Data) Save() {
	for row, line := range d.Figure.Type {
		for col, block := range line {
			if block == 1 {
				xCord := d.Figure.XCord + col
				yCord := d.Figure.YCord + row

				d.Field[yCord][xCord] = d.Figure.Color
			}
		}
	}
}