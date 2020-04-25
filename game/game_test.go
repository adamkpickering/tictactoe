package game

import (
	"testing"
)

func TestPlayTurn(t *testing.T) {

	g := NewGame(nil)

	err := g.PlayTurn(0, 3)
	if err == nil {
		t.Fail()
	}

	err = g.PlayTurn(-1, 2)
	if err == nil {
		t.Fail()
	}

	err = g.PlayTurn(0, 1)
	if err != nil {
		t.Fail()
	}
}

func TestCheckWin(t *testing.T) {

	// test all win conditions
	{
		wc_array := [8][3][2]int{
			[3][2]int{
				[2]int{0, 0},
				[2]int{0, 1},
				[2]int{0, 2},
			},
			[3][2]int{
				[2]int{1, 0},
				[2]int{1, 1},
				[2]int{1, 2},
			},
			[3][2]int{
				[2]int{2, 0},
				[2]int{2, 1},
				[2]int{2, 2},
			},
			[3][2]int{
				[2]int{0, 0},
				[2]int{1, 0},
				[2]int{2, 0},
			},
			[3][2]int{
				[2]int{0, 1},
				[2]int{1, 1},
				[2]int{2, 1},
			},
			[3][2]int{
				[2]int{0, 2},
				[2]int{1, 2},
				[2]int{2, 2},
			},
			[3][2]int{
				[2]int{0, 0},
				[2]int{1, 1},
				[2]int{2, 2},
			},
			[3][2]int{
				[2]int{0, 2},
				[2]int{1, 1},
				[2]int{2, 0},
			},
		}
		board := [3][3]byte{}
		for _, letter := range [2]byte{'X', 'O'} {
			for _, wc := range wc_array {
				board = [3][3]byte{}
				board[wc[0][0]][wc[0][1]] = letter
				board[wc[1][0]][wc[1][1]] = letter
				board[wc[2][0]][wc[2][1]] = letter
				g := NewGame(board)
				result, _ := g.CheckWin()
				if result != Win {
					t.Fail()
				}
			}
		}
	}

	// test some other conditions
	{
		board := [3][3]byte{
			[3]byte{'X', 'X', 'O'},
			[3]byte{0, 0, 0},
			[3]byte{0, 0, 0},
		}
		g := NewGame(board)
		result, _ := g.CheckWin()
		if result != KeepPlaying {
			t.Fail()
		}
	}
	{
		board := [3][3]byte{
			[3]byte{'O', 0, 0},
			[3]byte{'X', 0, 0},
			[3]byte{'O', 0, 0},
		}
		g := NewGame(board)
		result, _ := g.CheckWin()
		if result != KeepPlaying {
			t.Fail()
		}
	}

	// test draw condition
	{
		board := [3][3]byte{
			[3]byte{'X', 'X', 'O'},
			[3]byte{'O', 'O', 'X'},
			[3]byte{'X', 'O', 'X'},
		}
		g := NewGame(board)
		result, _ := g.CheckWin()
		if result != Draw {
			t.Fail()
		}
	}

}
