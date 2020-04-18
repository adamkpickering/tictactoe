package main

import (
	"fmt"
	//"os"
	//"time"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
)

func main() {

//	// set tcell up
//	s, err := tcell.NewScreen()
//	if err != nil {
//		fmt.Printf("main: %s", err.Error())
//		os.Exit(1)
//	}
//	err = s.Init()
//	if err != nil {
//		fmt.Printf("main: %s", err.Error())
//		os.Exit(1)
//	}
//	s.Clear()

	a := &views.Application{}
	w := NewTicTacToeWidget(a)
	a.SetRootWidget(w)
	a.Start()
	a.Wait()


//	width, height := s.Size()
//	wh := fmt.Sprintf("%d\u2318%d", width, height)
//
//	i := 0
//	for _, r := range wh {
//		s.SetContent(i, 0, r, nil, tcell.StyleDefault)
//		i = i + 1
//	}
//	s.Show()
//
//	time.Sleep(time.Second * 4)
//
//	s.Fini()
}


type TicTacToeWidget struct {
	application *views.Application
}

func NewTicTacToeWidget(a *views.Application) *TicTacToeWidget {
	return &TicTacToeWidget{
		application: a,
	}
}

func (t *TicTacToeWidget) Draw() {
	fmt.Println("Draw: not implemented\n")
}


func (t *TicTacToeWidget) Resize() {
	fmt.Println("Resize: not implemented\n")
}

func (t *TicTacToeWidget) HandleEvent(i tcell.Event) bool {
	fmt.Println("HandleEvent: implemented\n")
	switch e := i.(type) {
	case *tcell.EventKey:
		if e.Key() == tcell.KeyRune {
			fmt.Printf("%s\n", e.Rune())
		}
		if e.Key() == tcell.KeyEnter {
			t.application.Quit()
		}
		return true
	default:
		return false
	}
}

func (t *TicTacToeWidget) SetView(v views.View) {
	fmt.Println("SetView: not implemented\n")
}

func (t *TicTacToeWidget) Size() (int, int) {
	fmt.Println("Size: not implemented\n")
	return 0, 0
}

func (t *TicTacToeWidget) Watch(handler tcell.EventHandler) {
	fmt.Println("Watch: not implemented\n")
}

func (t *TicTacToeWidget) Unwatch(handler tcell.EventHandler) {
	fmt.Println("Unwatch: not implemented\n")
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
