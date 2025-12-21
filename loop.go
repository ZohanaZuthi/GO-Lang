package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
		if j == 3 {
			break
		}
	}

	// for 1.22 version

	for i := range 11 {
		fmt.Println(i)
	}

	for i := range []int{10, 20, 30, 40, 50} {
		fmt.Println(i)
	}

}
