package main

import "fmt"

const a = 10

var p = 100

func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println(z)
	}

	add(5, 6)
	add(p, a)
}

func main() {
	// invoked function
	call()

	fmt.Println(a)

}

func init() {
	fmt.Println("This is the first function")

}
