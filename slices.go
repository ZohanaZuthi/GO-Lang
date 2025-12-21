package main

import "fmt"

// slice is dynamic
// most used construct in go
// useful methods
func main() {
	// uninitialized slice is nil
	var nums []int
	fmt.Println("Slice:", nums)
	fmt.Println(nums == nil)
	fmt.Println(cap(nums))

	var nums1 = make([]int, 5) // length 5, capacity 5
	fmt.Println("Slice:", nums1)
	fmt.Println(cap(nums1))
	nums1 = append(nums1, 10)
	fmt.Println("Slice:", nums1)

	var nums2 = make([]int, 3, 5) // length 3, capacity 5
	fmt.Println("Slice:", nums2)
	fmt.Println(cap(nums2))

	var nums3 = make([]int, len(nums2))
	fmt.Println("Slice:", nums3)
	fmt.Println(cap(nums3))

	copied := make([]int, len(nums2))
	copy(copied, nums2)
	fmt.Println("Copied slice:", copied)

	fmt.Println("Slice[1:3]:", nums2[1:3])

}
