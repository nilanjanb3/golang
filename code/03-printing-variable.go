package main

import "fmt"

func main() {

	var fname string = "Nilanjan"
	var lname string = "Bhattacharjee"
	var address string = "Kolkata"
	var language string = "golang"
	var temperature float32 = 32.8

	fmt.Print(fname, " ", lname, "\n")
	fmt.Println(address)
	fmt.Println(language)

	fmt.Printf("first name is %s \n", fname) // %s formats in string
	fmt.Printf("last name is %v \n", lname)
	fmt.Printf("temperature is %v C \n", temperature) // %v formats the value in default

	/*
		format specifiers are like c language
		%q is for quoted character or string
		%T type of the value
		%t is for true or false
	*/

}
