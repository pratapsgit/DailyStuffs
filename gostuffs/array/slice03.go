package main

import (
	"fmt"
)

func main() {
	fmt.Println("make demo")
	a := make([]int, 15, 20)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a[10] = 55
	a = append(a, 56)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
