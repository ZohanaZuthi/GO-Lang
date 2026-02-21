package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in Go")
	Zuthi := User{"Zuthi", "nazifa.aip@gmail.com", 22, true}
	fmt.Printf("User details: %+v\n", Zuthi)
	fmt.Println(Zuthi)
	fmt.Printf("User Name: %v\n", Zuthi.Name)
	Zuthi.GetStatus()
	Zuthi.NewEmail()

	fmt.Println("The email after calling NewEmail method is", Zuthi.Email)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

// methods is function attached to a type mostly struct
func (u User) GetStatus() {
	fmt.Printf("The status of the user is %v\n", u.Status)
}

func (u *User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("The new email is", u.Email)
}

func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
	fmt.Println("The email has been updated to", u.Email)
}
