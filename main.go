package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p person) speak() {
	fmt.Printf("From a person - this is my name", p.first)
}

func main() {
	p1 := person{
		first: "James",
	}

	fmt.Printf("%T\n", p1)

	var x human
	x = p1

	fmt.Printf("%T\n", x)

}
