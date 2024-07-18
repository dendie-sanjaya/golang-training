package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"praisindo/entity"
	"time"

	"github.com/redis/go-redis/v9"
)

// IUserService mendefinisikan interface untuk layanan pengguna
type IUserService interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

// IUserRepository mendefinisikan interface untuk repository pengguna
type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)
}

// userService adalah implementasi dari IUserService yang menggunakan IUserRepository
type userService struct {
	userRepo IUserRepository
	rdb      *redis.Client
}

// NewUserService membuat instance baru dari userService
func NewUserService(userRepo IUserRepository, rdb *redis.Client) IUserService {
	return &userService{userRepo: userRepo, rdb: rdb}
}

// CreateUser membuat pengguna baru
func (s *userService) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	// Memanggil CreateUser dari repository untuk membuat pengguna baru
	createdUser, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("gagal membuat pengguna: %v", err)
	}

	redisKey := fmt.Sprintf("createdUser:%d", createdUser.ID)
	createdUserJSON, err := json.Marshal(createdUser)
	if err != nil {
		log.Println("gagal marshal json")
	}
	if err := s.rdb.Set(ctx, redisKey, createdUserJSON, 60*time.Second).Err(); err != nil {
		log.Println("error when set redis key", redisKey)
	}

	return createdUser, nil
}

// GetUserByID mendapatkan pengguna berdasarkan ID
func (s *userService) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	// Memanggil GetUserByID dari repository untuk mendapatkan pengguna berdasarkan ID
	user, err := s.userRepo.GetUserByID(ctx, id)

	redisKey := fmt.Sprintf("createdUser:%d", id)
	val, err := s.rdb.Get(ctx, redisKey).Result()
	if err == nil {
		log.Println("data tersedia di redis")
		err = json.Unmarshal([]byte(val), &user)
		if err != nil {
			log.Println("gagal marshall data di redis, coba ambil ke database")
		}
	}

	if err != nil {
		log.Println("data tidak tersedia di redis, ambil dari database")
		user, err = s.userRepo.GetUserByID(ctx, id)
		if err != nil {
			log.Println("gagal ambil data di database")
			return entity.User{}, fmt.Errorf("gagal mendapatkan pengguna berdasarkan ID: %v", err)
		}
	}

	return user, nil
}

// UpdateUser memperbarui data pengguna
func (s *userService) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	// Memanggil UpdateUser dari repository untuk memperbarui data pengguna
	updatedUser, err := s.userRepo.UpdateUser(ctx, id, user)
	if err != nil {
		return entity.User{}, fmt.Errorf("gagal memperbarui pengguna: %v", err)
	}
	return updatedUser, nil
}

// DeleteUser menghapus pengguna berdasarkan ID
func (s *userService) DeleteUser(ctx context.Context, id int) error {
	// Memanggil DeleteUser dari repository untuk menghapus pengguna berdasarkan ID
	err := s.userRepo.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus pengguna: %v", err)
	}
	return nil
}

// GetAllUsers mendapatkan semua pengguna
func (s *userService) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	// Memanggil GetAllUsers dari repository untuk mendapatkan semua pengguna
	users, err := s.userRepo.GetAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan semua pengguna: %v", err)
	}
	return users, nil
}
