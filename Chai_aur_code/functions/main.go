package main

import "fmt"

func main() {

	fmt.Println("welcome to functions in go")
	greeter()

	result := add(2, 3)
	fmt.Println("The result of addition is", result)
	proresult, _ := proadd(1, 2, 3, 4, 5)
	fmt.Println("The result of proadd is", proresult)

}

func add(a int, b int) int {
	return a + b
}
func proadd(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "The result of proadd is"
}

func greeter() {
	fmt.Println("Welcome to the greeter function")
}
