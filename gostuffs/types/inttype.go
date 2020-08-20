package main

import (
	"fmt"
)

var x int32
var y rune
var a uint8
var b bool

func main() {

	fmt.Printf("rune type %T\n", y)
	fmt.Printf("int32 type %T\n", x)
	fmt.Printf("uint8 type %T\n", a)
	fmt.Printf("bool type %T\n", b)
}
