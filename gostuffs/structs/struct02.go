package main

import (
	"fmt"
)

type Person struct {
	fname string
	lname string
	age   int
}

type Employee struct {
	Person //It's the type
	edesg  string
	eid    int
}

func main() {
	p1 := Employee{
		Person: Person{ //Use the actual type to create the instance
			fname: "Pierie",
			lname: "Multi",
			age:   40,
		},
		edesg: "Manager",
		eid:   7383,
	}

	fmt.Println(p1)
	fmt.Printf("First Name:%32s\nLast Name:%32s\nAge:%10d\nDesignation:%32s\nEmployee Id:%10d\n", p1.fname, p1.lname, p1.age, p1.edesg, p1.eid)
}
