package main

import "fmt"

const PI = 3.14

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return PI * c.radius * c.radius
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (r Rectangle) area() float64 {
	return r.breadth * r.length
}

type Shape interface {
	area() float64
}

func main() {
	c := Circle{radius: 7}
	fmt.Println(c.area())

	var r Rectangle
	r.length = 3
	r.breadth = 4

	fmt.Println(r.area())

	shapes := []Shape{c, r}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
