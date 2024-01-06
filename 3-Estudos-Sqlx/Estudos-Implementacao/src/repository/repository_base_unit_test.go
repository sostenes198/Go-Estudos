//go:build unittest

package repository

import (
	testfixtures "3-Estudos-Sqlx/Estudos-Implementacao/test/fixtures"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	insert string = "INSERT INTO TEST"
	query  string = "SELECT id, name FROM TEST WHERE none:=none"
)

type repositoryModelTest struct {
	Id   int
	Name string
}

type invalidRepositoryModelTest struct {
	id   int
	name string
}

func TestRepositoryBase(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()

	param := map[string]interface{}{
		"none": "NONE",
	}

	t.Run("Exec Query", func(t *testing.T) {
		t.Parallel()

		t.Run("Should exec query", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, mock := testfixtures.CreateSqlXDbMock(t)

			mock.ExpectExec(insert).WillReturnResult(sqlmock.NewResult(1, 1))

			// act
			result, err := ExecQuery[map[string]interface{}](insert, nil, dbSqlX, &ctx, nil)

			// assert
			assert.Nil(t, err)
			rows, err := result.RowsAffected()
			assert.Nil(t, err)
			assert.Equal(t, int64(1), rows)
			lastId, err := result.LastInsertId()
			assert.Nil(t, err)
			assert.Equal(t, int64(1), lastId)
		})

		t.Run("Should return error when failed to exec query", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, _ := testfixtures.CreateSqlXDbMock(t)

			// act
			result, err := ExecQuery[map[string]interface{}](insert, nil, dbSqlX, &ctx, nil)

			// assert
			assert.EqualError(t, err, "all expectations were already fulfilled, call to ExecQuery 'INSERT INTO TEST' with args [] was not expected")
			assert.Nil(t, result)
		})

		t.Run("Should exec query with transaction", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, mock := testfixtures.CreateSqlXDbMock(t)

			mock.ExpectBegin()
			mock.ExpectExec(insert).WillReturnResult(sqlmock.NewResult(1, 1))

			tx, err := dbSqlX.BeginTxx(ctx, nil)

			// act
			result, err := ExecQuery[map[string]interface{}](insert, nil, dbSqlX, &ctx, tx)

			// assert
			assert.Nil(t, err)
			rows, err := result.RowsAffected()
			assert.Nil(t, err)
			assert.Equal(t, int64(1), rows)
			lastId, err := result.LastInsertId()
			assert.Nil(t, err)
			assert.Equal(t, int64(1), lastId)
		})

		t.Run("Should return error when failed to exec query with transaction", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, mock := testfixtures.CreateSqlXDbMock(t)

			mock.ExpectBegin()

			tx, err := dbSqlX.BeginTxx(ctx, nil)

			// act
			result, err := ExecQuery[map[string]interface{}](insert, nil, dbSqlX, &ctx, tx)

			// assert
			assert.EqualError(t, err, "all expectations were already fulfilled, call to ExecQuery 'INSERT INTO TEST' with args [] was not expected")
			assert.Nil(t, result)
		})
	})

	t.Run("Get Exec Query", func(t *testing.T) {
		t.Parallel()

		t.Run("Should Get Exec Query", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, mock := testfixtures.CreateSqlXDbMock(t)

			rows := sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "Test")

			mock.ExpectPrepare("^SELECT (.+) FROM TEST WHERE").ExpectQuery().WillReturnRows(rows)

			dest := repositoryModelTest{}

			expectedResult := repositoryModelTest{Id: 1, Name: "Test"}

			// act
			err := GetExecQuery(query, param, &dest, dbSqlX, &ctx, nil)

			// assert
			assert.Nil(t, err)
			assert.Equal(t, expectedResult, dest)
		})

		t.Run("Should return error when failed to prepare statment to Get Exec Query", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, _ := testfixtures.CreateSqlXDbMock(t)

			dest := repositoryModelTest{}

			expectedResult := repositoryModelTest{}

			// act
			err := GetExecQuery(query, param, &dest, dbSqlX, &ctx, nil)

			// assert
			assert.EqualError(t, err, "all expectations were already fulfilled, call to Prepare 'SELECT id, name FROM TEST WHERE none:=none' query was not expected")
			assert.Equal(t, expectedResult, dest)

		})

		t.Run("Should return error when failed to get context in Get Exec Query", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, mock := testfixtures.CreateSqlXDbMock(t)

			rows := sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "Test")

			mock.ExpectPrepare("^SELECT (.+) FROM TEST WHERE").ExpectQuery().WillReturnRows(rows)

			dest := invalidRepositoryModelTest{}

			expectedResult := invalidRepositoryModelTest{}

			// act
			err := GetExecQuery(query, param, &dest, dbSqlX, &ctx, nil)

			// assert
			assert.EqualError(t, err, "scannable dest type struct with >1 columns (2) in result")
			assert.Equal(t, expectedResult, dest)
		})

		t.Run("Should Get Exec Query with transaction", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, mock := testfixtures.CreateSqlXDbMock(t)

			mock.ExpectBegin()

			tx, _ := dbSqlX.BeginTxx(ctx, nil)

			rows := sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "Test")

			mock.ExpectPrepare("^SELECT (.+) FROM TEST WHERE").ExpectQuery().WillReturnRows(rows)

			dest := repositoryModelTest{}

			expectedResult := repositoryModelTest{Id: 1, Name: "Test"}

			// act
			err := GetExecQuery(query, param, &dest, dbSqlX, &ctx, tx)

			// assert
			assert.Nil(t, err)
			assert.Equal(t, expectedResult, dest)
		})

		t.Run("Should return error when failed to prepare statment to Get Exec Query with transaction", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, mock := testfixtures.CreateSqlXDbMock(t)

			mock.ExpectBegin()

			tx, _ := dbSqlX.BeginTxx(ctx, nil)

			dest := repositoryModelTest{}

			expectedResult := repositoryModelTest{}

			// act
			err := GetExecQuery(query, param, &dest, dbSqlX, &ctx, tx)

			// assert
			assert.EqualError(t, err, "all expectations were already fulfilled, call to Prepare 'SELECT id, name FROM TEST WHERE none:=none' query was not expected")
			assert.Equal(t, expectedResult, dest)

		})

		t.Run("Should return error when failed to get context in Get Exec Query with transaction", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, mock := testfixtures.CreateSqlXDbMock(t)

			mock.ExpectBegin()

			tx, _ := dbSqlX.BeginTxx(ctx, nil)

			rows := sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "Test")

			mock.ExpectPrepare("^SELECT (.+) FROM TEST WHERE").ExpectQuery().WillReturnRows(rows)

			dest := invalidRepositoryModelTest{}

			expectedResult := invalidRepositoryModelTest{}

			// act
			err := GetExecQuery(query, param, &dest, dbSqlX, &ctx, tx)

			// assert
			assert.EqualError(t, err, "scannable dest type struct with >1 columns (2) in result")
			assert.Equal(t, expectedResult, dest)
		})
	})

	t.Run("List Exec Query", func(t *testing.T) {
		t.Parallel()

		t.Run("Should List Exec Query", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, mock := testfixtures.CreateSqlXDbMock(t)

			rows := sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "Test").
				AddRow(2, "Test2")

			mock.ExpectPrepare("^SELECT (.+) FROM TEST WHERE").ExpectQuery().WillReturnRows(rows)

			dest := []repositoryModelTest{}

			expectedResult := []repositoryModelTest{
				{Id: 1, Name: "Test"},
				{Id: 2, Name: "Test2"},
			}

			// act
			err := ListExecQuery(query, param, &dest, dbSqlX, &ctx, nil)

			// assert
			assert.Nil(t, err)
			assert.Equal(t, expectedResult, dest)
		})

		t.Run("Should return error when failed to prepare statment to List Exec Query", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, _ := testfixtures.CreateSqlXDbMock(t)

			dest := []repositoryModelTest{}

			expectedResult := []repositoryModelTest{}

			// act
			err := ListExecQuery(query, param, &dest, dbSqlX, &ctx, nil)

			// assert
			assert.EqualError(t, err, "all expectations were already fulfilled, call to Prepare 'SELECT id, name FROM TEST WHERE none:=none' query was not expected")
			assert.Equal(t, expectedResult, dest)
		})

		t.Run("Should return error when failed to select context in List Exec Query", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, mock := testfixtures.CreateSqlXDbMock(t)

			rows := sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "Test")

			mock.ExpectPrepare("^SELECT (.+) FROM TEST WHERE").ExpectQuery().WillReturnRows(rows)

			dest := []invalidRepositoryModelTest{}

			expectedResult := []invalidRepositoryModelTest{}

			// act
			err := ListExecQuery(query, param, &dest, dbSqlX, &ctx, nil)

			// assert
			assert.EqualError(t, err, "non-struct dest type struct with >1 columns (2)")
			assert.Equal(t, expectedResult, dest)
		})

		t.Run("Should List Exec Query with transaction", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, mock := testfixtures.CreateSqlXDbMock(t)

			mock.ExpectBegin()

			tx, _ := dbSqlX.BeginTxx(ctx, nil)

			rows := sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "Test").
				AddRow(2, "Test2")

			mock.ExpectPrepare("^SELECT (.+) FROM TEST WHERE").ExpectQuery().WillReturnRows(rows)

			dest := []repositoryModelTest{}

			expectedResult := []repositoryModelTest{
				{Id: 1, Name: "Test"},
				{Id: 2, Name: "Test2"},
			}

			// act
			err := ListExecQuery(query, param, &dest, dbSqlX, &ctx, tx)

			// assert
			assert.Nil(t, err)
			assert.Equal(t, expectedResult, dest)
		})

		t.Run("Should return error when failed to prepare statment to List Exec Query with transaction", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, mock := testfixtures.CreateSqlXDbMock(t)

			mock.ExpectBegin()

			tx, _ := dbSqlX.BeginTxx(ctx, nil)

			dest := []repositoryModelTest{}

			expectedResult := []repositoryModelTest{}

			// act
			err := ListExecQuery(query, param, &dest, dbSqlX, &ctx, tx)

			// assert
			assert.EqualError(t, err, "all expectations were already fulfilled, call to Prepare 'SELECT id, name FROM TEST WHERE none:=none' query was not expected")
			assert.Equal(t, expectedResult, dest)
		})

		t.Run("Should return error when failed to select context in List Exec Query with transaction", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, mock := testfixtures.CreateSqlXDbMock(t)

			mock.ExpectBegin()

			tx, _ := dbSqlX.BeginTxx(ctx, nil)

			rows := sqlmock.NewRows([]string{"id", "name"}).
				AddRow(1, "Test")

			mock.ExpectPrepare("^SELECT (.+) FROM TEST WHERE").ExpectQuery().WillReturnRows(rows)

			dest := []invalidRepositoryModelTest{}

			expectedResult := []invalidRepositoryModelTest{}

			// act
			err := ListExecQuery(query, param, &dest, dbSqlX, &ctx, tx)

			// assert
			assert.EqualError(t, err, "non-struct dest type struct with >1 columns (2)")
			assert.Equal(t, expectedResult, dest)
		})
	})
}
