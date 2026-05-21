package main

import "fmt"

func main() {

	fmt.Println("Welcome to structs in Go")
	// no inheritence in golang: no super r parent
	Zuthi := User{"Zuthi", "nazifa.aip@gmail.com", 22, true, customer{"natasha", 22}}
	fmt.Printf("User details: %+v\n", Zuthi)
	fmt.Println(Zuthi)
	fmt.Printf("User Name: %v\n", Zuthi.Name)
}

type customer struct {
	name string
	age  int
}

// as the class and the field will be public they are capitalized
type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
	customer
}
