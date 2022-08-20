package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"gitlab.ozon.dev/ArtoriasAbW/homework-01/internal/pkg/service/models"
	pb "gitlab.ozon.dev/ArtoriasAbW/homework-01/pkg/api"
)

func TestNewUserService(t *testing.T) {
	t.Run("success add user", func(t *testing.T) {
		// arrange
		f := serviceSetUp(t)
		user := models.User{
			Username: "some username",
		}
		// mock repository call
		req := &pb.UserCreateRequest{
			Username: "some username",
		}
		f.repo.EXPECT().UserCreate(gomock.Any(), req).Return(&pb.UserCreateResponse{Id: 1}, nil).Times(1)
		// act
		resp, err := f.service.AddUser(f.Ctx, user)
		// assert
		require.NoError(t, err)
		assert.Equal(t, resp, uint64(1))
	})

	t.Run("add user with error [empty username]", func(t *testing.T) {
		// arrange
		f := serviceSetUp(t)
		user := models.User{
			Username: "",
		}
		req := &pb.UserCreateRequest{
			Username: "",
		}
		f.repo.EXPECT().UserCreate(gomock.Any(), req).Return(&pb.UserCreateResponse{Id: 1}, nil).Times(0)

		// act
		_, err := f.service.AddUser(f.Ctx, user)
		// assert
		assert.Equal(t, err.Error(), errors.Wrap(ErrValidation, "field: [username] cannot be empty").Error())
	})
	t.Run("add user with error [storage service error]", func(t *testing.T) {
		// arrange
		f := serviceSetUp(t)
		user := models.User{
			Username: "user1",
		}
		req := &pb.UserCreateRequest{
			Username: "user1",
		}
		f.repo.EXPECT().UserCreate(gomock.Any(), req).Return(nil, errors.New("some error")).Times(1)
		// act
		_, err := f.service.AddUser(f.Ctx, user)
		// assert
		assert.Equal(t, err.Error(), errors.New("some error").Error())
	})
}

