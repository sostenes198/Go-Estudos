package domain_person

type Person struct {
	FirstName string
	LastName  string
	Email     string
}

func NewPerson(firstName string, lastName string, email string) Person {
	return Person{FirstName: firstName, LastName: lastName, Email: email}
}

func EmptyPerson() Person {
	return Person{}
}
