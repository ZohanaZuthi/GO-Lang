package main

import "fmt"

// defer is Last in firt out queue that works very last of the code

func main() {
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	defer fmt.Println("World")
	fmt.Println("Hello")

	math()

}

func math() {
	for i := 0; i < 5; i++ {
		fmt.Print(i)

	}
}
