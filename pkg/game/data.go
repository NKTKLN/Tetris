package game

import (
	"math/rand"
	"time"

	"github.com/NKTKLN/tetris/models"
)

type Data struct {
	Figure     models.Figure
	Field      [][]int
	FigureType int

	StartTime   time.Time
	PlayerScore int
}

// Creates a variable with default game data
func DefaultData() (data Data) {
	// Saving start time
	data.StartTime = time.Now()

	// Creating an empty field
	data.Field = make([][]int, 19)

	for yCord := range data.Field {
		data.Field[yCord] = make([]int, 10)
	}

	// Randomization of the first figure
	data.FigureType = rand.Intn(7)

	return
}
