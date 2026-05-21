package main

import (
	"fmt"

	"example.com/mathlib"
)

var a = 10

func main() {
	fmt.Println("Tela Mathai tel dao")
	var k = mathlib.Add(12, 3)
	fmt.Println(k)
	fmt.Println(a)

}

// init function doesn't take input or output
// at first global is called , then init function , and then main function
// it is a first function that is called
func init() {
	fmt.Println("This is the first function that is called")
	a = 20
}
