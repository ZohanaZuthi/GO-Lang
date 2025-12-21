package main

import "fmt"

func main() {
	// declaring and initializing an array of integers
	var arr1 [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array 1:", arr1) // [10 20 30 40 50]

	var name [3]string
	name[0] = "golang"
	fmt.Println("Array 2:", name) // [golang  ]

	names := [3]string{"Zuthi", "Arafat", "Sakib"}
	fmt.Println("Array 3:", names[2]) // [Zuthi Arafat Sakib]

	//multi dimensional array
	var multiArr [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("Multi Dimensional Array:", multiArr) // [[1 2 3] [4 5 6]]

	// fixed size that is predictable
	// memory optimization
	// constant time access

}
