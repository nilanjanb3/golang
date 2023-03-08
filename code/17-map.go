package main

import "fmt"

func main() {
	map1 := map[int]string{1: "a", 2: "b"}

	fmt.Println(map1)

	value, found := map1[1]

	fmt.Println(value, "=>", found)

	fmt.Println("=============================================")

	value, found = map1[11]
	fmt.Println(value, "=>", found) // if key is not present value will be default value of the data type

	fmt.Println("=============================================")
	// Adding a key, value

	map1[3] = "c"
	fmt.Println(map1)

	fmt.Println("=============================================")
	// modifying a value

	map1[3] = "z"
	fmt.Println(map1)

	fmt.Println("=============================================")
	// delete a key

	delete(map1, 3)
	fmt.Println(map1)

	fmt.Println("=============================================")
	// looping a map
	var key int
	for key, value = range map1 {
		fmt.Println(key, "=>", value)
	}
	fmt.Println("=============================================")
	// truncate a map
	// 1 using deleting each key, value
	// 2 re initializing with a new map

	map1 = map[int]string{}

	fmt.Println(map1)

}
