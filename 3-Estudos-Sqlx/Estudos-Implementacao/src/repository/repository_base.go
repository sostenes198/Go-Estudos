package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
)

func ExecQuery[TParam interface{}](query string, param TParam, db *sqlx.DB, ctx *context.Context, tx *sqlx.Tx) error {
	if tx != nil {
		_, err := tx.NamedExecContext(*ctx, query, param)
		return err
	}

	if _, err := db.NamedExecContext(*ctx, query, param); err != nil {
		return err
	}

	return nil
}

func GetExecQuery[TParam map[string]interface{}, TModelDest interface{}](query string, param TParam, dest *TModelDest, db *sqlx.DB, ctx *context.Context, tx *sqlx.Tx) error {
	namedContext, err := _prepareStmt(query, db, ctx, tx)
	if err != nil {
		return err
	}

	err = namedContext.GetContext(*ctx, dest, param)
	if err != nil {
		return err
	}

	return nil
}

func ListExecQuery[TParam map[string]interface{}, TModelDest interface{}](query string, param TParam, dest *[]TModelDest, db *sqlx.DB, ctx *context.Context, tx *sqlx.Tx) error {
	namedContext, err := _prepareStmt(query, db, ctx, tx)
	if err != nil {
		return err
	}

	err = namedContext.SelectContext(*ctx, dest, param)
	if err != nil {
		return err
	}

	return nil
}

func _prepareStmt(query string, db *sqlx.DB, ctx *context.Context, tx *sqlx.Tx) (*sqlx.NamedStmt, error) {
	var namedContext *sqlx.NamedStmt
	var err error

	if tx != nil {
		namedContext, err = tx.PrepareNamedContext(*ctx, query)
		if err != nil {
			return nil, err
		}
	} else {
		namedContext, err = db.PrepareNamedContext(*ctx, query)
		if err != nil {
			return nil, err
		}
	}

	return namedContext, nil
}
