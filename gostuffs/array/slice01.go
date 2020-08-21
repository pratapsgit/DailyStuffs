package main

import (
	"fmt"
)

func main() {
	x := []int{6, 2, 9, 34, 1}
	fmt.Println(len(x))
	for i, v := range x {
		fmt.Printf("Value at index %d is %d\n", i, v)
	}

	fmt.Println(x[1:])
	fmt.Println(x[:4])
	fmt.Println(x[1:4])

	for i := 0; i < len(x); i++ {
		fmt.Printf("Value at index %d is %d\n", i, x[i])
	}

	fmt.Println("Append demo")
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9}
	fmt.Println(a)
	fmt.Println(b)

	c := append(a, b...)
	fmt.Println(c)
}
