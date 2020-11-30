package game

// An interface that represents win info.
type WinInfo interface {
	// Returns the letter of the player that won.
	Letter() byte

	// Returns the squares that contain the line that the
	// player won with. Can be used for (for example)
	// highlighting them for better user feedback of results.
	Squares() [3][2]int
}

// An implementation of WinInfo.
type MuhWin struct {
	letter byte
	squares [3][2]int
}

func (mw MuhWin) Letter() byte {
	return mw.letter
}

func (mw MuhWin) Squares() [3][2]int {
	return mw.squares
}
