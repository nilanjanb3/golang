package main

import (
	"fmt"
	"math"
)

type Student struct {
	name   string
	rollNo int32
	marks  []int
	grades map[string]string
}
type Point struct {
	x int
	y int
}

func getDistance(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2))
}

func main() {

	var nilanjan Student
	aman := new(Student)
	fmt.Println(nilanjan)
	fmt.Println(aman)
	fmt.Println("====================================")
	var p1 Point = Point{
		x: 3,
		y: 2,
	}

	fmt.Println(p1)
	fmt.Print(p1.x, p1.y, "\n")
	fmt.Println("====================================")

	p2 := new(Point)
	p2.x = 3
	p2.y = 8

	fmt.Println(getDistance(p1, *p2))
	fmt.Printf("%+v \n", *p2)
	fmt.Println("====================================")

}
