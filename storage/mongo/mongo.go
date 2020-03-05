package mongo

import architecture "github.com/danielsan22/golang-architecture"

type Mongo map[int]architecture.Person

func (m Mongo) Save(n int, p architecture.Person) {
	m[n] = p
}

func (m Mongo) Retrieve(n int) architecture.Person {
	return m[n]
}
