package main

import (
	"fmt"

	architecture "github.com/danielsan22/golang-architecture"
	mongo "github.com/danielsan22/golang-architecture/storage/mongo"
	postg "github.com/danielsan22/golang-architecture/storage/postgres"
)

func main() {
	dbm := mongo.Mongo{}
	dbp := postg.Postg{}

	p1 := architecture.Person{
		First: "Jenny",
	}

	p2 := architecture.Person{
		First: "James",
	}

	architecture.Put(dbm, 1, p1)
	architecture.Put(dbm, 2, p2)
	fmt.Println(architecture.Get(dbm, 1))
	fmt.Println(architecture.Get(dbm, 2))

	// or store in some other data storage
	architecture.Put(dbp, 1, p1)
	architecture.Put(dbp, 2, p2)
	fmt.Println(architecture.Get(dbp, 1))
	fmt.Println(architecture.Get(dbp, 2))
}
