package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	default:
		fmt.Println("i is not 1 or 2")
	}

	// multiple condition on switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is weekend")
	default:
		fmt.Println("it is a weekday")
	}

	// declaring a function
	// interface{} can take any type
	whoAmI := func(name interface{}) {
		switch v := name.(type) {
		case int:
			fmt.Println("I am an integer", v)
		case string:
			fmt.Println("I am a string", v)
		default:
			fmt.Println("I am of a different type")
		}
	}
	whoAmI(10)

}
