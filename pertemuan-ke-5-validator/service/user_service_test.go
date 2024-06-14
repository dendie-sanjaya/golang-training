package service_test

import (
	//"github.com/ibrahimker/golang-praisindo-advanced/session-4-unit-test-crud-user/test/mock/repository"
	"github.com/stretchr/testify/assert"
	"praisindo/entity"
	"praisindo/service"
	"testing"
)

/*
func TestCreateUser(t *testing.T) {
	mockRepo := &repository.MockUserRepository{}
	userService := service.NewUserService(mockRepo)

	t.Run("CreateUser - Success", func(t *testing.T) {
		user := &entity.User{Name: "Test", Email: "test@example.com", Password: "password"}
		createdUser := userService.CreateUser(user)

		assert.Equal(t, 0, createdUser.ID)
		assert.Equal(t, "Test", createdUser.Name)
		assert.NotZero(t, createdUser.CreatedAt)
		assert.NotZero(t, createdUser.UpdatedAt)
	})
}
*/

func TestGetUserByID(t *testing.T) {
	mockRepo := &repository.MockUserRepository{}
	userService := service.NewUserService(mockRepo)

	user := &entity.User{Name: "Test", Email: "test@example.com", Password: "password"}
	createdUser := userService.CreateUser(user)

	t.Run("GetUserByID - Success", func(t *testing.T) {
		retrievedUser, err := userService.GetUserByID(createdUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, createdUser.ID, retrievedUser.ID)
		assert.Equal(t, createdUser.Name, retrievedUser.Name)
	})

	t.Run("GetUserByID - UserNotFound", func(t *testing.T) {
		_, err := userService.GetUserByID(99)
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}

func TestUpdateUser(t *testing.T) {
	mockRepo := &repository.MockUserRepository{}
	userService := service.NewUserService(mockRepo)

	user := &entity.User{Name: "Test", Email: "test@example.com", Password: "password"}
	createdUser := userService.CreateUser(user)

	t.Run("UpdateUser - Success", func(t *testing.T) {
		updatedUser := entity.User{Name: "Updated", Email: "updated@example.com", Password: "password"}
		result, err := userService.UpdateUser(createdUser.ID, updatedUser)

		assert.NoError(t, err)
		assert.Equal(t, "Updated", result.Name)
		assert.Equal(t, "updated@example.com", result.Email)
	})

	t.Run("UpdateUser - UserNotFound", func(t *testing.T) {
		updatedUser := entity.User{Name: "Updated", Email: "updated@example.com", Password: "password"}
		_, err := userService.UpdateUser(99, updatedUser)

		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}

func TestDeleteUser(t *testing.T) {
	mockRepo := &repository.MockUserRepository{}
	userService := service.NewUserService(mockRepo)

	user := &entity.User{Name: "Test", Email: "test@example.com", Password: "password"}
	createdUser := userService.CreateUser(user)

	t.Run("DeleteUser - Success", func(t *testing.T) {
		err := userService.DeleteUser(createdUser.ID)
		assert.NoError(t, err)
	})

	t.Run("DeleteUser - UserNotFound", func(t *testing.T) {
		err := userService.DeleteUser(99)
		assert.Error(t, err)
		assert.Equal(t, "user not found", err.Error())
	})
}

func TestGetAllUsers(t *testing.T) {
	mockRepo := &repository.MockUserRepository{}
	userService := service.NewUserService(mockRepo)

	user1 := &entity.User{Name: "Test1", Email: "test1@example.com", Password: "password"}
	user2 := &entity.User{Name: "Test2", Email: "test2@example.com", Password: "password"}

	userService.CreateUser(user1)
	userService.CreateUser(user2)

	t.Run("GetAllUsers - Success", func(t *testing.T) {
		users := userService.GetAllUsers()
		assert.Equal(t, 2, len(users))
		assert.Equal(t, "Test1", users[0].Name)
		assert.Equal(t, "Test2", users[1].Name)
	})
}
