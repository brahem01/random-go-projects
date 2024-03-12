package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

// Define the board (3x3 grid)
var board [3][3]string

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
	for i,row:= range board{
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
		if i != len(board) -1 {
			fmt.Println("\n-----------")
		}
	}
	fmt.Println()
}

// Get the next move from the player
func getNextMove(player string) (int, int) {
	fmt.Printf("%s's turn. Enter row (0-2): ", player)
	rowInput := readInput()
	row, _ := strconv.Atoi(rowInput)

	fmt.Printf("%s's turn. Enter column (0-2): ", player)
	colInput := readInput()
	col, _ := strconv.Atoi(colInput)

	return row, col
}

// Read input from the console
func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input[:len(input)-1]
}

// Main function
func main() {
	initializeBoard()
	currentPlayer := "X"

	for {
		displayBoard()
		row, col := getNextMove(currentPlayer)

		if board[row][col] != " " {
			fmt.Println("Cell already occupied. Try again.")
			continue
		}

		board[row][col] = currentPlayer

		// Switch players
		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
		if someoneWin() {
			fmt.Println("the winner is: ", CheckWin())
			displayBoard()
			break
		}
		if fullBoard() {
			fmt.Println("is a ta3adol")
			displayBoard()
			break
		}
	}
}

func CheckWin() string {
	// Check rows
	for i := 0; i < 3; i++ {
		if board[i][0] != " " && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0]
		}
	}

	// Check columns
	for j := 0; j < 3; j++ {
		if board[0][j] != " " && board[0][j] == board[1][j] && board[1][j] == board[2][j] {
			return board[0][j]
		}
	}

	// Check diagonals
	if board[0][0] != " " && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] != " " && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}

	return "" // No winner
}


func someoneWin() bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if board[i][0] != " " && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return true
		}
	}

	// Check columns
	for j := 0; j < 3; j++ {
		if board[0][j] != " " && board[0][j] == board[1][j] && board[1][j] == board[2][j] {
			return true
		}
	}

	// Check diagonals
	if board[0][0] != " " && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return true
	}
	if board[0][2] != " " && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return true
	}

	return false
}


func fullBoard() bool {
	for _, v:= range board{
		for _, char := range v {
			if char == " "  {
				return false
			}
		}
	}
	return true
}

