package main

import "fmt"

func main() {

	var names [5]string
	var numbers [3]int = [3]int{1, 2, 3}
	fmt.Println(names)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(numbers[1])
	fmt.Println("====================================")
	names = [5]string{"a", "b", "c", "d", "e"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	fmt.Println("====================================")

	for index, element := range names {
		fmt.Println(index, "=>", element)
	}
	// fmt.Println(cap(names))
	fmt.Println("====================================")

	mat := [2][2]int{{1, 2}, {3, 4}}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			fmt.Print(mat[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println("====================================")
}
