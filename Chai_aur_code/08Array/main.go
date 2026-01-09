package main

import "fmt"

func main() {
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
}
