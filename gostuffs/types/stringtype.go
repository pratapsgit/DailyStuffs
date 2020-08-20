package main

import (
	"fmt"
)

func main() {
	var x string = "Hello wrold"
	fmt.Printf("x variable type is %T\n", x)
	fmt.Printf("Value in x is %s\n", x)

	b := []byte(x)
	fmt.Printf("Type of b is %T\n", b)
	fmt.Println(b)

	for i := 0; i < len(x); i++ {
		fmt.Printf("%#U ", b[i])
	}
	fmt.Println()

	for i, v := range x {
		fmt.Printf("%v - %v - %#x\n", i, v, v)
	}
}
