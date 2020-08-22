package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		fname string
		lname string
		age   int
	}{
		fname: "Pierie",
		lname: "Mutahi",
		age:   43,
	}

	fmt.Println("Displaying data from a anonymous struct")
	fmt.Println(p1)
}
