package game

import (
	"fmt"
)

type Game struct {
	Board [3][3]byte
	Turn  byte
}

func NewGame(i interface{}) *Game {
	switch v := i.(type) {
	case [3][3]byte:
		return &Game{
			Board: v,
			Turn:  'X',
		}
	default:
		return &Game{
			Board: [3][3]byte{},
			Turn:  'X',
		}
	}
}

func (g *Game) PlayTurn(row, col int) error {

	if row < 0 || row > 2 || col < 0 || col > 2 {
		return fmt.Errorf("row %d col %d is not valid; must be in range 0-2", row, col)
	}

	if g.Board[row][col] != 0 {
		return fmt.Errorf("square %d,%d is already occupied", row, col)
	}

	g.Board[row][col] = g.Turn

	if g.Turn == 'X' {
		g.Turn = 'O'
	} else {
		g.Turn = 'X'
	}

	return nil
}

func (g *Game) CheckWin() byte {

	// test for win condition
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
	win := true
	for _, letter := range [2]byte{'X', 'O'} {
		for _, wc := range wc_array {
			win = true
			for _, coords := range wc {
				if g.Board[coords[0]][coords[1]] != letter {
					win = false
				}
			}
			if win {
				return g.Board[wc[0][0]][wc[0][1]]
			}
		}
	}

	// test for draw condition
	draw := true
	for _, row := range g.Board {
		for _, square := range row{
			if square == 0 {
				draw = false
				break
			}
		}
	}
	if draw {
		return 1
	}

	return 0
}
