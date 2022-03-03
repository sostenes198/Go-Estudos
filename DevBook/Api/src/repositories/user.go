package repositories

import (
	"database/sql"
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories/contracts"
	"devbook/src/repositories/queries"
	"fmt"
)

type userRepository struct {
	db *sql.DB
}

// NewUserRepository cria uma instância do repositório de usuário
func NewUserRepository() contracts.UserRepository {
	return &userRepository{database.Db}
}

func (repository *userRepository) List(filter contracts.UserListFilter) ([]models.User, error) {
	var users []models.User
	nameOrNick := fmt.Sprintf("%%%s%%", filter.NameOrNick)

	rows, err := repository.db.Query(queries.UserList, nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if rows.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.Password, &user.CreateAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository *userRepository) GetById(id uint64) (*models.User, error) {
	var user models.User

	row, err := repository.db.Query(queries.UserGetById, id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if row.Next() {
		if err := row.Scan(&user.Id, &user.Name, &user.Nick, &user.Email, &user.Password, &user.CreateAt); err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (repository *userRepository) Create(user *models.User) error {
	statement, err := repository.db.Prepare(queries.UserCreate)
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.Id = uint64(lastInsertId)
	return nil
}
