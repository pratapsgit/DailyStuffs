package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Pastero": 50,
		"Samero":  13,
	}

	fmt.Println(m)

	fmt.Println("Check for the key present or not")
	if v, ok := m["Morasi"]; ok {
		fmt.Println("Morasi is present")
		fmt.Println(v)
	} else {
		fmt.Println("Morasi is absent")
	}
}
