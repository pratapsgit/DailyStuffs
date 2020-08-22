package main

import (
	"fmt"
)

type Person struct {
	fname string
	lname string
	age   int
}

func main() {
	per1 := Person{
		fname: "Peirie",
		lname: "Musti",
		age:   40,
	}

	per2 := Person{
		fname: "Jerrie",
		lname: "Musti",
		age:   3,
	}

	fmt.Println(per1)
	fmt.Println(per2)

	fmt.Printf("%-20s:%s:%10d\n", per1.lname, per1.fname, per1.age)
	fmt.Printf("%-20s:%s:%10d\n", per2.lname, per2.fname, per2.age)
}
