package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Go")
	for i := 0; i < 5; i++ {
		fmt.Printf("Value of i: %v\n", i)
	}
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for i := range days {
		fmt.Printf("The value is %v\n", days[i])
	}

	for index, day := range days {
		fmt.Printf("The index is %v and the value is %v\n", index, day)
	}
	rouge := 0

	for rouge < 9 {
		if rouge == 2 {
			goto start
		}

		fmt.Printf("Rouge is %v\n", rouge)
		rouge++
	}

	// go to sytax
start:
	fmt.Println("This is the start of the loop")

}
