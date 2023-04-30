package main

import (
	"fmt"
	"math"
)

type Shape interface {
	getArea() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (r Rectangle) getArea() float64 {
	return r.length * r.length
}

func (c Circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s Shape) float64 {
	return s.getArea()
}

func main() {
	rect := Rectangle{length: 32, breadth: 54}
	circ := Circle{radius: 32}

	fmt.Println("Area of a rectangle ", getArea(rect))
	fmt.Println("Area of a circle ", getArea(circ))
}
