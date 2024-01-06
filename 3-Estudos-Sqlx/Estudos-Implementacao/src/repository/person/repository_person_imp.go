package repository_person

import (
	pkgSql "3-Estudos-Sqlx/Estudos-Implementacao/pkg/sql"
	domainperson "3-Estudos-Sqlx/Estudos-Implementacao/src/domain/person"
	"3-Estudos-Sqlx/Estudos-Implementacao/src/repository"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepositoryPersonImp struct {
	sqlService pkgSql.SqlService
	parser     repository.ParserModel[domainperson.Person, PersonModel]
}

func NewRepositoryPerson(sqlService pkgSql.SqlService, parser repository.ParserModel[domainperson.Person, PersonModel]) RepositoryPerson {
	return RepositoryPersonImp{sqlService: sqlService, parser: parser}
}

func (repo RepositoryPersonImp) Create(person domainperson.Person, ctx *context.Context, tx *sqlx.Tx) error {
	db, err := repo.sqlService.GetDb()
	if err != nil {
		return err
	}

	model := repo.parser.ParseToModel(person)

	if _, err = repository.ExecQuery(createPersonQuery, model, db, ctx, tx); err != nil {
		return err
	}

	return nil
}

func (repo RepositoryPersonImp) Update(person domainperson.Person, ctx *context.Context, tx *sqlx.Tx) error {
	var err error = nil

	db, err := repo.sqlService.GetDb()
	if err != nil {
		return err
	}

	model := repo.parser.ParseToModel(person)

	if _, err = repository.ExecQuery(updatePersonQuery, model, db, ctx, tx); err != nil {
		return err
	}

	return nil
}

func (repo RepositoryPersonImp) DeleteByEmail(email string, ctx *context.Context, tx *sqlx.Tx) error {
	var err error = nil

	db, err := repo.sqlService.GetDb()
	if err != nil {
		return err
	}

	param := map[string]interface{}{
		"email": email,
	}

	if _, err = repository.ExecQuery(deletePersonByEmailQuery, param, db, ctx, tx); err != nil {
		return err
	}

	return nil
}

func (repo RepositoryPersonImp) GetFirstOrDefaultByLastName(lastname string, ctx *context.Context, tx *sqlx.Tx) (domainperson.Person, error) {
	var err error = nil

	db, err := repo.sqlService.GetDb()
	if err != nil {
		return domainperson.EmptyPerson(), err
	}

	model := PersonModel{}
	param := map[string]interface{}{
		"last_name": lastname,
	}

	if err = repository.GetExecQuery(getByLastNameQuery, param, &model, db, ctx, tx); err != nil {
		return domainperson.EmptyPerson(), err
	}

	return repo.parser.ParseToEntity(model), nil
}

func (repo RepositoryPersonImp) ListByLastName(lastname string, ctx *context.Context, tx *sqlx.Tx) ([]domainperson.Person, error) {
	var err error = nil
	emptyPersons := make([]domainperson.Person, 0)

	db, err := repo.sqlService.GetDb()
	if err != nil {
		return emptyPersons, err
	}

	var models []PersonModel
	param := map[string]interface{}{
		"last_name": lastname,
	}

	err = repository.ListExecQuery(getByLastNameQuery, param, &models, db, ctx, tx)
	if err != nil {
		return nil, err
	}

	return _parseModelsTOEntities(repo.parser, models), nil
}

func _parseModelsTOEntities(parser repository.ParserModel[domainperson.Person, PersonModel], models []PersonModel) []domainperson.Person {
	var persons []domainperson.Person
	for _, model := range models {
		persons = append(persons, parser.ParseToEntity(model))
	}
	return persons
}
