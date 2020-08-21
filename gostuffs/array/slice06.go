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

	fmt.Println("Lets add an element to the existing map")
	m["Protero"] = 78
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("Lets delete an element from the existing map")
	delete(m, "Protero")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
