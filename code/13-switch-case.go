package main

import "fmt"

func main() {
	var a, b int = 100, 5

	switch a / b {
	case 10:
		fmt.Println(10)
	case 20:
		fmt.Println(20)
	case 30:
		fmt.Println(30)
	default:
		fmt.Println("Unknown")

	}

	fmt.Println("=========================================================")

	// fallthrough will execute the following case statemts as well
	day := "wednesday"
	switch day {
	case "monday":
		fmt.Println("monday")
	case "tuesday":
		fmt.Println("tuesday")
	case "wednesday":
		fmt.Println("wednesday")
		fallthrough
	case "thursday":
		fmt.Println("thursday")
		fallthrough
	case "friday":
		fmt.Println("friday")
	case "saturday", "sunday":
		fmt.Println("weekend")
	default:
		fmt.Println("default")
	}

	fmt.Println("=========================================================")

	// we can match multiple case at a time
	day = "saturday"
	switch day {
	case "monday":
		fmt.Println("monday")
	case "tuesday":
		fmt.Println("tuesday")
	case "wednesday":
		fmt.Println("wednesday")
	case "thursday":
		fmt.Println("thursday")
	case "friday":
		fmt.Println("friday")
	case "saturday", "sunday":
		fmt.Println("weekend")
	default:
		fmt.Println("default")
	}

	fmt.Println("=========================================================")

	// we can write switch with out expression as well

	switch {
	case a+b == 105:
		fmt.Println(105)

	case a+b == 110:
		fmt.Println(110)

	case a+b == 120:
		fmt.Println(120)

	default:
		fmt.Println("default")

	}

}
