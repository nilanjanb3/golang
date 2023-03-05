package main

import "fmt"

func main() {

	// if we don't declare any variable it will catch default values

	var s string // ""
	var b bool
	var f float32
	var i int32
	var d float64

	fmt.Println(s)
	fmt.Println(b) // false
	fmt.Println(f) // 0
	fmt.Println(d) // 0
	fmt.Println(i) // 0

}
