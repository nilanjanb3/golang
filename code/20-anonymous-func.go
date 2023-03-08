package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := func(a int, b int) int {
		return a + b
	}

	// fmt.Printf("%T \n", x)
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(x(3, 2))

}
