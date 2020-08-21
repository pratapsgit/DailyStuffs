package main

import (
	"fmt"
)

const (
	_ = iota
	a = (2016 + iota)
	b = (2016 + iota)
	c = (2016 + iota)
	d = (2016 + iota)
)

func main() {
	fmt.Printf("First constant a : %v\n", a)
	fmt.Printf("second constant b : %v\n", b)
	fmt.Printf("Third constant c : %v\n", c)
	fmt.Printf("Fourth constant d : %v\n", d)
}
