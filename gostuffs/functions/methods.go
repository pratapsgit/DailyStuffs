package main

import (
	"fmt"
)

type Person struct {
	fname string
	lname string
}

type Employee struct {
	Person
	empid  int
	empsal float32
}

func (e Employee) printDetails() {
	fmt.Printf("%d,%s,%s,%v\n", e.empid, e.fname, e.lname, e.empsal)
}

func main() {
	p1 := Employee{
		Person: Person{
			fname: "Pierie",
			lname: "Mustard",
		},
		empid:  2324,
		empsal: 459822.45,
	}

	fmt.Println(p1)
	p1.printDetails()
}
