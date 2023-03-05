package main

import "fmt"

func main() {

	for i := 0; i < 3; i++ {
		fmt.Println(i)

	}
	fmt.Println("=========================================================")

	// x := 0

	// Infinite loop
	// for {
	// 	fmt.Println(x)
	// 	x++
	// }
	fmt.Println("=========================================================")

	for i := 0; i <= 5; i++ {
		if i < 3 {
			continue
		} else if i == 3 {
			fmt.Println(3)
		} else {
			fmt.Println("else case: ", i)
		}
	}
}
