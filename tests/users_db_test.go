//go:build integration
// +build integration

package tests

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/models"
	repository "gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/repository/postgres"
)

func TestCreateUser(t *testing.T) {
	t.Run("success add user", func(t *testing.T) {
		// arrange
		Db.Setup(t)
		defer Db.TearDown()
		repo := repository.NewRepository(Db.DB)
		// act
		_, err := repo.AddUser(context.Background(), models.User{
			Username: "user1",
		})
		// assert
		assert.NoError(t, err)
	})
}

func TestGetUser(t *testing.T) {
	t.Run("success get user", func(t *testing.T) {
		// arrange
		Db.Setup(t)
		defer Db.TearDown()
		repo := repository.NewRepository(Db.DB)
		ctx := context.Background()
		id, _ := repo.AddUser(ctx, models.User{
			Username: "user1",
		})
		// act
		user, err := repo.GetUserByID(ctx, uint(id))

		// assert
		require.NoError(t, err)
		assert.EqualValues(t, id, user.ID)
		assert.Equal(t, "user1", user.Username)

	})
	t.Run("no user in db", func(t *testing.T) {
		// arrange
		Db.Setup(t)
		defer Db.TearDown()
		repo := repository.NewRepository(Db.DB)
		ctx := context.Background()
		id := 1
		_, err := repo.GetUserByID(ctx, uint(id))

		assert.EqualError(t, err, "no user")
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("success delete user", func(t *testing.T) {
		// arrange
		Db.Setup(t)
		defer Db.TearDown()
		repo := repository.NewRepository(Db.DB)
		ctx := context.Background()
		id, _ := repo.AddUser(ctx, models.User{
			Username: "user1",
		})
		// act
		err := repo.DeleteUser(ctx, uint(id))

		// assert
		require.NoError(t, err)

		_, err = repo.GetUserByID(ctx, uint(id))

		assert.EqualError(t, err, "no user")
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("success update user", func(t *testing.T) {
		Db.Setup(t)
		defer Db.TearDown()
		repo := repository.NewRepository(Db.DB)
		ctx := context.Background()
		id, _ := repo.AddUser(ctx, models.User{
			Username: "user1",
		})
		newUser := models.User{
			ID:       uint(id),
			Username: "new_username",
		}
		err := repo.UpdateUser(ctx, newUser)

		require.NoError(t, err)

		user, err := repo.GetUserByID(ctx, uint(id))

		require.NoError(t, err)
		assert.EqualValues(t, id, user.ID)
		assert.Equal(t, newUser.Username, user.Username)
	})
}

func TestListUsers(t *testing.T) {
	t.Run("list users", func(t *testing.T) {
		Db.Setup(t)
		defer Db.TearDown()

		repo := repository.NewRepository(Db.DB)
		ctx := context.Background()
		for i := 1; i < 6; i++ {
			repo.AddUser(ctx, models.User{
				Username: "user" + strconv.Itoa(i),
			})
		}

		listParams := models.ListInput{
			Limit:  3,
			Offset: 2,
			Order:  "DESC",
		}

		users, err := repo.ListUsers(ctx, listParams)

		require.NoError(t, err)
		require.Len(t, users, 3)
		assert.Equal(t, users[0].Username, "user3")
		assert.Equal(t, users[1].Username, "user2")
		assert.Equal(t, users[2].Username, "user1")

	})
}
