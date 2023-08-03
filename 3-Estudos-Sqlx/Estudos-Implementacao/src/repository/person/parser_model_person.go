package repository_person

import (
	domainperson "3-Estudos-Sqlx/Estudos-Implementacao/src/domain/person"
	"3-Estudos-Sqlx/Estudos-Implementacao/src/repository"
)

type parseModelPerson struct {
}

func NewParserModelPerson() repository.ParserModel[domainperson.Person, personModel] {
	return parseModelPerson{}
}

func (p parseModelPerson) ParseToModel(entity domainperson.Person) personModel {
	return newPersonModel(entity.FirstName, entity.LastName, entity.Email)
}

func (p parseModelPerson) ParseToEntity(model personModel) domainperson.Person {
	return domainperson.NewPerson(model.FirstName, model.LastName, model.Email)
}
