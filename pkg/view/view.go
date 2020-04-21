package view

import (
	"fmt"
	"errors"
	"github.com/gdamore/tcell"
)


type View struct {
	SelectedRow, SelectedCol int
	screen tcell.Screen
}

func NewView(s tcell.Screen) *View {
	return &View{
		screen: s,
		SelectedRow: 0,
		SelectedCol: 0,
	}
}

func (v *View) MoveSelect(key tcell.Key) error {
	switch key {
	case tcell.KeyUp:
		v.SelectedRow = v.SelectedRow - 1
		if v.SelectedRow < 0 {
			v.SelectedRow = 2
		}
	case tcell.KeyRight:
		v.SelectedCol = v.SelectedCol + 1
		if v.SelectedCol > 2 {
			v.SelectedCol = 0
		}
	case tcell.KeyDown:
		v.SelectedRow = v.SelectedRow + 1
		if v.SelectedRow > 2 {
			v.SelectedRow = 0
		}
	case tcell.KeyLeft:
		v.SelectedCol = v.SelectedCol - 1
		if v.SelectedCol < 0 {
			v.SelectedCol = 2
		}
	default:
		return fmt.Errorf("%v is not a valid key", key)
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