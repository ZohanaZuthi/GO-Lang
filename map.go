package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30

	fmt.Println(m["a"])
	fmt.Println(m["b"])
	fmt.Println(m["w"])
	fmt.Println(len(m))

	delete(m, "b")
	fmt.Println(len(m))

	n := map[string]int{
		"x": 100,
		"y": 200,
	}
	fmt.Println(n)

	v, ok := m["c"]
	if ok {
		fmt.Println("Value:", v)
	} else {
		fmt.Println("Key not found")

	}

	fmt.Println(maps.Equal(m, n))
	for index, value := range m {
		fmt.Println("Key:", index, "Value:", value)

	}

	for index := range m {
		fmt.Println("Key:", index)
	}
}
