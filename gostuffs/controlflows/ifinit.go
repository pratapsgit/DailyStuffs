package main

import (
	"fmt"
)

func main() {
	if f := 67; f == 32 {
		fmt.Printf("We are in a 32 condition\n")
	} else {
		fmt.Println("We are in no 32 condition")
	}

	if f := 88; f == 88 {
		fmt.Printf("We are in a 88 condition\n")
	} else {
		fmt.Println("We are in no 88 condition")
	}
}
