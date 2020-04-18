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


type View struct {
	SelectedSquare int
	screen tcell.Screen
}

func NewView(s tcell.Screen) *View {
	return &View{
		screen: s,
		SelectedSquare: 0,
	}
}

func (v *View) MoveSelect(key tcell.Key) error {
	switch v.SelectedSquare {
	case 0:
		switch key {
		case tcell.KeyUp:
			v.SelectedSquare = 6
		case tcell.KeyRight:
			v.SelectedSquare = 1
		case tcell.KeyDown:
			v.SelectedSquare = 3
		case tcell.KeyLeft:
			v.SelectedSquare = 2
		default:
			return fmt.Errorf("%s is not a valid key")
		}
	case 1:
		switch key {
		case tcell.KeyUp:
			v.SelectedSquare = 7
		case tcell.KeyRight:
			v.SelectedSquare = 2
		case tcell.KeyDown:
			v.SelectedSquare = 4
		case tcell.KeyLeft:
			v.SelectedSquare = 0
		default:
			return fmt.Errorf("%s is not a valid key")
		}
	case 2:
		switch key {
		case tcell.KeyUp:
			v.SelectedSquare = 8
		case tcell.KeyRight:
			v.SelectedSquare = 0
		case tcell.KeyDown:
			v.SelectedSquare = 5
		case tcell.KeyLeft:
			v.SelectedSquare = 1
		default:
			return fmt.Errorf("%s is not a valid key")
		}
	case 3:
		switch key {
		case tcell.KeyUp:
			v.SelectedSquare = 0
		case tcell.KeyRight:
			v.SelectedSquare = 4
		case tcell.KeyDown:
			v.SelectedSquare = 6
		case tcell.KeyLeft:
			v.SelectedSquare = 5
		default:
			return fmt.Errorf("%s is not a valid key")
		}
	case 4:
		switch key {
		case tcell.KeyUp:
			v.SelectedSquare = 1
		case tcell.KeyRight:
			v.SelectedSquare = 5
		case tcell.KeyDown:
			v.SelectedSquare = 7
		case tcell.KeyLeft:
			v.SelectedSquare = 3
		default:
			return fmt.Errorf("%s is not a valid key")
		}
	case 5:
		switch key {
		case tcell.KeyUp:
			v.SelectedSquare = 2
		case tcell.KeyRight:
			v.SelectedSquare = 3
		case tcell.KeyDown:
			v.SelectedSquare = 8
		case tcell.KeyLeft:
			v.SelectedSquare = 4
		default:
			return fmt.Errorf("%s is not a valid key")
		}
	case 6:
		switch key {
		case tcell.KeyUp:
			v.SelectedSquare = 3
		case tcell.KeyRight:
			v.SelectedSquare = 7
		case tcell.KeyDown:
			v.SelectedSquare = 0
		case tcell.KeyLeft:
			v.SelectedSquare = 8
		default:
			return fmt.Errorf("%s is not a valid key")
		}
	case 7:
		switch key {
		case tcell.KeyUp:
			v.SelectedSquare = 4
		case tcell.KeyRight:
			v.SelectedSquare = 8
		case tcell.KeyDown:
			v.SelectedSquare = 1
		case tcell.KeyLeft:
			v.SelectedSquare = 6
		default:
			return fmt.Errorf("%s is not a valid key")
		}
	case 8:
		switch key {
		case tcell.KeyUp:
			v.SelectedSquare = 5
		case tcell.KeyRight:
			v.SelectedSquare = 6
		case tcell.KeyDown:
			v.SelectedSquare = 2
		case tcell.KeyLeft:
			v.SelectedSquare = 7
		default:
			return fmt.Errorf("%s is not a valid key")
		}
	default:
		return errors.New("SelectedSquare out of range")
	}
	return nil
}


func (v *View) Draw() {

	board := v.getFreshBoard()

	select_style := tcell.StyleDefault.Background(tcell.ColorWhite)
	switch v.SelectedSquare {
	case 0:
		board[1][1].Style = select_style
		board[1][2].Style = select_style
		board[1][3].Style = select_style
	case 1:
		board[1][5].Style = select_style
		board[1][6].Style = select_style
		board[1][7].Style = select_style
	case 2:
		board[1][9].Style = select_style
		board[1][10].Style = select_style
		board[1][11].Style = select_style
	case 3:
		board[3][1].Style = select_style
		board[3][2].Style = select_style
		board[3][3].Style = select_style
	case 4:
		board[3][5].Style = select_style
		board[3][6].Style = select_style
		board[3][7].Style = select_style
	case 5:
		board[3][9].Style = select_style
		board[3][10].Style = select_style
		board[3][11].Style = select_style
	case 6:
		board[5][1].Style = select_style
		board[5][2].Style = select_style
		board[5][3].Style = select_style
	case 7:
		board[5][5].Style = select_style
		board[5][6].Style = select_style
		board[5][7].Style = select_style
	case 8:
		board[5][9].Style = select_style
		board[5][10].Style = select_style
		board[5][11].Style = select_style
	}

	for j, row := range board {
		for i, cell := range row {
			v.screen.SetContent(i, j, cell.Rune, nil, cell.Style)
		}
	}

	v.screen.Show()

}


