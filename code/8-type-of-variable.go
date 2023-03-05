package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := "hello world"
	fmt.Printf("%T\n", s) // string

	fmt.Printf("%v\n", reflect.TypeOf(s))

}
