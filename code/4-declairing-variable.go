package main

import "fmt"

func main() {

	var str1, str2 string = "hello", "world"

	fmt.Println(str1, str2) // (,) by default provides one space

	var (
		fname string = "Nilanjan"
		lname string = "Bhattacharjee"
	)

	fmt.Println(fname, lname)

	s := "this is short variable declairation" // dynamically typed, but data type is fixed at first assignment, we can't change

	s = "short variable value has changed"

	// s = 100 // This is an error

	fmt.Println(s)

}
