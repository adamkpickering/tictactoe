package game

import (
	"adampickering.ca/tictactoe/pkg/game"
	"testing"
)

func TestPlayTurn(t *testing.T) {

	g := game.NewGame(nil)

	err := g.PlayTurn(9)
	if err == nil {
		t.Fail()
	}

	err = g.PlayTurn(-1)
	if err == nil {
		t.Fail()
	}

	err = g.PlayTurn(1)
	if err != nil {
		t.Fail()
	}
}

func TestCheckWin(t *testing.T) {

	// test all win conditions
	{
		win_conditions := [8][3]int{
			[3]int{0, 1, 2},
			[3]int{3, 4, 5},
			[3]int{6, 7, 8},
			[3]int{0, 3, 6},
			[3]int{1, 4, 7},
			[3]int{2, 5, 8},
			[3]int{0, 4, 8},
			[3]int{2, 4, 6},
		}
		board := [9]byte{}
		for _, letter := range [2]byte{'X', 'O'} {
			for _, condition := range win_conditions {
				board = [9]byte{}
				board[condition[0]] = letter
				board[condition[1]] = letter
				board[condition[2]] = letter
				g := game.NewGame(board)
				winner := g.CheckWin()
				if winner != letter {
					t.Fail()
				}
			}
		}
	}

	// test some other conditions
	{
		board := [9]byte{'X', 'O', 'X'}
		g := game.NewGame(board)
		winner := g.CheckWin()
		if winner != 0 {
			t.Fail()
		}
	}
	{
		board := [9]byte{'O', 0, 0, 'X', 0, 0, 'O'}
		g := game.NewGame(board)
		winner := g.CheckWin()
		if winner != 0 {
			t.Fail()
		}
	}
}