func (v *View) getFreshBoard() [7][13]struct{Rune rune; Style tcell.Style} {

	var board [7][13]struct{
		Rune rune
		Style tcell.Style
	}

	for j, _ := range board {
		for i, _ := range board[j] {
			board[j][i].Style = tcell.StyleDefault
		}
	}

	board[0][0].Rune = tcell.RuneULCorner
	board[0][1].Rune = tcell.RuneHLine
	board[0][2].Rune = tcell.RuneHLine
	board[0][3].Rune = tcell.RuneHLine
	board[0][4].Rune = tcell.RuneTTee
	board[0][5].Rune = tcell.RuneHLine
	board[0][6].Rune = tcell.RuneHLine
	board[0][7].Rune = tcell.RuneHLine
	board[0][8].Rune = tcell.RuneTTee
	board[0][9].Rune = tcell.RuneHLine
	board[0][10].Rune = tcell.RuneHLine
	board[0][11].Rune = tcell.RuneHLine
	board[0][12].Rune = tcell.RuneURCorner

	board[1][0].Rune = tcell.RuneVLine
	board[1][1].Rune = ' '
	board[1][2].Rune = ' '
	board[1][3].Rune = ' '
	board[1][4].Rune = tcell.RuneVLine
	board[1][5].Rune = ' '
	board[1][6].Rune = ' '
	board[1][7].Rune = ' '
	board[1][8].Rune = tcell.RuneVLine
	board[1][9].Rune = ' '
	board[1][10].Rune = ' '
	board[1][11].Rune = ' '
	board[1][12].Rune = tcell.RuneVLine

	board[2][0].Rune = tcell.RuneLTee
	board[2][1].Rune = tcell.RuneHLine
	board[2][2].Rune = tcell.RuneHLine
	board[2][3].Rune = tcell.RuneHLine
	board[2][4].Rune = tcell.RunePlus
	board[2][5].Rune = tcell.RuneHLine
	board[2][6].Rune = tcell.RuneHLine
	board[2][7].Rune = tcell.RuneHLine
	board[2][8].Rune = tcell.RunePlus
	board[2][9].Rune = tcell.RuneHLine
	board[2][10].Rune = tcell.RuneHLine
	board[2][11].Rune = tcell.RuneHLine
	board[2][12].Rune = tcell.RuneRTee

	board[3][0].Rune = tcell.RuneVLine
	board[3][1].Rune = ' '
	board[3][2].Rune = ' '
	board[3][3].Rune = ' '
	board[3][4].Rune = tcell.RuneVLine
	board[3][5].Rune = ' '
	board[3][6].Rune = ' '
	board[3][7].Rune = ' '
	board[3][8].Rune = tcell.RuneVLine
	board[3][9].Rune = ' '
	board[3][10].Rune = ' '
	board[3][11].Rune = ' '
	board[3][12].Rune = tcell.RuneVLine

	board[4][0].Rune = tcell.RuneLTee
	board[4][1].Rune = tcell.RuneHLine
	board[4][2].Rune = tcell.RuneHLine
	board[4][3].Rune = tcell.RuneHLine
	board[4][4].Rune = tcell.RunePlus
	board[4][5].Rune = tcell.RuneHLine
	board[4][6].Rune = tcell.RuneHLine
	board[4][7].Rune = tcell.RuneHLine
	board[4][8].Rune = tcell.RunePlus
	board[4][9].Rune = tcell.RuneHLine
	board[4][10].Rune = tcell.RuneHLine
	board[4][11].Rune = tcell.RuneHLine
	board[4][12].Rune = tcell.RuneRTee

	board[5][0].Rune = tcell.RuneVLine
	board[5][1].Rune = ' '
	board[5][2].Rune = ' '
	board[5][3].Rune = ' '
	board[5][4].Rune = tcell.RuneVLine
	board[5][5].Rune = ' '
	board[5][6].Rune = ' '
	board[5][7].Rune = ' '
	board[5][8].Rune = tcell.RuneVLine
	board[5][9].Rune = ' '
	board[5][10].Rune = ' '
	board[5][11].Rune = ' '
	board[5][12].Rune = tcell.RuneVLine

	board[6][0].Rune = tcell.RuneLLCorner
	board[6][1].Rune = tcell.RuneHLine
	board[6][2].Rune = tcell.RuneHLine
	board[6][3].Rune = tcell.RuneHLine
	board[6][4].Rune = tcell.RuneBTee
	board[6][5].Rune = tcell.RuneHLine
	board[6][6].Rune = tcell.RuneHLine
	board[6][7].Rune = tcell.RuneHLine
	board[6][8].Rune = tcell.RuneBTee
	board[6][9].Rune = tcell.RuneHLine
	board[6][10].Rune = tcell.RuneHLine
	board[6][11].Rune = tcell.RuneHLine
	board[6][12].Rune = tcell.RuneLRCorner

	return board
}




//	g := game.NewGame(nil)
//
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

//func printGame(b [9]byte) {
//	pb := [9]byte{}
//	for i, element := range b {
//		if element == 0 {
//			pb[i] = ' '
//		} else {
//			pb[i] = element
//		}
//	}
//	fmt.Printf("\tBoard:\t\t\tLegend:\n")
//	fmt.Printf("\t-------------------\t-------------------\n")
//	fmt.Printf("\t|  %c  |  %c  |  %c  |\t|  0  |  1  |  2  |\n", pb[0], pb[1], pb[2])
//	fmt.Printf("\t-------------------\t-------------------\n")
//	fmt.Printf("\t|  %c  |  %c  |  %c  |\t|  3  |  4  |  5  |\n", pb[3], pb[4], pb[5])
//	fmt.Printf("\t-------------------\t-------------------\n")
//	fmt.Printf("\t|  %c  |  %c  |  %c  |\t|  6  |  7  |  8  |\n", pb[6], pb[7], pb[8])
//	fmt.Printf("\t-------------------\t-------------------\n")
//}
