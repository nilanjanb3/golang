package main

import "fmt"

var globalVariable string = "This is a global variable"

func main() {

	city := "London" // This is a local variable

	{
		country := "UK"

		fmt.Println(city)
		fmt.Println(country)
	}

	fmt.Println(city)
	// fmt.Print(country) // This line has scope error
	fmt.Println(globalVariable)

}
