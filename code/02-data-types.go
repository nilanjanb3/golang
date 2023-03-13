package main

import "fmt"

func main() {
	str := "nilanjan" // string data type

	/*
		uint8, unit16, uint32, uint64 --> unsigned integer
		int8, int16, int32, int64 --> these can store both signed and unsigned integers

		int --> 4 byte in 32bit , 8 byte 64 bit machine
	*/
	x := 182 // dynamically typed

	var name string = "nilanjan" // statically typed variable
	var f1 float32 = 32.32
	var f2 float64 = 64.64
	var b bool = true
	var i1 int8 = 75

	fmt.Println(x)
	fmt.Println(str)
	fmt.Println(name)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(b)
	fmt.Print(i1)
}
