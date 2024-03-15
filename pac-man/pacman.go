package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
)

var maze []string

func main() {
	// initialize the game 



	// load ressources 
	err := loadMaze("maze01.txt")
	if err != nil {
		log. Println("failed to load maze:", err)
		return 
	}

	// game loop
	for {
		//update screen 
		printScreen()

		//process input


		//process movement

		//process collisions

		// check game over 

		break
	}
}

func loadMaze(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err 
	}
	defer f.Close() 

	scanner := bufio.NewScanner(f)
	for scanner. Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}
	return nil 
}

func printScreen() {
	for line:= range maze {
		fmt.Println(line)
	}
}
