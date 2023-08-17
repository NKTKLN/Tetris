package game

// Checks for filled lines, removes them if found and increases the score
func (d *Data) CheckFilledLines() {
lines:
	for yCord, line := range d.Field {
		for _, color := range line {
			if color == 0 {
				continue lines
			}
		}

		d.Field = append(d.Field[:yCord], d.Field[yCord+1:]...)
		d.Field = append([][]int{make([]int, 10)}, d.Field...)
		d.PlayerScore += 10
	}
}
