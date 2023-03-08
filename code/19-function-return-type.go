package main

import "fmt"

func greet(name string) {
	fmt.Println(name)
}

func swap(a int, b int) (int, int) {
	return b, a
}

func swap2(a int, b int) (c int, d int) {
	c = b
	d = a
	return
}

func variadicFunc(numbers ...int) {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	fmt.Println(sum)
}
func main() {
	greet("nilanjan")

	a, b := swap(3, 5)

	fmt.Println(a, b)

	fmt.Println("====================================")

	a, b = swap2(10, 20)

	fmt.Println(a, b)
	fmt.Println("====================================")

	variadicFunc(1, 2, 3, 4)

}
