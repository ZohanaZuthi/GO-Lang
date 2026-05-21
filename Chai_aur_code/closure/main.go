package main

import "fmt"

const a = 10

var p = 100

func Outer() func() {
	money := 100
	age := 30

	fmt.Println("Age=", age)
	show := func() {
		money = money + p + a
		fmt.Println(money)
	}

	return show
}

func call() {
	incr1 := Outer()
	incr1()
	incr1()

	incr2 := Outer()
	incr2()
	incr2()
}

/*
 there are two phases
1. compilation phase (compile time)
2. execution time

*********compile Phase*************

*** code segment ***

a=10
Outer=func(){....}
outerAnonymous function= func 1
call= func()

./main


********run time******
**** data segment ****
p=100






*/
