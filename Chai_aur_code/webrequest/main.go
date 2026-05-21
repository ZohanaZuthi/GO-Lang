package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)
	checkErr(err)

	fmt.Printf("Response is of type:%T\n", response)

	defer response.Body.Close()
	databytes, err := io.ReadAll(response.Body)
	checkErr(err)

	fmt.Printf(string(databytes))

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
