package main

import (
	"fmt"
	"math"
)

func main() {
	defer fmt.Println("Done!")
	c := Circle{radius: 12}
	fmt.Println("Area of the circle ", c.getArea())

	//method call with nil implementation
	var s Shape
	var circ Circle
	s = &circ

	fmt.Println("Call without any implementation object created ", s.getArea())

	fmt.Println(c)

}

type Shape interface {
	getArea() float64
}

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) String() string {
	return "This is a Circle class"
}
