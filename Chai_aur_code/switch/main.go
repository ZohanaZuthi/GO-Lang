package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Structs in Go are a way to group together related data and functions. They are similar to classes in other programming languages, but they do not have methods or inheritance. Structs are defined using the "type" keyword, followed by the name of the struct and the fields that it contains. Each field has a name and a type, and can be accessed using the dot notation.
func main() {
	fmt.Println("Welcome to switch statements in Go")
	rand.Seed(time.Now().UnixNano()) // generate a random age between 0 and 99
	age := rand.Intn(100) + 1
	fmt.Printf("Your age is: %v\n", age)
	switch age {
	case 18:
		fmt.Println("You are 18 years old")
		fallthrough
	case 21:
		fmt.Println("You are 21 years old")
		fallthrough
	case 65:
		fmt.Println("You are 65 years old")
		fallthrough
	default:
		fmt.Println("Your age is not 18, 21 or 65")
	}
}
