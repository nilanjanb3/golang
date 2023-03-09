package main

import "fmt"

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

func main() {

	var nilanjan Student
	aman := new(Student)
	fmt.Println(nilanjan)
	fmt.Println(aman)

	var p1 Point = Point{
		x: 3,
		y: 2,
	}

	fmt.Println(p1)

}
