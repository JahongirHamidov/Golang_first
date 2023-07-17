package main

import (
	"fmt"
	"math"
)

// Define interface

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}
func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circl := Circle{x: 0, y: 0, radius: 5}
	rectangl := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle Area : %f\n", getArea(circl))
	fmt.Printf("Rectangle Area : %f\n", getArea(rectangl))

}
