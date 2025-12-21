package main

import "fmt"

func update(val *int) {
	*val = *val + 10
}
func main() {

	a := 10 // a normal variable
	p := &a // p is a pointer to a
	fmt.Println("a:", a)
	fmt.Println("p:", p)   // memory address
	fmt.Println("*p:", *p) // value at that address

	*p = 20 // modify value through the pointer
	fmt.Println("a after change:", a)

	num := 5
	fmt.Println("Before:", num)
	update(&num)
	fmt.Println("After:", num)

}
