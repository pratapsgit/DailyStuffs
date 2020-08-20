package main

import (
	"fmt"
)

const (
	a int     = 32
	b float32 = 45.67
	c string  = "Hello"
)

const (
	d = 32
	e = 45.67
	f = "Hello"
)

func main() {
	fmt.Printf("a has type %T\n", a)
	fmt.Printf("b has type %T\n", b)
	fmt.Printf("c has type %T\n", c)

	fmt.Printf("d has type %T\n", d)
	fmt.Printf("e has type %T\n", e)
	fmt.Printf("f has type %T\n", f)
}
