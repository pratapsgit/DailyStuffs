package main

import (
	"fmt"
)

func main() {
	fmt.Println("Delete demo")
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9}
	fmt.Println(a)
	fmt.Println(b)

	c := append(a, b...)
	fmt.Println(c)

	c = append(c[:2], c[4:]...)
	fmt.Println(c)
}
