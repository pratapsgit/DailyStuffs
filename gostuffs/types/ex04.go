package main

import (
	"fmt"
)

var x int = 8

func main() {
	fmt.Printf("Value of x in dec %d bin %b hex %#x\n", x, x, x)
	y := x << 1
	fmt.Printf("Value of y holding x shifted to left by 1 dec %d bin %b hex %#x\n", y, y, y)
}
