package main

import (
	"fmt"
	"strings"
)

func main() {
	myMap := map[string]int{}

	str1 := "the quick brown fox jumps over the lazy dog"

	words := strings.Split(str1, " ")

	for _, word := range words {
		// fmt.Println(word)
		myMap[word] = myMap[word] + 1
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	val, err := myMap["the"]
	fmt.Println(val, err)
}