func TestGetUser(t *testing.T) {
	t.Run("success get user", func(t *testing.T) {
		// arrange
		f := serviceSetUp(t)
		id := uint(1)
		req := &pb.UserGetRequest{
			Id: uint64(id),
		}
		f.repo.EXPECT().UserGet(gomock.Any(), req).Return(&pb.UserGetResponse{Id: 1, Username: "some"}, nil).Times(1)
		// act
		resp, err := f.service.GetUser(f.Ctx, id)
		// assert
		require.NoError(t, err)
		assert.Equal(t, resp, models.User{
			Id:       1,
			Username: "some",
		})
	})

	t.Run("get user with error [storage service error]", func(t *testing.T) {
		// arrange
		f := serviceSetUp(t)
		id := uint(1)
		req := &pb.UserGetRequest{
			Id: uint64(id),
		}
		f.repo.EXPECT().UserGet(gomock.Any(), req).Return(nil, errors.New("some error")).Times(1)
		// act
		_, err := f.service.GetUser(f.Ctx, id)
		// assert
		assert.Equal(t, err.Error(), errors.New("some error").Error())
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("success delete user", func(t *testing.T) {
		// arrange
		f := serviceSetUp(t)
		id := uint(1)
		req := &pb.UserDeleteRequest{
			Id: uint64(id),
		}
		f.repo.EXPECT().UserDelete(gomock.Any(), req).Return(&pb.UserDeleteResponse{}, nil).Times(1)

		// act
		err := f.service.DeleteUser(f.Ctx, id)
		// assert
		require.NoError(t, err)
	})

	t.Run("delete user with error [storage service error]", func(t *testing.T) {
		// arrange
		f := serviceSetUp(t)
		id := uint(1)
		req := &pb.UserDeleteRequest{
			Id: uint64(id),
		}
		f.repo.EXPECT().UserDelete(gomock.Any(), req).Return(nil, errors.New("some error")).Times(1)

		// act
		err := f.service.DeleteUser(f.Ctx, id)

		// assert
		assert.Equal(t, "some error", err.Error())
	})
}

func TestUserUpdate(t *testing.T) {
	t.Run("success update user", func(t *testing.T) {
		// arrange
		f := serviceSetUp(t)
		newUser := models.User{
			Id:       1,
			Username: "new_username",
		}
		oldUser := models.User{
			Id:       1,
			Username: "old_username",
		}
		req := &pb.UserUpdateRequest{
			Id:       uint64(newUser.Id),
			Username: newUser.Username,
		}
		getReq := &pb.UserGetRequest{
			Id: uint64(oldUser.Id),
		}
		f.repo.EXPECT().UserUpdate(gomock.Any(), req).Return(&pb.UserUpdateResponse{}, nil).Times(1)
		f.repo.EXPECT().UserGet(gomock.Any(), getReq).
			Return(&pb.UserGetResponse{Id: uint64(oldUser.Id), Username: oldUser.Username}, nil).
			Times(1)

		// act
		err := f.service.UpdateUser(f.Ctx, newUser)

		// assert
		require.NoError(t, err)
	})
	t.Run("user update with error [storage service user get returns error]", func(t *testing.T) {
		// arrange
		f := serviceSetUp(t)
		newUser := models.User{
			Id:       1,
			Username: "new_username",
		}
		oldUser := models.User{
			Id:       1,
			Username: "old_username",
		}
		req := &pb.UserUpdateRequest{
			Id:       uint64(newUser.Id),
			Username: newUser.Username,
		}
		getReq := &pb.UserGetRequest{
			Id: uint64(oldUser.Id),
		}
		f.repo.EXPECT().UserUpdate(gomock.Any(), req).Return(&pb.UserUpdateResponse{}, nil).Times(0)
		f.repo.EXPECT().UserGet(gomock.Any(), getReq).
			Return(nil, errors.New("some error")).
			Times(1)

		// act
		err := f.service.UpdateUser(f.Ctx, newUser)

		// assert
		require.EqualError(t, err, "user with this id doesn't exist: invalid data")
	})
	t.Run("update user with error [empty username]", func(t *testing.T) {
		// arrange
		f := serviceSetUp(t)
		newUser := models.User{
			Id:       1,
			Username: "",
		}
		oldUser := models.User{
			Id:       1,
			Username: "old_username",
		}
		req := &pb.UserUpdateRequest{
			Id:       uint64(newUser.Id),
			Username: newUser.Username,
		}
		getReq := &pb.UserGetRequest{
			Id: uint64(oldUser.Id),
		}
		f.repo.EXPECT().UserUpdate(gomock.Any(), req).Return(&pb.UserUpdateResponse{}, nil).Times(0)
		f.repo.EXPECT().UserGet(gomock.Any(), getReq).
			Return(&pb.UserGetResponse{Id: uint64(oldUser.Id), Username: oldUser.Username}, nil).
			Times(0)
		// act
		err := f.service.UpdateUser(f.Ctx, newUser)
		require.EqualError(t, err, "field: [username] cannot be empty: invalid data")
	})
}

func TestUserList(t *testing.T) {
	t.Run("success list users", func(t *testing.T) {
		// arrange
		f := serviceSetUp(t)
		params := models.ListInput{
			Offset: 1,
			Limit:  2,
			Order:  "desc",
		}
		req := &pb.UserListRequest{
			Offset: params.Offset,
			Limit:  params.Limit,
			Order:  "DESC",
		}
		usersResp := []*pb.UserListResponse_User{{Id: 1, Username: "user1"}, {Id: 2, Username: "user2"}}
		f.repo.EXPECT().UserList(gomock.Any(), req).Return(&pb.UserListResponse{Users: usersResp}, nil).Times(1)

		// act
		users, err := f.service.ListUsers(f.Ctx, params)
		require.NoError(t, err)
		require.Len(t, users, len(usersResp))
		assert.Equal(t, users[0].Id, uint(usersResp[0].Id))
		assert.Equal(t, users[0].Username, usersResp[0].Username)
		assert.Equal(t, users[1].Id, uint(usersResp[1].Id))
		assert.Equal(t, users[1].Username, usersResp[1].Username)
	})
	t.Run("list users with error [storage service error]", func(t *testing.T) {
		// arrange
		f := serviceSetUp(t)
		params := models.ListInput{
			Offset: 1,
			Limit:  2,
			Order:  "ads",
		}
		req := &pb.UserListRequest{
			Offset: params.Offset,
			Limit:  params.Limit,
			Order:  "ASC",
		}
		f.repo.EXPECT().UserList(gomock.Any(), req).Return(nil, errors.New("some error")).Times(1)

		// act
		_, err := f.service.ListUsers(f.Ctx, params)
		require.EqualError(t, err, "some error")
	})
}
