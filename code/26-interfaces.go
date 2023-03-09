package main

import "fmt"

type Shape interface {
	area() float64
}
type Rectangle struct {
	length  float64
	breadth float64
}

func (r Rectangle) area() float64 {
	return 2 * (r.length + r.breadth)
}

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return 4 * s.side
}

func printArea(shape Shape) {
	fmt.Println(shape.area())
}
func main() {
	r := Rectangle{
		length:  5,
		breadth: 10,
	}

	s := Square{
		side: 10,
	}

	printArea(r)
	printArea(s)
}
