package main

import (
	"fmt"
)

const x int = 56
const y = 99

func main() {
	fmt.Printf("A typed constant : %d\n", x)
	fmt.Printf("A un-typed constant : %v\n", y)
}
