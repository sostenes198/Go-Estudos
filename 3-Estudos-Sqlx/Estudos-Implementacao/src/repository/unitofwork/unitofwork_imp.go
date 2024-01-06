package repository_unitofwork

import (
	pkgsql "3-Estudos-Sqlx/Estudos-Implementacao/pkg/sql"
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
)

type unitOfWorkImp struct {
	sqlService pkgsql.SqlService
}

func NewUnitOfWork(sqlService pkgsql.SqlService) UnitOfWork {
	return unitOfWorkImp{sqlService: sqlService}
}

func (u unitOfWorkImp) Execute(exec func(tx *sqlx.Tx, ctx *context.Context) (interface{}, error), ctx *context.Context) (interface{}, error) {
	db, err := u.sqlService.GetDb()
	if err != nil {
		return nil, err
	}
	tx := db.MustBeginTx(*ctx, nil)
	result, err := exec(tx, ctx)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return nil, errors.New(err.Error() + "\n" + rollbackErr.Error())
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return result, err
}
