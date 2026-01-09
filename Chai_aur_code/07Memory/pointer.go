package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in Go")
	myNumber := 23
	var ptr *int = &myNumber
	fmt.Println("Value of myNumber:", myNumber)
	fmt.Println("Value of ptr (address of myNumber):", ptr)
	fmt.Println("Value at the address stored in ptr:", *ptr)
	*ptr = *ptr + 20
	fmt.Println("New value of myNumber after updating via ptr:", myNumber)
}
