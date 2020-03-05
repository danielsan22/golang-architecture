package architecture

// Person is how the architecture package stores a person
type Person struct {
	First string
}

// Accessor is how to store or retrieve a person
// When retriving a person if they not exist, return the zero value
type Accessor interface {
	Save(n int, p Person)
	Retrieve(n int) Person
}

// PersonService is the basic ser
type PersonService struct {
	a Accessor
}

func NewPersonService(a Accessor) PersonService {
	return PersonService{
		a: a,
	}
}

func Put(a Accessor, n int, p Person) {
	a.Save(n, p)
}

func Get(a Accessor, n int) Person {
	return a.Retrieve(n)
}
