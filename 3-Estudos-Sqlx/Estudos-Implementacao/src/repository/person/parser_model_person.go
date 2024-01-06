package repository_person

import (
	domainperson "3-Estudos-Sqlx/Estudos-Implementacao/src/domain/person"
	"3-Estudos-Sqlx/Estudos-Implementacao/src/repository"
)

type parseModelPerson struct {
}

func NewParserModelPerson() repository.ParserModel[domainperson.Person, PersonModel] {
	return parseModelPerson{}
}

func (p parseModelPerson) ParseToModel(entity domainperson.Person) PersonModel {
	return NewPersonModel(entity.FirstName, entity.LastName, entity.Email)
}

func (p parseModelPerson) ParseToEntity(model PersonModel) domainperson.Person {
	return domainperson.NewPerson(model.FirstName, model.LastName, model.Email)
}
