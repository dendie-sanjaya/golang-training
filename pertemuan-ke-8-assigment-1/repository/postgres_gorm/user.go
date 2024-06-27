package postgres_gorm

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
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		log.Printf("Error creating user: %v\n", err)
		return entity.User{}, err
	}
	return *user, nil
}

// GetUserByID mengambil pengguna berdasarkan ID
func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.UserDetail, error) {
	fmt.Print("masuk ke GetUserByID gorm xxxx - wow \n ")
	//var user entity.User
	var userDetail entity.UserDetail
	if err := r.db.WithContext(ctx).Table("users").Select("users.id", "name", "email", "submissions.risk_score", "submissions.risk_category", "users.created_at", "users.updated_at").
		Joins("join submissions on submissions.user_id = users.id").
		Where("users.id = ?", id).
		Order("submissions.id DESC").
		Limit(1).
		Find(&userDetail).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.UserDetail{}, nil
		}
		log.Printf("Error getting user by ID: %v\n", err)
		return entity.UserDetail{}, err
	}

	var riskCategory string
	var definition string
	riskCategory = userDetail.RiskCategory

	if riskCategory == "Conservative" {
		definition = "Tujuan utama Anda adalah untuk melindungi modal/dana yang ditempatkan dan Anda tidak memiliki toleransi " +
			"sama sekali terhadap perubahan harga/nilai dari dana investasinya tersebut. " +
			"Anda memiliki pengalaman yang sangat terbatas atau tidak memiliki pengalaman sama sekali mengenai produk investasi."
	}

	if riskCategory == "Moderate" {
		definition = "Anda memiliki toleransi yang rendah dengan perubahan harga/nilai dari dana investasi dan risiko investasi."
	}

	if riskCategory == "Balanced" {
		definition = "Anda memiliki toleransi yang cukup terhadap produk investasi dan dapat menerima perubahan yang besar dari " +
			"harga/nilai dari harga yang diinvestasikan."
	}

	if riskCategory == "Growth" {
		definition = "Anda memiliki toleransi yang cukup tinggi dan dapat menerima perubahan yang besar dari harga/nilai portfolio" +
			"pada produk investasi yang diinvestasikan." +
			"Pada umumnya Anda sudah pernah atau berpengalaman dalam berinvestasi di produk investasi."
	}

	if riskCategory == "Aggresive" {
		definition = "Anda sangat berpengalaman terhadap produk investasi dan memiliki toleransi yang sangat tinggi atas" +
			"produk-produk investasi. Anda bahkan dapat menerima perubahan signifikan pada modal/nilai investasi." +
			"Pada umumnya portfolio Anda sebagian besar dialokasikan pada produk investasi."
	}

	userDetail.RiskDefinition = definition

	return userDetail, nil
}

// UpdateUser memperbarui informasi pengguna dalam basis data
func (r *userRepository) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	// Menemukan pengguna yang akan diperbarui
	var existingUser entity.User
	if err := r.db.WithContext(ctx).Select("id", "name", "email", "created_at", "updated_at").First(&existingUser, id).Error; err != nil {
		log.Printf("Error finding user to update: %v\n", err)
		return entity.User{}, err
	}

	// Memperbarui informasi pengguna
	existingUser.Name = user.Name
	existingUser.Email = user.Email
	if err := r.db.WithContext(ctx).Save(&existingUser).Error; err != nil {
		log.Printf("Error updating user: %v\n", err)
		return entity.User{}, err
	}
	return existingUser, nil
}

// DeleteUser menghapus pengguna berdasarkan ID
func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		log.Printf("Error deleting user: %v\n", err)
		return err
	}
	return nil
}

// GetAllUsers mengambil semua pengguna dari basis data
func (r *userRepository) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	if err := r.db.WithContext(ctx).Select("id", "name", "email", "created_at", "updated_at").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, nil
		}
		log.Printf("Error getting all users: %v\n", err)
		return nil, err
	}
	return users, nil
}
