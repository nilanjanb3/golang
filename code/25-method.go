package main

import "fmt"

var PI float64 = 3.145

type Circle struct {
	radius float64
	area   float64
}

func (c *Circle) calculateArea() { // if we don't use pointer then only instance value will be changed
	c.area = PI * c.radius * c.radius
}
func main() {

	circle1 := Circle{radius: 5}

	circle1.calculateArea()

	fmt.Printf("%+v", circle1)

}
