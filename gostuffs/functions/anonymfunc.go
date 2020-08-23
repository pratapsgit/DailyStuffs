package main

import (
	"fmt"
)

func main() {
	func(d int) {
		fmt.Printf("Value we have now : %v\n", d)
	}(67)
}
