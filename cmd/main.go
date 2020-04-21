package main

import (
	"fmt"
	"os"
	"errors"
	"github.com/gdamore/tcell"
)

func main() {

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

	// create game object
	g := game.NewGame(nil)

	view := NewView(s)
	for {
		view.Draw()
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
				s.Fini()
				return
			default:
				view.MoveSelect(key)
			}
		}
	}

	s.Fini()
}


//	for {
//		printGame(g.Board)
//		fmt.Printf("It is player %c's turn.\n", g.Turn)
//		fmt.Printf("Pick a square: ")
//
//		input := make([]byte, 10)
//		if _, err := os.Stdin.Read(input); err != nil {
//			fmt.Fprintf(os.Stderr, "There was an error reading input: %s\n", err.Error())
//			continue
//		}
//
//		square := int(input[0]) - 48
//		if err := g.PlayTurn(square); err != nil {
//			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
//			continue
//		}
//
//		result := g.CheckWin()
//		switch result {
//		case 'X':
//			fmt.Printf("Player %c won. Good Game!\n", result)
//			printGame(g.Board)
//			os.Exit(0)
//		case 'O':
//			fmt.Printf("Player %c won. Good Game!\n", result)
//			printGame(g.Board)
//			os.Exit(0)
//		case 1:
//			fmt.Print("It is a draw. Good game!\n")
//			printGame(g.Board)
//			os.Exit(0)
//		}
//	}
