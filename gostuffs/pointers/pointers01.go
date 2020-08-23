package main

import (
	"fmt"
)

func detailFunc(v *int) {
	fmt.Println("got the value ", *v)
	fmt.Println("holds the address ", v)

	*v = 6

	fmt.Println("changed the value ", *v)
	fmt.Println("holds the address ", v)
}
func main() {
	x := 2

	fmt.Println("Value at x ", x)
	fmt.Println("Address of x ", &x)

	detailFunc(&x)

	fmt.Println("Value at x ", x)
	fmt.Println("Address of x ", &x)
	fmt.Println("dereference of x ", *&x)

}
