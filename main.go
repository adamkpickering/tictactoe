// The main loop.
package main

import (
	"fmt"
	"github.com/gdamore/tcell"
	"os"

	"github.com/adamkpickering/tictactoe/game"
	"github.com/adamkpickering/tictactoe/view"
)

func main() {

	// set up tcell screen
	s, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("main: %s", err.Error())
		os.Exit(1)
	}
	err = s.Init()
	if err != nil {
		fmt.Printf("main: %s", err.Error())
		os.Exit(1)
	}
	s.Clear()

	// create game and view object
	g := game.NewGame(nil)
	view := view.NewView(s)

	for {
		view.Draw(g)
		event := s.PollEvent()
		switch e := event.(type) {
		case *tcell.EventKey:
			key := e.Key()
			switch key {
			case tcell.KeyRune:
				if e.Rune() == 'q' {
					s.Fini()
					return
				}
			case tcell.KeyEnter:
				_ = g.PlayTurn(view.SelectedX, view.SelectedY)
				// need to handle error - goes to log
			default:
				view.MoveSelect(key)
			}
		}
	}

	s.Fini()
}
