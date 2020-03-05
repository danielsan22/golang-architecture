package architecture

import (
	"fmt"
	"testing"
)

type Mongo map[int]Person

func (m Mongo) Save(n int, p Person) {
	m[n] = p
}

func (m Mongo) Retrieve(n int) Person {
	return m[n]
}

func TestPut(t *testing.T) {

	mdb := Mongo{}
	p := Person{
		First: "Daniel",
	}

	Put(mdb, 1, p)
	got := mdb.Retrieve(1)
	if got != p {
		t.Fatalf("Whant %v, got %v", p, got)
	}
}

func ExamplePut() {

	mdb := Mongo{}
	p := Person{
		First: "Daniel",
	}

	Put(mdb, 1, p)
	got := mdb.Retrieve(1)
	fmt.Println(got)

	// Output: {Daniel}
}
