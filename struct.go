package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	// Create struct values in different ways
	var p1 Person
	p1.name = "Zuthi"
	p1.age = 20

	// Using shorthand
	p2 := Person{name: "Rahul", age: 25}

	// Positional initialization (order matters)
	p3 := Person{"Aditi", 30}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	fmt.Println("Name:", p1.name)
	fmt.Println("Age:", p1.age)
}
