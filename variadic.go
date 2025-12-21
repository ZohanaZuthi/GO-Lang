package main

import "fmt"

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(Sum(1, 2, 3))        // 3 numbers
	fmt.Println(Sum(10, 20, 30, 40)) // 4 numbers
	fmt.Println(Sum())               // 0 numbers — also fine!

	nums := []int{5, 10, 15}
	fmt.Println(Sum(nums...)) // passing a slice

}
