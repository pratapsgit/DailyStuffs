package main

import (
	"fmt"
)

func printData(e ...int){
	fmt.Printf("Processing elements ")
	for _, v := range e{
		fmt.Printf(" %v", v)
	}
	fmt.Println()
}

func calculateSum(input ...int) (int, float64) {
	var sum int
	var average float64

	printData(input...)

	for _, v := range input {
		sum += v
	}

	average = float64(sum) / float64(len(input))

	return sum, average
}

//This is wrong, as the variadic argument should be the final argument
//func invalidFunction(input ...int, s string)

func main() {
	r1, r2 := calculateSum(1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Printf("Sum : %v Average : %v\n", r1, r2)

	xslice := []int{ 2, 3, 4, 5, 6, 7, 8 ,9}

	r1, r2 = calculateSum(xslice...)

	fmt.Printf("Sum : %v Average : %v\n", r1, r2)
}
