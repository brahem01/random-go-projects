package main 

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("you should enter the number of character in the password!!!")
		return 
	}
	password:= make([]byte, atoi(os.Args[1]))
	rand.Read(password)
	fmt.Printf("the password is: %x\n", password)
}

func atoi(s string) (n int) {
	for _, v:= range s {
		digit := int(v - 48)
		n = n * 10 + digit
	}
	return 
}
