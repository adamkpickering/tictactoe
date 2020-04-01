package main

import (
	"fmt"
	"os"
	"adampickering.ca/tictactoe/pkg/game"
)

func main() {

	g := game.NewGame(nil)

	for {
		printGame(g.Board)
		fmt.Printf("It is player %c's turn.\n", g.Turn)
		fmt.Printf("Pick a square: ")

		input := make([]byte, 10)
		if _, err := os.Stdin.Read(input); err != nil {
			fmt.Fprintf(os.Stderr, "There was an error reading input: %s\n", err.Error())
			continue
		}

		square := int(input[0]) - 48
		if err := g.PlayTurn(square); err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			continue
		}

		result := g.CheckWin()
		switch result {
		case 'X':
			fmt.Printf("Player %c won. Good Game!\n", result)
			os.Exit(0)
		case 'O':
			fmt.Printf("Player %c won. Good Game!\n", result)
			os.Exit(0)
		case 1:
			fmt.Print("It is a draw. Good game!\n")
			os.Exit(0)
		}
	}
}

func printGame(board [9]byte) {
	fmt.Printf("\tBoard:\t\t\tLegend:\n")
	fmt.Printf("\t-------------------\t-------------------\n")
	fmt.Printf("\t|  %c  |  %c  |  %c  |\t|  0  |  1  |  2  |\n", board[0], board[1], board[2])
	fmt.Printf("\t-------------------\t-------------------\n")
	fmt.Printf("\t|  %c  |  %c  |  %c  |\t|  3  |  4  |  5  |\n", board[3], board[4], board[5])
	fmt.Printf("\t-------------------\t-------------------\n")
	fmt.Printf("\t|  %c  |  %c  |  %c  |\t|  6  |  7  |  8  |\n", board[6], board[7], board[8])
	fmt.Printf("\t-------------------\t-------------------\n")
}
