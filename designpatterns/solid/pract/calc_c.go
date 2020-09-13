package main

import (
	"fmt"
)

func main() {
	v := 200000.0
	r := 0.05
	final_sum := 0.0
	n_y := 10
	t_w := n_y * 52

	final_sum = v
	for i := 1; i <= t_w; i++ {
		in := final_sum * r
		final_sum += in
		if i>1 && (i-1)%4 == 0{
			fmt.Println()
		}
		fmt.Printf("%3d. %20.2f\t", i, final_sum)
	}

	fmt.Printf("\nTotal sum : %.2f\n", final_sum)
}
