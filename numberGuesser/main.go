package main

import (
	"fmt"
	"math/rand"
//	"strconv"
)

func main() {
	fmt.Println("think about your secret number, it must be between 1 and 100")
	
	low:= 0
	high:= 101
	var input string
	n:= 0
	for input != "correct" {
		n++
		guess:= rand.Int()%high
		for guess <= low{
			guess = rand.Int()%high
		}
		fmt.Println("I guess the number is: ", guess)
		fmt.Scan(&input)
		if input == "low" {
			high = guess
		}else if input == "high" {
			low = guess
		}else if input != "correct" {
			fmt.Println("error in the input!")
			return
		}
	}
	fmt.Printf("I guessed the number after %d times", n)
}
