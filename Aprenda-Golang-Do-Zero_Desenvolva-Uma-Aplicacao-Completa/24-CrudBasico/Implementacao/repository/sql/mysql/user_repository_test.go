package mysql_test

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"implementacao/entitity"
	"implementacao/repository/sql/mysql"
	"log"
	"testing"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCreate(t *testing.T) {
	t.Parallel()

	// arrange
	const lastInsertIdExpected = 10
	defaultUser, _ := entitity.NewUser(entitity.Params{Name: "Test", Email: "Test@Testing.com"})

	t.Run("Create User", func(t *testing.T) {
		t.Parallel()

		// arrange
		db, mock := NewMock()
		defer db.Close()
		repo := mysql.NewUserRepository(db)

		prep := mock.ExpectPrepare("INSERT INTO usuario")
		prep.ExpectExec().WithArgs(defaultUser.Name, defaultUser.Email).WillReturnResult(sqlmock.NewResult(int64(lastInsertIdExpected), 1))

		// act
		lastInsertId, err := repo.Create(defaultUser)

		// assert
		require.NoError(t, err)
		require.Equal(t, lastInsertIdExpected, lastInsertId)
	})

	t.Run("Error when create defaultUser and fail to prepare query", func(t *testing.T) {
		t.Parallel()

		// arrange
		expectedError := errors.New("fail prepare query")
		db, mock := NewMock()
		defer db.Close()
		repo := mysql.NewUserRepository(db)

		prep := mock.ExpectPrepare("INSERT INTO usuario")
		prep.WillReturnError(expectedError)

		// act
		lastInsertId, err := repo.Create(defaultUser)

		// assert
		require.EqualError(t, expectedError, err.Error())
		require.Equal(t, 0, lastInsertId)
	})

	t.Run("Erro when create defaultUser and fail to exec query", func(t *testing.T) {
		t.Parallel()

		// arrange
		expectedError := errors.New("fail exec query")

		db, mock := NewMock()
		defer db.Close()
		repo := mysql.NewUserRepository(db)

		// act
		prep := mock.ExpectPrepare("INSERT INTO usuario")
		prep.ExpectExec().WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(expectedError)

		// act
		lastInsertId, err := repo.Create(defaultUser)

		// assert
		require.EqualError(t, expectedError, err.Error())
		require.Equal(t, 0, lastInsertId)
	})

	t.Run("Erro when create defaultUser and fail to get last insert id", func(t *testing.T) {
		t.Parallel()

		// arrange
		expectedError := errors.New("fail to get last insert id")

		db, mock := NewMock()
		defer db.Close()
		repo := mysql.NewUserRepository(db)

		// act
		prep := mock.ExpectPrepare("INSERT INTO usuario")
		prep.ExpectExec().WithArgs(defaultUser.Name, defaultUser.Email).WillReturnResult(sqlmock.NewErrorResult(expectedError))

		// act
		lastInsertId, err := repo.Create(defaultUser)

		// assert
		require.EqualError(t, expectedError, err.Error())
		require.Equal(t, 0, lastInsertId)
	})
}

func TestList(t *testing.T) {
	t.Parallel()

	// arrange
	expectedSql := "SELECT (.+) FROM usuario"
	columns := []string{
		"id", "nome", "email",
	}

	t.Run("List Users", func(t *testing.T) {
		t.Parallel()

		// arrange
		data := map[int]map[string]string{
			0: {
				"name":  "Test",
				"email": "Test@Test.com",
			},
			1: {
				"name":  "Test1",
				"email": "Test1@Test.com",
			},
			2: {
				"name":  "Test2",
				"email": "Test2@Test.com",
			},
		}
		expectedUsers := []entitity.User{
			{Id: 0, Name: data[0]["name"], Email: data[0]["email"]},
			{Id: 1, Name: data[1]["name"], Email: data[1]["email"]},
			{Id: 2, Name: data[2]["name"], Email: data[2]["email"]},
		}

		rows := sqlmock.NewRows(columns)
		rows.AddRow(0, data[0]["name"], data[0]["email"])
		rows.AddRow(1, data[1]["name"], data[1]["email"])
		rows.AddRow(2, data[2]["name"], data[2]["email"])

		db, mock := NewMock()
		defer db.Close()
		repository := mysql.NewUserRepository(db)

		mock.ExpectQuery(expectedSql).WillReturnRows(rows)

		// act
		users, err := repository.List()

		// assert
		require.NoError(t, err)
		require.Equal(t, expectedUsers, users)
	})

	t.Run("Erro when failed to exec query", func(t *testing.T) {
		t.Parallel()

		// arrange

		expectedError := errors.New("fail to exec query")

		db, mock := NewMock()
		defer db.Close()
		repository := mysql.NewUserRepository(db)

		mock.ExpectQuery(expectedSql).WillReturnError(expectedError)

		// act
		users, err := repository.List()

		// assert
		require.Nil(t, users)
		require.EqualError(t, expectedError, err.Error())
	})

	t.Run("Erro when failed to scan row", func(t *testing.T) {
		t.Parallel()

		// arrange

		expectedError := errors.New("sql: expected 1 destination arguments in Scan, not 3")

		db, mock := NewMock()
		defer db.Close()
		repository := mysql.NewUserRepository(db)

		rows := sqlmock.NewRows([]string{"a"})
		rows.AddRow("N")

		mock.ExpectQuery(expectedSql).WillReturnRows(rows)

		// act
		users, err := repository.List()

		// assert
		require.Nil(t, users)
		require.EqualError(t, expectedError, err.Error())
	})
}

func TestGetById(t *testing.T) {
	t.Parallel()

	// arrange
	id := 44
	expectedSql := "SELECT (.+) FROM usuario WHERE id=?"
	columns := []string{
		"id", "nome", "email",
	}

	t.Run("Get user by id", func(t *testing.T) {
		t.Parallel()

		// arrange
		name := "Test"
		email := "Testing@Testing.com"
		expectedUser := &entitity.User{Id: id, Name: name, Email: email}

		db, mock := NewMock()
		defer db.Close()
		repository := mysql.NewUserRepository(db)

		mock.ExpectQuery(expectedSql).WithArgs(id).WillReturnRows(sqlmock.NewRows(columns).AddRow(id, name, email))

		// act
		user, err := repository.Get(id)

		// assert
		require.Nil(t, err)
		require.Equal(t, expectedUser, user)
	})

	t.Run("Error when failed to exec query", func(t *testing.T) {
		t.Parallel()

		// arrange
		expectedError := errors.New("fail to exec query")

		db, mock := NewMock()
		defer db.Close()
		repository := mysql.NewUserRepository(db)

		mock.ExpectQuery(expectedSql).WithArgs(sqlmock.AnyArg()).WillReturnError(expectedError)

		// act
		user, err := repository.Get(id)

		// assert
		require.Nil(t, user)
		require.EqualError(t, expectedError, err.Error())
	})

	t.Run("Erro when failed to scan row", func(t *testing.T) {
		t.Parallel()

		// arrange
		expectedError := errors.New("sql: expected 1 destination arguments in Scan, not 3")

		db, mock := NewMock()
		defer db.Close()
		repository := mysql.NewUserRepository(db)

		rows := sqlmock.NewRows([]string{"a"})
		rows.AddRow("N")

		mock.ExpectQuery(expectedSql).WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)

		// act
		user, err := repository.Get(id)

		// assert
		require.Nil(t, user)
		require.EqualError(t, expectedError, err.Error())
	})
}
