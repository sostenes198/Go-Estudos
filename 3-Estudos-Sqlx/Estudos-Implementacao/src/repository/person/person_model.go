package repository_person

type personModel struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

func newPersonModel(firstName string, lastName string, email string) personModel {
	return personModel{FirstName: firstName, LastName: lastName, Email: email}
}
