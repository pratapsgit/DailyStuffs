package main

import (
	"fmt"
)

type Shape interface {
	setHeight(h int)
	setWidth(w int)
	getWidth() int
	getHeight() int
}

type Rectangle struct {
	height, width int
}

func (r *Rectangle) setHeight(h int) {
	r.height = h
}

func (r *Rectangle) setWidth(w int) {
	r.width = w
}

func (r *Rectangle) getWidth() int {
	return r.width
}

func (r *Rectangle) getHeight() int {
	return r.height
}

func CalculateArea(s Shape) int {
	w := s.getWidth()
	s.setHeight(10)

	areaLocal := 10 * w
	areaActual := s.getWidth() * s.getHeight()

	fmt.Print("Area Local : ", areaLocal, " Area Actual : ", areaActual, "\n")

	return areaActual
}

type Square struct{
	Rectangle
}

func getSquare(s int) *Square{
	sq := Square{}
	sq.width = s
	sq.height = s

	return &sq
}

func (s *Square) setHeight(h int) {
	s.height = h
	s.width = h
}

func (s *Square) setWidth(w int) {
	s.height = w
	s.width = w
}

func main() {

	rect := &Rectangle{10, 20}

	CalculateArea(rect)

	sqr :=  getSquare(7)

	CalculateArea(sqr)
}
