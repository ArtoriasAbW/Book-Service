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
	// arrange
	t.Run("success", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		queryStore := regexp.QuoteMeta(`INSERT INTO users (username) VALUES ($1) RETURNING id`)
		rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
		f.dbMock.ExpectQuery(queryStore).WithArgs("username").WillReturnRows(rows)

		// act
		id, err := f.bookRepo.AddUser(context.Background(), models.User{
			Username: "username",
		})

		// assert
		require.NoError(t, err)
		assert.Equal(t, uint64(1), id)
	})

	t.Run("fail", func(t *testing.T) {
		f := setUp(t)
		defer f.tearDown()

		queryStore := regexp.QuoteMeta(`INSERT INTO users (username) VALUES ($1) RETURNING id`)
		f.dbMock.ExpectQuery(queryStore).WithArgs("username").WillReturnError(errors.New("some error"))

		// act
		_, err := f.bookRepo.AddUser(context.Background(), models.User{
			Username: "username",
		})

		// assert
		assert.Equal(t, err, fmt.Errorf("Repository.AddUser: to sql: %w", errors.New("some error")))

	})
}
