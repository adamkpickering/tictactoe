package main

import (
	"fmt"
	"os"
	"time"
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

	draw_screen(s)

	s.Show()

	time.Sleep(time.Second * 2)

	s.Fini()
}


func draw_screen(screen tcell.Screen) {

	screen.SetContent(0, 0, tcell.RuneULCorner, nil, tcell.StyleDefault)
	screen.SetContent(1, 0, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(2, 0, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(3, 0, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(4, 0, tcell.RuneTTee, nil, tcell.StyleDefault)
	screen.SetContent(5, 0, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(6, 0, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(7, 0, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(8, 0, tcell.RuneTTee, nil, tcell.StyleDefault)
	screen.SetContent(9, 0, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(10, 0, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(11, 0, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(12, 0, tcell.RuneURCorner, nil, tcell.StyleDefault)

	screen.SetContent(0, 1, tcell.RuneVLine, nil, tcell.StyleDefault)
	screen.SetContent(1, 1, ' ', nil, tcell.StyleDefault)
	screen.SetContent(2, 1, ' ', nil, tcell.StyleDefault)
	screen.SetContent(3, 1, ' ', nil, tcell.StyleDefault)
	screen.SetContent(4, 1, tcell.RuneVLine, nil, tcell.StyleDefault)
	screen.SetContent(5, 1, ' ', nil, tcell.StyleDefault)
	screen.SetContent(6, 1, ' ', nil, tcell.StyleDefault)
	screen.SetContent(7, 1, ' ', nil, tcell.StyleDefault)
	screen.SetContent(8, 1, tcell.RuneVLine, nil, tcell.StyleDefault)
	screen.SetContent(9, 1, ' ', nil, tcell.StyleDefault)
	screen.SetContent(10, 1, ' ', nil, tcell.StyleDefault)
	screen.SetContent(11, 1, ' ', nil, tcell.StyleDefault)
	screen.SetContent(12, 1, tcell.RuneVLine, nil, tcell.StyleDefault)

	screen.SetContent(0, 2, tcell.RuneLTee, nil, tcell.StyleDefault)
	screen.SetContent(1, 2, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(2, 2, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(3, 2, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(4, 2, tcell.RunePlus, nil, tcell.StyleDefault)
	screen.SetContent(5, 2, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(6, 2, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(7, 2, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(8, 2, tcell.RunePlus, nil, tcell.StyleDefault)
	screen.SetContent(9, 2, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(10, 2, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(11, 2, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(12, 2, tcell.RuneRTee, nil, tcell.StyleDefault)

	screen.SetContent(0, 3, tcell.RuneVLine, nil, tcell.StyleDefault)
	screen.SetContent(1, 3, ' ', nil, tcell.StyleDefault)
	screen.SetContent(2, 3, ' ', nil, tcell.StyleDefault)
	screen.SetContent(3, 3, ' ', nil, tcell.StyleDefault)
	screen.SetContent(4, 3, tcell.RuneVLine, nil, tcell.StyleDefault)
	screen.SetContent(5, 3, ' ', nil, tcell.StyleDefault)
	screen.SetContent(6, 3, ' ', nil, tcell.StyleDefault)
	screen.SetContent(7, 3, ' ', nil, tcell.StyleDefault)
	screen.SetContent(8, 3, tcell.RuneVLine, nil, tcell.StyleDefault)
	screen.SetContent(9, 3, ' ', nil, tcell.StyleDefault)
	screen.SetContent(10, 3, ' ', nil, tcell.StyleDefault)
	screen.SetContent(11, 3, ' ', nil, tcell.StyleDefault)
	screen.SetContent(12, 3, tcell.RuneVLine, nil, tcell.StyleDefault)

	screen.SetContent(0, 4, tcell.RuneLTee, nil, tcell.StyleDefault)
	screen.SetContent(1, 4, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(2, 4, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(3, 4, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(4, 4, tcell.RunePlus, nil, tcell.StyleDefault)
	screen.SetContent(5, 4, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(6, 4, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(7, 4, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(8, 4, tcell.RunePlus, nil, tcell.StyleDefault)
	screen.SetContent(9, 4, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(10, 4, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(11, 4, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(12, 4, tcell.RuneRTee, nil, tcell.StyleDefault)

	screen.SetContent(0, 5, tcell.RuneVLine, nil, tcell.StyleDefault)
	screen.SetContent(1, 5, ' ', nil, tcell.StyleDefault)
	screen.SetContent(2, 5, ' ', nil, tcell.StyleDefault)
	screen.SetContent(3, 5, ' ', nil, tcell.StyleDefault)
	screen.SetContent(4, 5, tcell.RuneVLine, nil, tcell.StyleDefault)
	screen.SetContent(5, 5, ' ', nil, tcell.StyleDefault)
	screen.SetContent(6, 5, ' ', nil, tcell.StyleDefault)
	screen.SetContent(7, 5, ' ', nil, tcell.StyleDefault)
	screen.SetContent(8, 5, tcell.RuneVLine, nil, tcell.StyleDefault)
	screen.SetContent(9, 5, ' ', nil, tcell.StyleDefault)
	screen.SetContent(10, 5, ' ', nil, tcell.StyleDefault)
	screen.SetContent(11, 5, ' ', nil, tcell.StyleDefault)
	screen.SetContent(12, 5, tcell.RuneVLine, nil, tcell.StyleDefault)

	screen.SetContent(0, 6, tcell.RuneLLCorner, nil, tcell.StyleDefault)
	screen.SetContent(1, 6, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(2, 6, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(3, 6, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(4, 6, tcell.RuneBTee, nil, tcell.StyleDefault)
	screen.SetContent(5, 6, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(6, 6, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(7, 6, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(8, 6, tcell.RuneBTee, nil, tcell.StyleDefault)
	screen.SetContent(9, 6, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(10, 6, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(11, 6, tcell.RuneHLine, nil, tcell.StyleDefault)
	screen.SetContent(12, 6, tcell.RuneLRCorner, nil, tcell.StyleDefault)
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
