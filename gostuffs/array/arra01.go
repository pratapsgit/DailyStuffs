package main

import (
	"fmt"
)

func main() {
	var myarr [5]int

	myarr[3] = 56

	fmt.Println(myarr[3])
	fmt.Println(myarr[0])
	fmt.Printf("Length oy array myarra is %d\n", len(myarr))
}
