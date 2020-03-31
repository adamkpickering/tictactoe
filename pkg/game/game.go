package game

import (
	"errors"
	"fmt"
)

type Game struct {
	board [9]byte
	turn  byte
}

func NewGame(i interface{}) *Game {
	switch v := i.(type) {
	case [9]byte:
		return &Game{
			board: v,
			turn:  'X',
		}
	default:
		return &Game{
			board: [9]byte{},
			turn:  'X',
		}
	}
}

func (g *Game) PlayTurn(number int) error {

	if number < 0 || number > 8 {
		return errors.New(fmt.Sprintf("%d is not a valid number; choose a number from 0 to 8", number))
	}

	if g.board[number] != 0 {
		return errors.New(fmt.Sprintf("square %d is already occupied", number))
	}

	g.board[number] = g.turn
	if g.turn == 'X' {
		g.turn = 'O'
	} else {
		g.turn = 'X'
	}

	return nil
}

func (g *Game) CheckWin() byte {

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
			if g.board[c[0]] == letter && g.board[c[1]] == letter && g.board[c[2]] == letter {
				return g.board[c[0]]
			}
		}
	}

	return 0
}
