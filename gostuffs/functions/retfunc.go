package main

import (
	"fmt"
)

const (
	NOTHING = iota
	ONEPOW  = iota
	SQUARE  = iota
	CUBE    = iota
	QUAD    = iota
	PENTA   = iota
	SIXER   = iota
)

func printString(name string) string {
	return fmt.Sprintf("Hello! %s", name)
}

func giveMea(powtype int) func(int) int {
	return func(x int) int {
		ret := 1
		for i := 0; i < powtype; i++ {
			ret *= x
		}
		return ret
	}
}

func main() {
	fmt.Println(printString("Pierie"))

	f := giveMea(SQUARE)
	r := f(4)
	fmt.Printf("Squared value we got %d\n", r)

	f = giveMea(CUBE)
	r = f(4)
	fmt.Printf("Cubed value we got %d\n", r)

	f = giveMea(QUAD)
	r = f(4)
	fmt.Printf("Quaded value we got %d\n", r)

	fmt.Printf("Sixer value we got for 2 is %d\n", giveMea(SIXER)(2))
}
