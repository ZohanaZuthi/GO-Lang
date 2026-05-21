package main

import "fmt"

// 1. parameter vs argument
// 2. First Order Function
// i. standard function or named function
// ii. anonymous function
// iii. IIFE
// iv. Function Expression
// 3. higher Order Function
// 4. callback function
//  fucntional function-> functional paradigm-> haskel
//  and racket

// discrete
// 1 . first order
// 2 . higher order

//  Logic
// 1. Object (people,animal, car etc)
// 2. Property (color, student)
// 3. relation

// Tutul is a student
// Apple is red
// Tutul is taller than Rakib

// Statement
//  Rule :(first order logic)
//  All customers must pay their pizza bills
//  All students must wear their uniforms

// Rule:(higher order logic)
// Any rules that applies to all customer must aldo apply to Tutul

//  higher order any rules among 3
//  1. parameter-> function
// 2. function return
// 3. both
func processOperation(a int, b int, op func(x int, y int)) {
	op(a, b)
}

// function return
//  the function is called or return in the higher order function
// higher order function is called first class function
// as it can be treated as first class citizen
func call() func(x int, y int) {
	return add
}

// 1st order function
func add(x int, y int) {
	z := x + y
	fmt.Println(z)

}

func main() {
	//  first class citizen any kind of variable can be assigned
	a := 3
	fmt.Println("Hello")
	processOperation(a, 5, add)

	sum := call()

	sum(2, 7)
}
