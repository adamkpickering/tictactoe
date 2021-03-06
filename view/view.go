// Package view contains all of the code that is concerned with communicating
// game information to the players, other than that which is in the main loop.
package view

import (
	"fmt"
	"strings"
	"github.com/gdamore/tcell"
	"github.com/adamkpickering/tictactoe/game"
)

// Represents the UI of the game; contains all UI state.
type View struct {
	SelectedY, SelectedX int
	screen tcell.Screen
}

// Used to create new View structs.
func NewView(s tcell.Screen) *View {
	return &View{
		screen: s,
		SelectedY: 0,
		SelectedX: 0,
	}
}

// Implements the logic of what happens when the user presses a key.
func (v *View) MoveSelect(key tcell.Key) error {
	switch key {
	case tcell.KeyUp:
		v.SelectedY = v.SelectedY - 1
		if v.SelectedY < 0 {
			v.SelectedY = 2
		}
	case tcell.KeyRight:
		v.SelectedX = v.SelectedX + 1
		if v.SelectedX > 2 {
			v.SelectedX = 0
		}
	case tcell.KeyDown:
		v.SelectedY = v.SelectedY + 1
		if v.SelectedY > 2 {
			v.SelectedY = 0
		}
	case tcell.KeyLeft:
		v.SelectedX = v.SelectedX - 1
		if v.SelectedX < 0 {
			v.SelectedX = 2
		}
	default:
		return fmt.Errorf("%v is not a valid key", key)
	}
	return nil
}

// Given a game.Game structure, turns the state of the game described
// in that structure into its UI representation.
func (v *View) Draw(g *game.Game) {

	// get center of display, and top left corners of widgets
	width, height := v.screen.Size()
	center_x := width/2
	center_y := height/2
	board_x := center_x - 8
	board_y := center_y - 10
	messages_x := center_x - 21
	messages_y := center_y - 2

	result, win_info := g.CheckWin()
	selection := [][2]int{}
	message_lines := []string{}
	switch result {
	case game.KeepPlaying:
		selection = [][2]int{
			[2]int{v.SelectedX, v.SelectedY},
		}
		message_lines = []string{
			fmt.Sprintf("It is player %c's turn", g.Turn),
			"Use the arrow keys to select a square",
			"Press enter to choose a square",
			"Press q to exit",
		}
	case game.Draw:
		message_lines = []string{
			"It's a draw!",
			"Press q to exit",
		}
	case game.Win:
		squares := win_info.Squares()
		selection = squares[:]
		message_lines = []string{
			fmt.Sprintf("Player %c won!", win_info.Letter()),
			"Press q to exit",
		}
	}

	// write screen
	v.screen.Clear()
	v.writeBoard(board_x, board_y, g.Board, selection)
	v.writeCenteredLines(messages_x, messages_y, 40, message_lines)
	v.screen.Show()
}

// Writes each string in lines to a box that is width characters wide,
// centering each line within the box. The box starts at coordinates
// given by x and y, where x is the number of characters from the left
// side of the screen and y is the number of characters from the top
// of the screen.
func (v *View) writeCenteredLines(x, y, width int, lines []string) {
	for j, line := range lines {
		line_length := len([]rune(line))
		difference := width - line_length
		var new_line string
		if difference < 0 {
			new_line = line[0:width]
		} else {
			padding := difference/2
			b := strings.Builder{}
			for i := 0; i < padding; i++ {
				b.WriteByte(' ')
			}
			b.WriteString(line)
			new_line = b.String()
		}
		i := 0
		for _, character := range new_line {
			v.screen.SetContent(x + i, y + j, character, nil, tcell.StyleDefault)
			i = i + 1
		}
	}
}

// Given a [3][3]byte that represents a board, writes the game board with the top
// left corner at the square given by x and y, where x is the number of characters
// from the left side of the screen and y is the number of characters from the
// top of the screen. Also takes highlightSquares, which is a slice of coordinates
// of squares to highlight.
func (v * View) writeBoard(x, y int, board [3][3]byte, highlightSquares [][2]int) {

	// put X's and O's into board
	view_board := v.getFreshBoard()
	for i, row := range board {
		for j, val := range row {
			if val != 0 {
				view_board[ i*2 + 1 ][ j*4 + 2].Rune = rune(val)
			}
		}
	}

	// set the color for selected cell to white
	select_style := tcell.StyleDefault.Background(tcell.ColorWhite)
	for _, square := range highlightSquares {
		view_board[ square[1]*2 + 1 ][ square[0]*4 + 1 ].Style = select_style
		view_board[ square[1]*2 + 1 ][ square[0]*4 + 2 ].Style = select_style
		view_board[ square[1]*2 + 1 ][ square[0]*4 + 3 ].Style = select_style
	}

	// write data to the screen
	for j, row := range view_board {
		for i, cell := range row {
			v.screen.SetContent(i + x, j + y, cell.Rune, nil, cell.Style)
		}
	}
}

// Returns a structure that represents an empty tic tac toe board. This is
// useful because it eliminates the boilerplate of writing the characters
// that are identical in each rewrite of the board - the calling function
// can instead worry about the characters that represent the state of the
// game.
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
