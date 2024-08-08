package postgres_gorm_raw

import (
	"context"
	"errors"
	"fmt"
	"log"
	"praisindo/entity"
	"praisindo/service"

	"gorm.io/gorm"
)

// GormDBIface defines an interface for GORM DB methods used in the repository
type GormDBIface interface {
	WithContext(ctx context.Context) *gorm.DB
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}

type userRepository struct {
	db GormDBIface
}

// NewUserRepository membuat instance baru dari userRepository
func NewUserRepository(db GormDBIface) service.IUserRepository {
	return &userRepository{db: db}
}

// CreateUser membuat pengguna baru dalam basis data
func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	fmt.Println("Gorm Raw CreateUser")
	query := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id"
	//query := fmt.Sprintf("INSERT INTO users (name, email, password, created_at, updated_at)
	// VALUES ('$1', $2, $3, NOW(), NOW()) RETURNING id",user.Name)
	var createdID int
	if err := r.db.WithContext(ctx).
		Raw(query, user.Name, user.Email, user.Password).Scan(&createdID).Error; err != nil {
		log.Printf("Error creating user: %v\n", err)
		return entity.User{}, err
	}
	user.ID = createdID
	return *user, nil
}

// GetUserByID mengambil pengguna berdasarkan ID
func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1"
	if err := r.db.WithContext(ctx).Raw(query, id).Scan(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, nil
		}
		log.Printf("Error getting user by ID: %v\n", err)
		return entity.User{}, err
	}
	return user, nil
}

// UpdateUser memperbarui informasi pengguna dalam basis data
func (r *userRepository) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	query := "UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3"
	// Memperbarui informasi pengguna
	if err := r.db.WithContext(ctx).Exec(query, user.Name, user.Email, id).Error; err != nil {
		log.Printf("Error updating user: %v\n", err)
		return entity.User{}, err
	}
	return user, nil
}

// DeleteUser menghapus pengguna berdasarkan ID
func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	if err := r.db.WithContext(ctx).Exec(query, id).Error; err != nil {
		log.Printf("Error deleting user: %v\n", err)
		return err
	}
	return nil
}

// GetAllUsers mengambil semua pengguna dari basis data
func (r *userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	query := "SELECT id, name, email, password, created_at, updated_at FROM users"
	var users []entity.User
	if err := r.db.WithContext(ctx).Raw(query).Scan(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, nil
		}
		log.Printf("Error getting all users: %v\n", err)
		return nil, err
	}
	return users, nil
}
