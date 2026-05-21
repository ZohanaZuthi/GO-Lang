package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This is for golang file creation"

	file, err := os.Create("./files/create_file.txt")

	checkErr(err)

	length, err := io.WriteString(file, content)
	checkErr(err)

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./files/create_file.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkErr(err)
	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}
