package postgres

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
)

func TestNewUsers(t *testing.T) {
	t.Run("success user add", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()
		queryStore := regexp.QuoteMeta(`INSERT INTO users (username) VALUES ($1) RETURNING id`)
		rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
		f.dbMock.ExpectQuery(queryStore).WithArgs("username").WillReturnRows(rows)

		// act
		id, err := f.repo.AddUser(context.Background(), models.User{
			Username: "username",
		})

		// assert
		require.NoError(t, err)
		assert.Equal(t, uint64(1), id)
	})

	t.Run("user add with error", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()
		queryStore := regexp.QuoteMeta(`INSERT INTO users (username) VALUES ($1) RETURNING id`)
		f.dbMock.ExpectQuery(queryStore).WithArgs("username").WillReturnError(errors.New("some error"))

		// act
		_, err := f.repo.AddUser(context.Background(), models.User{
			Username: "username",
		})

		// assert
		assert.Equal(t, err, fmt.Errorf("Repository.AddUser: to sql: %w", errors.New("some error")))

	})
}

func TestGetUser(t *testing.T) {
	t.Run("success user get", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()
		queryStore := regexp.QuoteMeta(`SELECT id, username FROM users WHERE id = $1`)
		id := uint(1)
		username := "user1"
		rows := sqlmock.NewRows([]string{"id", "username"}).AddRow(id, username)
		f.dbMock.ExpectQuery(queryStore).WithArgs(id).WillReturnRows(rows)

		// act
		user, err := f.repo.GetUserById(context.Background(), id)

		// assert
		require.NoError(t, err)
		assert.Equal(t, models.User{
			Id:       id,
			Username: username,
		}, user)
	})
	t.Run("user get with error", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()
		queryStore := regexp.QuoteMeta(`SELECT id, username FROM users WHERE id = $1`)
		id := uint(1)
		f.dbMock.ExpectQuery(queryStore).WithArgs(id).WillReturnError(errors.New("some error"))

		// act
		_, err := f.repo.GetUserById(context.Background(), id)

		// assert
		assert.Equal(t, err, fmt.Errorf("Repository.GetUserById: to sql: %w", errors.New("some error")))
	})
	t.Run("invalid data in user table", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()
		queryStore := regexp.QuoteMeta(`SELECT id, username FROM users WHERE id = $1`)
		id := uint(1)
		rows := sqlmock.NewRows([]string{"id", "username"}).AddRow("dadasd", "dasdasd")
		f.dbMock.ExpectQuery(queryStore).WithArgs(id).WillReturnRows(rows)

		// act
		_, err := f.repo.GetUserById(context.Background(), id)

		// assert
		assert.Equal(t, err, fmt.Errorf("Repository.GetUserById: to sql: invalid struct scan"))
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("success user update", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()
		queryStore := regexp.QuoteMeta(`UPDATE users SET username = $1 WHERE id = $2`)
		id := uint(1)
		username := "user2"
		f.dbMock.ExpectExec(queryStore).WithArgs(username, id).WillReturnResult(sqlmock.NewResult(1, 1))

		// act
		err := f.repo.UpdateUser(context.Background(), models.User{
			Id:       id,
			Username: username,
		})
		assert.NoError(t, err)
	})
	t.Run("user update with error", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()
		queryStore := regexp.QuoteMeta(`UPDATE users SET username = $1 WHERE id = $2`)
		id := uint(1)
		username := "user2"
		f.dbMock.ExpectExec(queryStore).WithArgs(username, id).WillReturnError(errors.New("some error"))

		// act
		err := f.repo.UpdateUser(context.Background(), models.User{
			Id:       id,
			Username: username,
		})
		assert.Equal(t, err, fmt.Errorf("Repository.UpdateUser: to sql: %w", errors.New("some error")))
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("success user delete", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()
		queryStore := regexp.QuoteMeta(`DELETE FROM users WHERE id = $1`)
		id := uint(1)
		f.dbMock.ExpectExec(queryStore).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

		// act
		err := f.repo.DeleteUser(context.Background(), id)

		// assert
		assert.NoError(t, err)
	})
	t.Run("user delete with error", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()
		queryStore := regexp.QuoteMeta(`DELETE FROM users WHERE id = $1`)
		id := uint(1)
		f.dbMock.ExpectExec(queryStore).WithArgs(id).WillReturnError(errors.New("some error"))

		// act
		err := f.repo.DeleteUser(context.Background(), id)

		// assert
		assert.Equal(t, err, fmt.Errorf("Repository.DeleteUser: to sql: %w", errors.New("some error")))
	})
}

func TestListUsers(t *testing.T) {
	t.Run("success user list", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()

		input := models.ListInput{
			Offset: 1,
			Limit:  2,
			Order:  "ACK",
		}
		queryStore := regexp.QuoteMeta(`SELECT id, username FROM users ORDER BY username ACK LIMIT 2 OFFSET 1`)
		userRows := []models.User{{Id: 1, Username: "user1"}, {Id: 2, Username: "user2"}}
		rows := sqlmock.NewRows([]string{"id", "username"}).
			AddRow(userRows[0].Id, userRows[0].Username).
			AddRow(userRows[1].Id, userRows[1].Username)
		f.dbMock.ExpectQuery(queryStore).WillReturnRows(rows)

		// act
		users, err := f.repo.ListUsers(context.Background(), input)

		// assert
		require.NoError(t, err)
		require.Len(t, users, 2)
		assert.Equal(t, users[0], userRows[0])
		assert.Equal(t, users[1], userRows[1])

	})
	t.Run("invalid user with error", func(t *testing.T) {
		// arrange
		f := setUp(t)
		defer f.tearDown()

		input := models.ListInput{
			Offset: 1,
			Limit:  2,
			Order:  "ACK",
		}
		queryStore := regexp.QuoteMeta(`SELECT id, username FROM users ORDER BY username ACK LIMIT 2 OFFSET 1`)
		f.dbMock.ExpectQuery(queryStore).WillReturnError(errors.New("some error"))

		// act
		_, err := f.repo.ListUsers(context.Background(), input)
		assert.Equal(t, err, fmt.Errorf("Repository.ListUsers: to sql: %w", errors.New("some error")))

	})
}
