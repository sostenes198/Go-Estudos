package mysql

import (
	"database/sql"
	"implementacao/entitity"
	"implementacao/repository/mysql/base"
	"implementacao/repository/mysql/models"
	"implementacao/repository/mysql/query"
	"log"
)

type userRepository struct {
	db *sql.DB
}

func (userRepository *userRepository) Create(user *entitity.User) (int, error) {
	repository, err := base.Open()
	if err != nil {
		return 0, err
	}
	defer repository.Close()

	insertUser := models.NewUserInsert(*user)

	statement, err := repository.Prepare(query.CreateUser)
	if err != nil {
		return 0, err
	}
	defer func() {
		err = statement.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	insertionResult, err := statement.Exec(insertUser.Name, insertUser.Email)
	if err != nil {
		return 0, err
	}

	lastIdInsert, err := insertionResult.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastIdInsert), nil
}
