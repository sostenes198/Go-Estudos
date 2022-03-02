package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySql
	"implementacao/entitity"
	"implementacao/repository/contracts"
	"implementacao/repository/sql/mysql/models"
	"implementacao/repository/sql/mysql/query"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) contracts.UserRepository {
	return &UserRepository{db: db}
}

func (repository *UserRepository) Create(user *entitity.User) (int, error) {
	insertUser := models.NewUserInsert(*user)

	statement, err := repository.db.Prepare(query.Create)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

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

func (repository *UserRepository) List() ([]entitity.User, error) {
	var users []entitity.User

	rows, err := repository.db.Query(query.List)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entitity.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository *UserRepository) Get(id int) (*entitity.User, error) {
	var user entitity.User

	row, err := repository.db.Query(query.GetById, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if row.Next() {
		if err := row.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (repository *UserRepository) Update(user *entitity.User) error {
	statement, err := repository.db.Prepare(query.Update)
	if err != nil {
		return nil
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, user.Id); err != nil {
		return err
	}

	return nil
}

func (repository *UserRepository) Delete(id int) error {
	statement, err := repository.db.Prepare(query.Delete)
	if err != nil {
		return nil
	}
	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		return err
	}

	return nil
}
