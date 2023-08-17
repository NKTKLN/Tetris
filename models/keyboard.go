package models

import "github.com/nsf/termbox-go"

var ControlKeys = map[termbox.Key]string{
	termbox.KeyArrowUp:    "rotate",
	termbox.KeyArrowDown:  "down",
	termbox.KeyArrowLeft:  "left",
	termbox.KeyArrowRight: "right",
	termbox.KeyEsc:        "quit",
	termbox.KeyEnter:      "bottom",
}

var ControlCharacters = map[string]string{
	"r": "rotate",
	"s": "down",
	"a": "left",
	"d": "right",
	"q": "quit",
	"f": "bottom",
}
