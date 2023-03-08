package main

import "fmt"

func printName() {
	fmt.Println("nilanjan")
}

func printRoll() {
	fmt.Println(10)
}

func printCollege() {
	fmt.Println("IIT")
}
func main() {

	printName()
	defer printRoll() // defer will wait for surroundings to complete first
	printCollege()
}
