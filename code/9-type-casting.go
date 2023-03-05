package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 100
	var f float32 = float32(i)

	fmt.Println(f)

	var f2 float32 = 123.456

	var i2 int32 = int32(f2)

	fmt.Println(i2)

	fmt.Printf("%q \n", strconv.Itoa(i))

	var s string = "256"
	i3, err := strconv.Atoi(s)
	fmt.Printf("%v %v \n", i3, err)

	var s2 string = "256abc"
	i4, err2 := strconv.Atoi(s2)
	fmt.Printf("%v %v \n", i4, err2) // 0 Error

}
