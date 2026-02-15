package main

import "fmt"

func main() {

	// 	var array_name = [length]datatype{values} // here length is defined
	// var array_name = [...]datatype{values} // here length is inferred
	// array_name := [length]datatype{values} // here length is defined
	// array_name := [...]datatype{values} // here length is inferred
	fmt.Println("Welcome to pointers in Go")
	var arr [7]int
	arr[0] = 23
	arr[1] = 45
	arr[2] = 67
	fmt.Println("Array:", arr)
	fmt.Println("Length of Array:", len(arr))

	var veg = [3]string{"Potato", "Tomato", "Cabbage"}
	fmt.Println("Vegetables:", veg)
	fmt.Println("Length of Vegetables Array:", len(veg))

	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
}
