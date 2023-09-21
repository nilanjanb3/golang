package main

import (
	"fmt"
)

func main() {
	i := 20
	var f1 float32 = float32(i)
	fmt.Println(i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%T \n", f1)
}
