package repository_person

import (
	pkgSql "3-Estudos-Sqlx/Estudos-Implementacao/pkg/sql"
	domainperson "3-Estudos-Sqlx/Estudos-Implementacao/src/domain/person"
	"3-Estudos-Sqlx/Estudos-Implementacao/src/repository"
	"context"
	"github.com/jmoiron/sqlx"
)

var _parser = NewParserModelPerson()

type repositoryPersonImp struct {
	sqlService *pkgSql.SqlService
}

func NewRepositoryPerson(sqlService *pkgSql.SqlService) RepositoryPerson {
	return repositoryPersonImp{sqlService: sqlService}
}

func (repo repositoryPersonImp) Create(person domainperson.Person, ctx *context.Context, tx *sqlx.Tx) error {
	db, err := repo.sqlService.GetDb()
	if err != nil {
		return err
	}

	model := _parser.ParseToModel(person)

	if err = repository.ExecQuery(createPersonQuery, model, db, ctx, tx); err != nil {
		return err
	}

	return nil
}

func (repo repositoryPersonImp) Update(person domainperson.Person, ctx *context.Context, tx *sqlx.Tx) error {
	var err error = nil

	db, err := repo.sqlService.GetDb()
	if err != nil {
		return err
	}

	model := _parser.ParseToModel(person)

	if err = repository.ExecQuery(updatePersonQuery, model, db, ctx, tx); err != nil {
		return err
	}

	return nil
}

func (repo repositoryPersonImp) DeleteByEmail(email string, ctx *context.Context, tx *sqlx.Tx) error {
	var err error = nil

	db, err := repo.sqlService.GetDb()
	if err != nil {
		return err
	}

	param := map[string]interface{}{
		"email": email,
	}

	if err = repository.ExecQuery(deletePersonByEmailQuery, param, db, ctx, tx); err != nil {
		return err
	}

	return nil
}

func (repo repositoryPersonImp) GetFirstOrDefaultByLastName(lastname string, ctx *context.Context, tx *sqlx.Tx) (domainperson.Person, error) {
	var err error = nil

	db, err := repo.sqlService.GetDb()
	if err != nil {
		return domainperson.EmptyPerson(), err
	}

	model := personModel{}
	param := map[string]interface{}{
		"last_name": lastname,
	}

	if err = repository.GetExecQuery(getByLastNameQuery, param, &model, db, ctx, tx); err != nil {
		return domainperson.EmptyPerson(), err
	}

	return _parser.ParseToEntity(model), nil
}

func (repo repositoryPersonImp) ListByLastName(lastname string, ctx *context.Context, tx *sqlx.Tx) ([]domainperson.Person, error) {
	var err error = nil
	emptyPersons := make([]domainperson.Person, 0)

	db, err := repo.sqlService.GetDb()
	if err != nil {
		return emptyPersons, err
	}

	var models []personModel
	param := map[string]interface{}{
		"last_name": lastname,
	}

	err = repository.ListExecQuery(getByLastNameQuery, param, &models, db, ctx, tx)
	if err != nil {
		return nil, err
	}

	return _parseModelsTOEntities(models), nil
}

func _parseModelsTOEntities(models []personModel) []domainperson.Person {
	var persons []domainperson.Person
	for _, model := range models {
		persons = append(persons, _parser.ParseToEntity(model))
	}
	return persons
}
