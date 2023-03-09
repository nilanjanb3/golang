package main

import "fmt"

func main() {
	x := 10
	var ptr *int = &x

	fmt.Println(ptr, *ptr)

}
