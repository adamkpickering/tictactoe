package main

import (
	"fmt"
	"os"
	"github.com/gdamore/tcell"

	"github.com/adamkpickering/tictactoe/pkg/game"
	"github.com/adamkpickering/tictactoe/pkg/view"
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
		// check for win conditions
		result := g.CheckWin()
		switch {
		case result == 'X' || result == 'O':
			s.Fini()
			fmt.Printf("Player %c won.\n", result)
			return
		case result == 1:
			s.Fini()
			fmt.Printf("It was a draw.\n")
			return
		}

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
				_ = g.PlayTurn(view.SelectedRow, view.SelectedCol)
				// need to handle error - goes to log
			default:
				view.MoveSelect(key)
			}
		}
	}

	s.Fini()
}
