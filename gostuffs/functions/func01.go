package main

import (
	"fmt"
)

func zeroArgFunction() {
	fmt.Println("Zero argument function")
}

func oneArgFunction(arg1 string) {
	fmt.Printf("One argument function with argument value : %s\n", arg1)
}

func twoArgFunction(arg1 string, arg2 bool) {
	fmt.Printf("Two argument function with argument value : %s : %v\n", arg1, arg2)
}

func returnArgFunction(arg1 string, arg2 int) string {
	fmt.Printf("Two argument function with argument value : %s : %b\n", arg1, arg2)
	r := fmt.Sprintf("%s-%v", arg1, arg2)

	return r
}

func main() {
	//func (associate AssociateType) fun-identifier(arg0 arg0-type, arg1 arg1-type, ...) (return-type) { }
	zeroArgFunction()
	oneArgFunction("Jumper")
	twoArgFunction("Elite", true)
	r := returnArgFunction("Boxer", 78)
	fmt.Printf("We go a value of %s\n", r)
}
