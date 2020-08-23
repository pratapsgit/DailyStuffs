package main

import (
	"fmt"
)

func functionOne() {
	fmt.Println("Function one")
}

func functionTwo() {
	fmt.Println("Function two")
}

func main() {
	defer functionOne()
	functionTwo()
}
