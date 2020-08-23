package main

import (
	"fmt"
)

type Person struct {
	fname string
	lname string
}

type SalariedEmployee struct {
	Person
	empid  int
	empsal float32
}

func (p Person) describe() {
	fmt.Printf("%v:%v\n", p.lname, p.fname)
}

func (e SalariedEmployee) describe() {
	fmt.Printf("%d,%s,%s,%v\n", e.empid, e.fname, e.lname, e.empsal)
}

type Employee interface {
	describe()
}

func acceptEmployees(e Employee) {
	switch e.(type) {
	case Person:
		fmt.Println("We found the person ", e.(Person))
	case Employee:
		fmt.Println("We found the salaried employee ", e.(SalariedEmployee))
	}
}

func main() {
	p1 := SalariedEmployee{
		Person: Person{
			fname: "Pierie",
			lname: "Mustard",
		},
		empid:  2324,
		empsal: 459822.45,
	}

	p2 := Person{
		fname: "Bolero",
		lname: "Mahindra",
	}

	acceptEmployees(p1)
	acceptEmployees(p2)
}
