package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	fmt.Println("List of languages:", languages)
	fmt.Println("JS stands for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of languages after deletion:", languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
