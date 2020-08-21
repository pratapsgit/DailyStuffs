package main

import (
	"fmt"
)

func main() {
	s := "Here"
	switch s {
	case "There":
		fmt.Println("There is a FOX hiding.")
	case "Here":
		fmt.Println("Here is the RABBIT")
	case "Another":
		fmt.Println("Another LION wants to be King")
	default:
		fmt.Println("Invalid request")
	}
}
