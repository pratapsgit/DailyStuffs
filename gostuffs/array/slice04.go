package main

import (
	"fmt"
)

func main() {
	s1 := []string{"Yureki", "Mokaira", "Bangalore", "Kaggadasa"}
	s2 := []string{"Soriminoa", "monasnta", "Bangalore", "California"}

	s := [][]string{s1, s2}
	fmt.Println(s)
}
