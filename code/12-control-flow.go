package main

import "fmt"

func main() {
	/*
		same if-else like c, c++ or java
		we have to maintain the following structure, it's sensitive
	*/

	var a, b string = "kolkata", "kolkata"
	if a == b {
		fmt.Println("strings are equal")
	} else {
		fmt.Println("strings are not equal")
	}
	fmt.Println("thank you!")

}
