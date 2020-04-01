package game

import (
	"errors"
	"fmt"
)

type Game struct {
	Board [9]byte
	Turn  byte
}

func NewGame(i interface{}) *Game {
	switch v := i.(type) {
	case [9]byte:
		return &Game{
			Board: v,
			Turn:  'X',
		}
	default:
		return &Game{
			Board: [9]byte{},
			Turn:  'X',
		}
	}
}

func (g *Game) PlayTurn(number int) error {

	if number < 0 || number > 8 {
		return errors.New(fmt.Sprintf("%d is not a valid number; choose a number from 0 to 8", number))
	}

	if g.Board[number] != 0 {
		return errors.New(fmt.Sprintf("square %d is already occupied", number))
	}

	g.Board[number] = g.Turn
	if g.Turn == 'X' {
		g.Turn = 'O'
	} else {
		g.Turn = 'X'
	}

	return nil
}

func (g *Game) CheckWin() byte {

	// test for win condition
	wc := [8][3]int{
		[3]int{0, 1, 2},
		[3]int{3, 4, 5},
		[3]int{6, 7, 8},
		[3]int{0, 3, 6},
		[3]int{1, 4, 7},
		[3]int{2, 5, 8},
		[3]int{0, 4, 8},
		[3]int{2, 4, 6},
	}
	for _, letter := range [2]byte{'X', 'O'} {
		for _, c := range wc {
			if g.Board[c[0]] == letter && g.Board[c[1]] == letter && g.Board[c[2]] == letter {
				return g.Board[c[0]]
			}
		}
	}

	// test for draw condition
	draw := true
	for _, element := range g.Board {
		if element == 0 {
			draw = false
			break
		}
	}
	if draw {
		return 1
	}

	return 0
}
