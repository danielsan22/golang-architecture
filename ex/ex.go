package main

import "fmt"

type person struct {
	name string
}

type mongodb map[int]person

type accessor interface {
	save(int, person)
	retrieve(int) person
}

func (m *mongodb) save(i int, p person) {
	v := *m
	v[i] = p
}

func (m *mongodb) retrieve(i int) person {
	v := *m
	return v[i]
}

func main() {
	p1 := person{
		name: "Daniel",
	}

	acc := mongodb{}

	put(&acc, 1, p1)

	lastPerson := get(&acc, 1)

	fmt.Println(lastPerson.name)

}

func put(acc accessor, id int, p person) {
	acc.save(id, p)
}

func get(acc accessor, id int) person {
	return acc.retrieve(id)
}
