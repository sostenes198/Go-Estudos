package repository_unitofwork

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type UnitOfWork interface {
	Execute(exec func(tx *sqlx.Tx, ctx *context.Context) (interface{}, error), ctx *context.Context) (interface{}, error)
}
