package main

import (
	"fmt"
)

func main() {
	fmt.Println("Mimicking while loop")
	i := 0
	for i < 10 {
		fmt.Printf("\"for condition\" looping at %d\n", i)
		i++
	}
	fmt.Println()

	fmt.Println("An expanded for loop")
	for j := 0; j < 5; j++ {
		fmt.Printf("%d loop with \" for init; cond; inc { } \"\n", j)
	}

	fmt.Println()
	fmt.Println("A loop with a break")
	for k := 0; k < 15; k++ {
		fmt.Printf("looping at iteration %d ->", k)
		if k%3 == 0 {
			fmt.Printf("continue...\n")
			continue
		} else if k%11 == 0 {
			fmt.Printf("break---\n")
			break
		} else {
			fmt.Printf("No action\n")
		}
	}

}
