package postg

import architecture "github.com/danielsan22/golang-architecture"

type Postg map[int]architecture.Person

func (pg Postg) Save(n int, p architecture.Person) {
	pg[n] = p
}

func (pg Postg) Retrieve(n int) architecture.Person {
	return pg[n]
}
