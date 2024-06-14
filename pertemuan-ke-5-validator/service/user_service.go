package service

import (
	"fmt"
	"praisindo/entity"
)

// IUserService mendefinisikan interface untuk layanan pengguna
type IUserService interface {
	CreateUser(user *entity.User) entity.User
	GetUserByID(id int) (entity.User, error)
	UpdateUser(id int, user entity.User) (entity.User, error)
	DeleteUser(id int) error
	GetAllUsers() []entity.User
}

// IUserRepository mendefinisikan interface untuk repository pengguna
type IUserRepository interface {
	CreateUser(user *entity.User) entity.User
	GetUserByID(id int) (entity.User, bool)
	UpdateUser(id int, user entity.User) (entity.User, bool)
	DeleteUser(id int) bool
	GetAllUsers() []entity.User
}

// userService adalah implementasi dari IUserService yang menggunakan IUserRepository
type userService struct {
	userRepo IUserRepository
}

// NewUserService membuat instance baru dari userService
func NewUserService(userRepo IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

// CreateUser membuat pengguna baru dengan menggunakan repository
func (s *userService) CreateUser(user *entity.User) entity.User {
	return s.userRepo.CreateUser(user)
}

// GetUserByID mendapatkan pengguna berdasarkan ID, mengembalikan error jika tidak ditemukan
func (s *userService) GetUserByID(id int) (entity.User, error) {
	user, found := s.userRepo.GetUserByID(id)
	if !found {
		return entity.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

// UpdateUser memperbarui pengguna berdasarkan ID, mengembalikan error jika tidak ditemukan
func (s *userService) UpdateUser(id int, user entity.User) (entity.User, error) {
	updatedUser, found := s.userRepo.UpdateUser(id, user)
	if !found {
		return entity.User{}, fmt.Errorf("user not found")
	}
	return updatedUser, nil
}

// DeleteUser menghapus pengguna berdasarkan ID, mengembalikan error jika tidak ditemukan
func (s *userService) DeleteUser(id int) error {
	if !s.userRepo.DeleteUser(id) {
		return fmt.Errorf("user not found")
	}
	return nil
}

// GetAllUsers mengembalikan semua pengguna yang ada di repository
func (s *userService) GetAllUsers() []entity.User {
	return s.userRepo.GetAllUsers()
}
