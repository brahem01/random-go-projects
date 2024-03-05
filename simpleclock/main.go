package main 

import (
	"fmt"
	"time"
) 

func main() {
	for {
		h, m, s:= time.Now().Clock()
		fmt.Printf("Time is: %.2d: %.2d: %.2d\r", h, m, s)
	}
}
