package main

// slices are dynamic arrays in Go, they can grow and shrink in size as needed. They are built on top of arrays and provide a more flexible way to work with collections of data. Slices have a length and a capacity, where the length is the number of elements in the slice and the capacity is the total number of elements that can be stored in the underlying array before it needs to be resized.
import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in Go")

	var fruitList = []string{"Apple", "Banana", "Grapes"}
	fmt.Println("Fruit List:", fruitList)
	fmt.Println("Length of Fruit List:", len(fruitList))
	fmt.Println("Capacity of Fruit List:", cap(fruitList))
	fruitList = append(fruitList, "Mango", "Orange")
	fmt.Println("Updated Fruit List:", fruitList)
	fmt.Println("New Length of Fruit List:", len(fruitList))
	fmt.Println("New Capacity of Fruit List:", cap(fruitList))
	slicedFruits := fruitList[1:4]
	fmt.Println("Sliced Fruits (index 1 to 3):", slicedFruits)
	fmt.Println("Length of Sliced Fruits:", len(slicedFruits))
	fmt.Println("Capacity of Sliced Fruits:", cap(slicedFruits))

	fruitList = append(fruitList, "Lichi", "Papaya")
	fmt.Println("Updated Fruit List:", fruitList)
	fmt.Println("New Length of Fruit List:", len(fruitList))
	fmt.Println("New Capacity of Fruit List:", cap(fruitList))
	//Go grows slices for speed, not by a fixed formula

	// The make() function can also be used to create a slice.
	highscore := make([]int, 4)
	highscore[0] = 100
	highscore[1] = 200
	highscore[2] = 300
	highscore[3] = 400
	// highscore[4] = 500 // this will cause an error because the length of the slice is 4, and we are trying to access index 4 which is out of range
	// if we add one more element it will crash
	highscore = append(highscore, 300, 220) // this will work because append will automatically resize the slice
	fmt.Println("High Scores:", highscore)

	sort.Ints(highscore)
	fmt.Println("Sorted High Scores:", highscore)

	highscore = append(highscore[:2], highscore[3:]...) // this will remove the element at index 2 (300) from the slice
	fmt.Println("High Scores after deletion:", highscore)

	// remove a value from the slice based on index

}
