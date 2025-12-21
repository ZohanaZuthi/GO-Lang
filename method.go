package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) birthday() {
	p.age++ // modifies actual struct because p is a pointer
	fmt.Println(p.name, "just had a birthday!")
}

type Order struct {
	id     int
	amount float32
	status string
}

func (o Order) str() string {
	if o.status == "paid" {
		return "Order is paid"
	}
	return "Order is not paid"
}

type customer struct {
	id   int
	name string
}

// embeded struct
type order2 struct {
	order_id int
	amount   float32
	customer customer
}

// Method attached to Person
func (p Person) greet() {
	fmt.Println("Hello, my name is", p.name)
}

func main() {
	p := Person{name: "Zuthi", age: 22}
	p.greet() // method call
	p.birthday()
	fmt.Println(p.age)

	q := Order{1, 50.00, "paid"}
	fmt.Println(q)
	fmt.Println(q.str())

	newOrder := order2{
		order_id: 101,
		amount:   250.50,
		customer: customer{
			id:   4,
			name: "Alice",
		},
	}

	language := struct {
		name   string
		isGood bool
	}{"Golang", true}
	fmt.Println(language)

	fmt.Println(newOrder)

}
