package main

import "fmt"

func main() {
	loginCount := 23
	var msg string
	if loginCount < 10 {
		msg = "Welcome to our website"
	} else if loginCount < 20 {
		msg = "Welcome back to our website"
	} else {
		msg = "Welcome again to our website"
	}
	fmt.Println(msg)
	// initialize and also check
	if num := 9; num%2 == 1 {
		fmt.Printf("%v is odd\n", num)
	}
	// if err != nil {
}
