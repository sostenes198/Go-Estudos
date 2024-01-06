package repository_person

type PersonModel struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

func NewPersonModel(firstName string, lastName string, email string) PersonModel {
	return PersonModel{FirstName: firstName, LastName: lastName, Email: email}
}
