package main

import "fmt"

func main() {
	var fname string
	var rollNo int16

	fmt.Println("enter name and roll no")
	// fmt.Scanf("%s %d", &fname, &rollNo)
	// fmt.Println("Name: ", fname, "Roll: ", rollNo)

	count, err := fmt.Scanf("%s %d", &fname, &rollNo)

	fmt.Println(count, err)
	/*
		nilanjan 3
		2 <nil>

		nilanjan hi
		1 expected integer
	*/

}
