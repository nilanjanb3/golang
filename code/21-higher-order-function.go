package main

import "fmt"

var PI float64 = 3.14

func getArea(r float64) float64 {
	return PI * r * r
}

func getPerimeter(r float64) float64 {
	return 2 * PI * r
}

func higherOrder(r float64, calc func(r float64) float64) {
	fmt.Println(calc(r))
}

func anotherHigherOrderFunction() func() {

	return func() {
		fmt.Println("Hello Go!")
	}
}
func main() {
	action := map[int]func(r float64) float64{
		1: getArea,
		2: getPerimeter,
	}

	var input int
	var r float64
	fmt.Println("enter option and radious carefully")
	fmt.Scanf("%d %f", &input, &r)

	fmt.Println(input, r)

	higherOrder(r, action[input])

	fmt.Println("=============================================")

	x := anotherHigherOrderFunction()
	x()

}
