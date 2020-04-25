package game

type WinInfo interface {
	Letter() byte
	Squares() [3][2]int
}

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
