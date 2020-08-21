package main

import (
	"fmt"
)

func main() {
	a := 45
	b := 56

	fmt.Printf("%d == %d ? Ans. %v\n", a, b, a == b)
	fmt.Printf("%d <= %d ? Ans. %v\n", a, b, a <= b)
	fmt.Printf("%d >= %d ? Ans. %v\n", a, b, a >= b)
	fmt.Printf("%d != %d ? Ans. %v\n", a, b, a != b)
	fmt.Printf("%d < %d ? Ans. %v\n", a, b, a < b)
	fmt.Printf("%d > %d ? Ans. %v\n", a, b, a > b)
}
