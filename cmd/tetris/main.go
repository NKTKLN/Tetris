package main

import (
	"flag"
	"fmt"

	"github.com/nsf/termbox-go"

	"github.com/NKTKLN/tetris/models"
	"github.com/NKTKLN/tetris/pkg/game"
)

func main() {
	// Announcement of game flags
	isKeysFlag := flag.Bool("keys", false, "Control information.")
	isInfoFlag := flag.Bool("info", false, "Application information.")
	flag.Parse()

	switch {
	case *isInfoFlag:
		fmt.Printf(models.InfoMessage, models.Version)
		return
	case *isKeysFlag:
		fmt.Print(models.KeysMessage)
		return
	}

	// Setting up the termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetInputMode(termbox.InputEsc)

	// Checking the terminal window size
	terminalWidth, terminalHeight := termbox.Size()
	if terminalWidth < 20 || terminalHeight < 17 {
		termbox.Close()
		fmt.Println(models.SizeErrorMessage)
		return
	}

	// Launching the main game function
	data := game.DefaultData()
	data.Game()
}
