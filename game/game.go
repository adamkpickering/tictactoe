package game

import (
	"fmt"
)

const (
	KeepPlaying int = iota
	Draw
	Win
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

func (g *Game) PlayTurn(x, y int) error {

	if x < 0 || x > 2 || y < 0 || y > 2 {
		return fmt.Errorf("x %d y %d is not valid; must be in range 0-2", x, y)
	}

	if g.Board[y][x] != 0 {
		return fmt.Errorf("square %d,%d is already occupied", x, y)
	}

	g.Board[y][x] = g.Turn

	if g.Turn == 'X' {
		g.Turn = 'O'
	} else {
		g.Turn = 'X'
	}

	return nil
}

// Returns either KeepPlaying, Draw, or Win in first argument.
// If the first argument is equal to Win, a WinInfo structure is
// returned in the second argument; otherwise, the second argument is nil.
func (g *Game) CheckWin() (int, WinInfo) {

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
	for _, letter := range [2]byte{'X', 'O'} {
		for _, wc := range wc_array {
			win := true
			for _, coords := range wc {
				if g.Board[coords[0]][coords[1]] != letter {
					win = false
				}
			}
			if win {
				return Win, MuhWin{
					letter: g.Board[wc[0][0]][wc[0][1]],
					squares: wc,
				}
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
		return Draw, nil
	}

	// neither win nor draw
	return KeepPlaying, nil
}
