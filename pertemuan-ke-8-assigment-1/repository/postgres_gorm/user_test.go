package postgres_gorm_test

import (
	"context"
	"praisindo/entity"
	"praisindo/repository/postgres_gorm"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MockGormDBIface adalah mock dari GormDBIface
type MockGormDBIface struct {
	mock.Mock
	*gorm.DB
}

func NewMockGormDBIface() *MockGormDBIface {
	// Gunakan database SQLite dalam memory untuk membuat dummy *gorm.DB
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	return &MockGormDBIface{DB: db}
}

func (m *MockGormDBIface) WithContext(ctx context.Context) *gorm.DB {
	args := m.Called(ctx)
	return args.Get(0).(*gorm.DB)
}

func (m *MockGormDBIface) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockGormDBIface) First(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func (m *MockGormDBIface) Save(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockGormDBIface) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(value, conds)
	return args.Get(0).(*gorm.DB)
}

func (m *MockGormDBIface) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func createDummyDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	return db
}

// TestCreateUser untuk menguji fungsi CreateUser
func TestCreateUser(t *testing.T) {
	mockDB := NewMockGormDBIface()
	ctx := context.TODO()
	user := &entity.User{ID: 1, Name: "John Doe", Email: "john.doe@example.com"}

	dummyDB := createDummyDB()

	mockDB.On("WithContext", ctx).Return(dummyDB)
	mockDB.On("Create", user).Return(dummyDB)

	repo := postgres_gorm.NewUserRepository(mockDB)
	result, err := repo.CreateUser(ctx, user)

	assert.Nil(t, err)
	assert.Equal(t, *user, result)
	mockDB.AssertExpectations(t)
}

// TestGetUserByID untuk menguji fungsi GetUserByID
func TestGetUserByID(t *testing.T) {
	mockDB := NewMockGormDBIface()
	ctx := context.TODO()
	user := entity.User{ID: 1, Name: "John Doe", Email: "john.doe@example.com"}

	dummyDB := createDummyDB()

	mockDB.On("WithContext", ctx).Return(dummyDB)
	mockDB.On("First", mock.Anything, 1).Return(dummyDB).Run(func(args mock.Arguments) {
		arg := reflect.ValueOf(args.Get(0)).Elem()
		arg.Set(reflect.ValueOf(user))
	})

	repo := postgres_gorm.NewUserRepository(mockDB)
	result, err := repo.GetUserByID(ctx, 1)

	assert.Nil(t, err)
	assert.Equal(t, user, result)
	mockDB.AssertExpectations(t)
}

// TestUpdateUser untuk menguji fungsi UpdateUser
func TestUpdateUser(t *testing.T) {
	mockDB := NewMockGormDBIface()
	ctx := context.TODO()
	existingUser := entity.User{ID: 1, Name: "John Doe", Email: "john.doe@example.com"}
	updatedUser := entity.User{Name: "John Updated", Email: "john.updated@example.com"}

	dummyDB := createDummyDB()

	mockDB.On("WithContext", ctx).Return(dummyDB)
	mockDB.On("First", mock.Anything, 1).Return(dummyDB).Run(func(args mock.Arguments) {
		arg := reflect.ValueOf(args.Get(0)).Elem()
		arg.Set(reflect.ValueOf(existingUser))
	})
	mockDB.On("Save", mock.Anything).Return(dummyDB)

	repo := postgres_gorm.NewUserRepository(mockDB)
	result, err := repo.UpdateUser(ctx, 1, updatedUser)

	assert.Nil(t, err)
	assert.Equal(t, updatedUser.Name, result.Name)
	assert.Equal(t, updatedUser.Email, result.Email)
	mockDB.AssertExpectations(t)
}

// TestDeleteUser untuk menguji fungsi DeleteUser
func TestDeleteUser(t *testing.T) {
	mockDB := NewMockGormDBIface()
	ctx := context.TODO()

	dummyDB := createDummyDB()

	mockDB.On("WithContext", ctx).Return(dummyDB)
	mockDB.On("Delete", mock.Anything, 1).Return(dummyDB)

	repo := postgres_gorm.NewUserRepository(mockDB)
	err := repo.DeleteUser(ctx, 1)

	assert.Nil(t, err)
	mockDB.AssertExpectations(t)
}

// TestGetAllUsers untuk menguji fungsi GetAllUsers
func TestGetAllUsers(t *testing.T) {
	mockDB := NewMockGormDBIface()
	ctx := context.TODO()

	users := []entity.User{
		{ID: 1, Name: "John Doe", Email: "john.doe@example.com"},
		{ID: 2, Name: "Jane Doe", Email: "jane.doe@example.com"},
	}

	dummyDB := createDummyDB()

	mockDB.On("WithContext", ctx).Return(dummyDB)
	mockDB.On("Find", mock.Anything).Return(dummyDB).Run(func(args mock.Arguments) {
		arg := reflect.ValueOf(args.Get(0)).Elem()
		arg.Set(reflect.ValueOf(users))
	})

	repo := postgres_gorm.NewUserRepository(mockDB)
	result, err := repo.GetAllUsers(ctx)

	assert.Nil(t, err)
	assert.Equal(t, users, result)
	mockDB.AssertExpectations(t)
}
