package repository_person

import (
	domainperson "3-Estudos-Sqlx/Estudos-Implementacao/src/domain/person"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepositoryPerson interface {
	Create(person domainperson.Person, ctx *context.Context, tx *sqlx.Tx) error
	Update(person domainperson.Person, ctx *context.Context, tx *sqlx.Tx) error
	DeleteByEmail(email string, ctx *context.Context, tx *sqlx.Tx) error
	GetFirstOrDefaultByLastName(lastname string, ctx *context.Context, tx *sqlx.Tx) (domainperson.Person, error)
	ListByLastName(lastname string, ctx *context.Context, tx *sqlx.Tx) ([]domainperson.Person, error)
}
