package main 

import (
	"os"
	"fmt"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		return 
	}
	sec := atoi(os.Args[1])
	for sec != 0 {
		fmt.Printf(" %.2d: %.2d\r", sec/60, sec%60)
		time.Sleep(1 * time.Second)
		sec--
	}
	fmt.Println("TimeOut")
}

func atoi(s string) (n int) {
	for _, v:= range s {
		digit := int(v - 48 ) 
		n = n * 10 + digit 
	}
	return 
}
