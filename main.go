package main

import (
	"fmt"
	"os"
	"github.com/gdamore/tcell"

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
		// check for win conditions
		result, win_info := g.CheckWin()
		switch result {
		case game.Win:
			s.Fini()
			fmt.Printf("Player %c won.\n", win_info.Letter())
			return
		case game.Draw:
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
				_ = g.PlayTurn(view.SelectedX, view.SelectedY)
				// need to handle error - goes to log
			default:
				view.MoveSelect(key)
			}
		}
	}

	s.Fini()
}
