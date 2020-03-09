package main

import "fmt"

type stringRepresentable interface {
	raw() string
}

type person struct {
	first string
	last  string
	age   int8
}

func (p person) raw() (str string) {
	str = fmt.Sprintf("name: %v %v, age: %v", p.first, p.last, p.age)
	return str
}

func main() {

	p1 := person{
		first: "Daniel",
		last:  "Sanchez R",
		age:   24,
	}

	var rr stringRepresentable
	rr = p1

	fmt.Println(rr.raw())

}
