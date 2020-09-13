package main

import (
	"fmt"
)

type Color int

const (
	blue Color = iota
	red
	green
	brown
	white
	purple
	black
	yellow
)

var colorNames = map[Color]string {
	red: "red",
	green: "green",
	blue: "blue",
	brown: "brown",
	black: "black",
	yellow: "yellow",
	purple: "purple",
	white: "white",
}

func getColorName(color Color) string{
	if val, ok :=  colorNames[color]; ok {
		return val
	}

	return "none"
}

type Size int

const (
	small Size = iota
	medium
	large
)

var sizeChart = map[Size]string {
	small: "small",
	medium: "medium",
	large: "large",
}

func getSizeName(size Size) string{
	if val, ok :=  sizeChart[size]; ok {
		return val
	}

	return "none"
}

type VegNFruit struct {
	name  string
	color Color
	size  Size
	price float64
}

func (v *VegNFruit) toString() string {
	return fmt.Sprintf("%-12s%8s%8v%8.2f", v.name, getColorName(v.color), getSizeName(v.size), v.price)
}

var totalItems int = 0

type VeggieBasket struct {
	basket []*VegNFruit
}

func (v *VeggieBasket) addItem(item *VegNFruit) int {
	totalItems = totalItems + 1
	v.basket = append(v.basket, item)
	return totalItems
}

func (v *VeggieBasket) displayItem() {
	fmt.Println("Items in basket")
	for i, v := range v.basket {
		fmt.Printf("%02d. %s\n", i, v.toString())
	}
}

type Filter struct {
}

func (f *Filter) filterByColor(v *VeggieBasket, color Color) []*VegNFruit{
	//result := []*VegNFruit{}
	result := make([]*VegNFruit, 0)

	for _, val := range v.basket {
		if val.color == color {
			result = append(result, val)
		}
	}

	return result
}


func (f *Filter) filterBySize(v *VeggieBasket, size Size) []*VegNFruit{
	result := make([]*VegNFruit, 0)

	for _, val := range v.basket{
		if val.size == size {
			result = append(result, val)
		}
	}

	return result
}

func (f *Filter) filterByColorNSize(v *VeggieBasket, color Color, size Size) []*VegNFruit{
	result := make([]*VegNFruit, 0)

	for _, val := range v.basket {
		if val.color == color && val.size == size {
			result = append(result, val)
		}
	}

	return result
}


type Properties interface{
	isValid(v *VegNFruit) bool
}

type ColorProperties struct{
	color Color
}

func (c ColorProperties) isValid(v *VegNFruit) bool {
	return c.color == v.color
}

type SizeProperties struct{
	size Size
}

func (c SizeProperties) isValid(v *VegNFruit) bool {
	return c.size == v.size
}


func main() {
	apple := VegNFruit{name: "apple", color: red, size: medium, price: 109.67}
	orange := VegNFruit{name: "orange", color: yellow, size: medium, price: 105.45}
	capsicum := VegNFruit{name: "capsicum", color: green, size: medium, price: 85.15}
	papaya := VegNFruit{name: "papaya", color: green, size: large, price: 98.55}
	tomato := VegNFruit{name: "tomato", color: red, size: medium, price: 56.25}

	v := VeggieBasket{}

	v.addItem(&apple)
	v.addItem(&orange)
	v.addItem(&papaya)
	v.addItem(&tomato)
	v.addItem(&capsicum)

	v.displayItem()

	f := Filter{}

	fmt.Println("\nGreen color fruits and vegetables")
	for i, vc := range f.filterByColor(&v, green) {
		fmt.Printf("%02d. %s\n", i, vc.toString())
	}

	fmt.Println("\nRed color fruits and vegetables")
	for i, vc := range f.filterByColor(&v, red) {
		fmt.Printf("%02d. %s\n", i, vc.toString())
	}

	fmt.Println("\nMedium size fruits and vegetables")
	for i, vc := range f.filterBySize(&v, medium) {
		fmt.Printf("%02d. %s\n", i, vc.toString())
	}

	fmt.Println("\nFruits and vegetables of size meium and color green")
	for i, vc := range f.filterByColorNSize(&v, green, medium) {
		fmt.Printf("%02d. %s\n", i, vc.toString())
	}

}
