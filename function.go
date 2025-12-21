package main

import "fmt"

func add(a int, b int) int {
	return a + b

}

// function takes another function and two integers
func function(fn func(a int, b int) int, a int, b int) {
	fmt.Println(fn(a, b))
}

func main() {
	result := add(10, 20)
	fmt.Println("Result:", result)

	function(add, 12, 3)

}
