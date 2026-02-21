package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in Go")
	// no inheritence in golang: no super r parent
	Zuthi := User{"Zuthi", "nazifa.aip@gmail.com", 22, true}
	fmt.Printf("User details: %+v\n", Zuthi)
	fmt.Println(Zuthi)
	fmt.Printf("User Name: %v\n", Zuthi.Name)
}

// as the class and the field will be public they are capitalized
type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}
