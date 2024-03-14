package main

import (
	"fmt"
	"os"
)

// func atoi(s string) (n int) {
// 	for _, v := range s {
// 		n = n*10 + int(rune(v)-48)
// 	}
// 	return
// }

// Define the board (3x3 grid)
var board [3][3]string
var currentPlayer string

// Initialize the board with empty cells
func initializeBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}
}

// Display the current state of the board
func displayBoard() {
	fmt.Println("Enter the resolution(0 1 2)")
	for i, row := range board {
		for j, squart := range row {
			fmt.Print(" ")
			switch squart {
			case " ":
				fmt.Print(" ")
			case "X":
				fmt.Print("X")
			case "O":
				fmt.Print("O")
			}
			if j != len(row)-1 {
				fmt.Print(" |")
			}
		}
		if i != len(board)-1 {
			fmt.Println("\n-----------")
		}
	}
	fmt.Println()
}

func getPlayerInfo(player string) (row, col int) {
	fmt.Printf("%v turn, Enter row (0-2): ", player)
	fmt.Scanln(&row)
	fmt.Printf("%v turn, Enter column (0-2): ", player)
	fmt.Scanln(&col)
	if !validInput(row, col) {
		fmt.Println("you just entered a location out of the gameBoard!")
		return getPlayerInfo(player)
	}
	if !spaceFree(row, col) {
		fmt.Println("the place is not empty")
		return getPlayerInfo(player)
	}
	return row, col
}

func validInput(row, col int) bool {
	return row >= 0 && row <= 2 && col >= 0 && col <= 2
}

func spaceFree(row, col int) bool {
	return board[row][col] == " "
}

func main() {
	initializeBoard()
	currentPlayer = os.Args[1]
	for {
		displayBoard()
		row, col := getPlayerInfo(currentPlayer)

		board[row][col] = currentPlayer

		end, winner := checkWinner()
		if end {
			displayBoard()
			fmt.Printf("Congratulation player %v, you are the winner", winner)
			break
		}
		if boarFull() {
			displayBoard()
			fmt.Println("End game, is a tie")
			break
		}
		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}
}
func boarFull() bool {
	for _, col := range board {
		for _, v := range col {
			if v == " " {
				return false
			}
		}
	}
	return true
}

func checkWinner() (bool, string) {
	for i := 0; i < 3; i++ {
		if board[i][0] != " " && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			return true, board[i][0]
		}
		if board[0][i] != " " && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			return true, board[0][i]
		}
	}
	if board[0][0] != " " && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return true, board[0][0]
	}
	if board[2][0] == board[1][1] && board[1][1] == board[0][2] && board[1][1] != " " {
		return true, board[1][1]
	}
	return false, ""
}
