# Go Function Concepts — Interview Preparation Notes

## 1. What is a Function in Go?

A function is a reusable block of code that performs a specific task.

Basic syntax:

```go
func functionName(parameters) returnType {
    // code
}
```

Example:

```go
func add(a int, b int) int {
    return a + b
}
```

---

# 2. Parameter vs Argument

## Parameter

Variables defined in the function definition.

```go
func add(a int, b int)
```

`a` and `b` are parameters.

---

## Argument

Actual values passed to the function.

```go
add(3, 5)
```

`3` and `5` are arguments.

---

# 3. First-Order Function

A first-order function works only with normal values.

Example:

```go
func add(x int, y int) {
    fmt.Println(x + y)
}
```

It does NOT:

- take functions as parameters
- return functions

---

# 4. First-Class Functions

In Go, functions are treated as **first-class citizens**.

This means:

- functions can be assigned to variables
- functions can be passed as arguments
- functions can be returned from functions

Example:

```go
sum := add
sum(2, 3)
```

---

# 5. Higher-Order Function

A higher-order function either:

1. Takes another function as parameter
2. Returns a function
3. Or both

---

## Example: Function as Parameter

```go
func process(a int, b int, op func(int, int)) {
    op(a, b)
}
```

Here:

- `process()` is a higher-order function
- because it accepts another function

---

## Example: Function Returning Function

```go
func getFunction() func(int, int) {
    return add
}
```

This is also a higher-order function.

---

# 6. Anonymous Function

A function without a name.

Example:

```go
func(a int, b int) {
    fmt.Println(a + b)
}
```

Useful for:

- short temporary logic
- callbacks
- closures

---

# 7. IIFE (Immediately Invoked Function Expression)

An anonymous function executed immediately.

Example:

```go
func(a int, b int) {
    fmt.Println(a + b)
}(3, 5)
```

Explanation:

- function is defined
- immediately called using `(3,5)`

---

# 8. Callback Function

A function passed into another function to be executed later.

Example:

```go
func process(a int, b int, op func(int, int)) {
    op(a, b)
}

process(3, 4, add)
```

Here:

- `add` is a callback function

---

# 9. Functional Programming Concept

Functional programming focuses on:

- functions as values
- immutability
- composition
- declarative style

Languages:

- Haskell
- Racket

Go supports some functional concepts but is not purely functional.

---

# 10. Logic Concepts (Interview Theory)

## Object

Entities like:

- people
- animal
- car

Example:

```text
Tutul
Apple
Rakib
```

---

## Property

Characteristics of objects.

Example:

```text
Apple is red
Tutul is a student
```

---

## Relation

Connection between objects.

Example:

```text
Tutul is taller than Rakib
```

---

# 11. First-Order Logic

Rules about objects and properties.

Example:

```text
All students must wear uniforms.
All customers must pay bills.
```

---

# 12. Higher-Order Logic

Rules applied to other rules.

Example:

```text
Any rule applied to all customers
must also apply to Tutul.
```

---

# 13. Full Interview Example

```go
package main

import "fmt"

// First-order function
func add(x int, y int) {
	result := x + y
	fmt.Println("Sum:", result)
}

// Higher-order function
func processOperation(a int, b int, operation func(int, int)) {
	operation(a, b)
}

// Function returning another function
func getAddFunction() func(int, int) {
	return add
}

func main() {

	// Normal function call
	add(3, 5)

	// Passing function as argument
	processOperation(10, 20, add)

	// Returning function from function
	sumFunction := getAddFunction()
	sumFunction(7, 8)

	// Anonymous function + IIFE
	func(a int, b int) {
		fmt.Println("Anonymous Sum:", a+b)
	}(4, 6)
}
```

---

# 14. Important Interview Questions

## Q1. What is a higher-order function?

A function that:

- takes another function as parameter
- returns a function
- or both

---

## Q2. Are Go functions first-class citizens?

Yes.

Functions in Go can:

- be stored in variables
- passed as arguments
- returned from functions

---

## Q3. Difference between anonymous function and named function?

| Named Function   | Anonymous Function |
| ---------------- | ------------------ |
| Has name         | No name            |
| Reusable         | Usually temporary  |
| Defined globally | Often inline       |

---

## Q4. What is IIFE?

Immediately Invoked Function Expression.

A function created and executed instantly.

---

## Q5. Why use higher-order functions?

- cleaner code
- reusable logic
- callbacks
- functional programming patterns
