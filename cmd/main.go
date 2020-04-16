package main

import (
	"fmt"
	//"os"
	"github.com/awesome-gocui/gocui"
	//"adampickering.ca/tictactoe/pkg/game"
)

func quit(_ *gocui.Gui, _ *gocui.View) error {
	return gocui.ErrQuit
}

func main() {

	g, err := gocui.NewGui(gocui.OutputGrayScale, false)
	if err != nil {
		fmt.Printf("NewGui failed: %s\n", err.Error())
		return
	}
	defer g.Close()

	x, y := g.Size()
	fmt.Printf("size of window is %d, %d", x, y)

	test_view, err := g.SetView("test", 0, 0, 10, 10, 0)
	if gocui.IsUnknownView(err) {
		fmt.Fprintf(test_view, "hello world")
	} else if err != nil {
		fmt.Printf("SetView failed: %s\n", err.Error())
		return
	}

	_, err = g.SetCurrentView("test")
	if err != nil {
		fmt.Printf("SetCurrentView failed: %s\n", err.Error())
		return
	}

	err = g.SetKeybinding("test", gocui.KeyEnter, gocui.ModNone, quit)
	if err != nil {
		fmt.Printf("SetKeybinding failed: %s\n", err.Error())
		return
	}

	if err := g.MainLoop(); err != nil && !gocui.IsQuit(err) {
		fmt.Printf("g.MainLoop failed: %s\n", err.Error())
		return
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
}

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
