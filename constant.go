package main

import "fmt"

func main() {
	const str = "Hello World"
	// we can not redefine after using str
	fmt.Println(str)

	const (
		num1 = 10
		str1 = "Golang"
	)

	fmt.Println(num1, str1)
}
