package main

import "fmt"

func main() {
	var arr [10]int = [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	slice1 := arr[2:5]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	fmt.Println("====================================")

	slice2 := make([]int, 5, 10)
	fmt.Println(len(slice2)) // length --> numbers of element present
	fmt.Println(cap(slice2)) // capacity --> underlying capacity or the array

	fmt.Println("====================================")

	slice1[0] = 99 // change on slice affect the underlying array
	fmt.Println(slice1)
	fmt.Println(arr)

	fmt.Println("====================================")

	slice2 = append(slice2, 10, 20, 30)
	fmt.Println(slice2)
	fmt.Println(len(slice2)) // append will modify length
	fmt.Println(cap(slice2)) // append will modify capacity

	fmt.Println("====================================")
	slice3 := make([]int, 0, 0)
	slice3 = append(slice3, slice1...) // adding a slice to another
	fmt.Println(slice3)
	fmt.Println("====================================")

	src := []int{1, 2, 3}
	fmt.Println(src)

	dest := make([]int, 3) // src & dest type and size matters while copying

	num := copy(dest, src)

	fmt.Println(dest)
	fmt.Println(num)

}
