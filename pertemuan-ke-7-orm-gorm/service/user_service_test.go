package service_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"praisindo/entity"
	"praisindo/service"
	mock_service "praisindo/test/mock/service"
	"testing"
	"time"
)

func TestUserService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_service.NewMockIUserRepository(ctrl)
	userService := service.NewUserService(mockRepo)

	ctx := context.Background()
	user := &entity.User{
		Name:      "John Doe",
		Email:     "john@example.com",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("PositiveCase", func(t *testing.T) {
		mockRepo.EXPECT().CreateUser(ctx, user).Return(*user, nil)

		createdUser, err := userService.CreateUser(ctx, user)
		assert.NoError(t, err)
		assert.Equal(t, *user, createdUser)
	})

	t.Run("NegativeCase", func(t *testing.T) {
		mockRepo.EXPECT().CreateUser(ctx, user).Return(entity.User{}, errors.New("failed to create user"))

		createdUser, err := userService.CreateUser(ctx, user)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create user")
		assert.Equal(t, entity.User{}, createdUser)
	})
}
