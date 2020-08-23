package main

import (
	"fmt"
)

func main() {
	fn := func(x ...int) int {
		sum := 0

		for _, v := range x {
			sum += v
		}

		return sum
	}

	inp := []int{5, 3, 6}
	r := fn(inp...)
	fmt.Printf("We got a value %v\n", r)
}
