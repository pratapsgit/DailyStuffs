package main

import (
	"fmt"
)

func sum(v ...int) int {
	s := 0

	for _, v := range v {
		s += v
	}

	return s
}

func sumCallback(f func(r ...int) int, eo bool, v ...int) int {
	a := []int{}

	for _, vv := range v {
		if vv%2 == 0 && eo == true {
			a = append(a, vv)
		} else if vv%2 != 0 && eo == false {
			a = append(a, vv)
		}
	}

	return f(a...)
}

func main() {
	ele := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("Sum of all numbers ", sum(ele...))
	e := sumCallback(sum, true, ele...)
	fmt.Printf("Sum of even numbers %d\n", e)
	o := sumCallback(sum, false, ele...)
	fmt.Printf("Sum of odd numbers %d\n", o)

}
