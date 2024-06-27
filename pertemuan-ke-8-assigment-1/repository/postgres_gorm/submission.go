package postgres_gorm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"praisindo/entity"
	"praisindo/service"

	"gorm.io/gorm"
)

type submissionRespository struct {
	db GormDBIface
}

// NewsubmissionRespository membuat instance baru dari submissionRespository
func NewsubmissionRespository(db GormDBIface) service.ISubmissionRepository {
	return &submissionRespository{db: db}
}

// GetSubmissionsByID mengambil pengguna berdasarkan ID
func (r *submissionRespository) GetUserByID(ctx context.Context, id int) (entity.Submission, error) {
	fmt.Print("masuk ke GetUserByID gorm submission \n ")
	var user entity.Submission
	if err := r.db.WithContext(ctx).Select("id", "user_id", "answers", "risk_score", "risk_category", "created_at", "updated_at").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Submission{}, nil
		}
		log.Printf("Error getting user by ID: %v\n", err)
		return entity.Submission{}, err
	}
	return user, nil
}

// CreateUser membuat pengguna baru dalam basis data
func (r *submissionRespository) CreateSubmissions(ctx context.Context, user *entity.Submission) (entity.Submission, error) {
	fmt.Print("masuk ke CreateUser gorm submission \n ")
	answare_json, _ := json.Marshal(user.Answers)
	submissionData := &entity.SubmissionData{
		UserId:       user.UserId,
		Answers:      answare_json,
		RiskScore:    user.RiskScore,
		RiskCategory: user.RiskCategory,
	}
	if err := r.db.WithContext(ctx).Table("submissions").Create(submissionData).Error; err != nil {
		log.Printf("Error creating user: %v\n", err)
		return entity.Submission{}, err
	}
	return *user, nil

}

// DeleteSubmissions menghapus pengguna berdasarkan ID
func (r *submissionRespository) DeleteSubmissions(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Submission{}, id).Error; err != nil {
		log.Printf("Error deleting submission: %v\n", err)
		return err
	}
	return nil
}

// GetAllSubmissions mengambil semua pengguna dari basis data
func (r *submissionRespository) GetAllSubmissions(ctx context.Context) ([]entity.Submission, error) {
	fmt.Print("masuk ke GetAllSubmissions gorm submission \n ")
	var users []entity.Submission
	if err := r.db.WithContext(ctx).Select("id", "user_id", "answers", "risk_score", "risk_category", "created_at", "updated_at").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users, nil
		}
		log.Printf("Error getting all users: %v\n", err)
		return nil, err
	}
	return users, nil
}

// GetUserByID mengambil pengguna berdasarkan ID
func (r *submissionRespository) GetSubmissionsByID(ctx context.Context, id int) (entity.Submission, error) {
	fmt.Print("masuk ke GetUserByID gorm \n ")
	var user entity.SubmissionData
	if err := r.db.WithContext(ctx).Table("submissions").Select("id", "user_id", "answers", "risk_score", "risk_category", "created_at", "updated_at").
		Where("user_id = ?", id).
		Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Submission{}, err
		}
		log.Printf("Error getting user by User ID: %v\n", err)
		return entity.Submission{}, err
	}

	//answare_json_decode, _ := json.Marshal(user.Answers)
	// submission := &entity.Submission{
	// 	UserId:       user.UserId,
	// 	RiskScore:    user.RiskScore,
	// 	RiskCategory: user.RiskCategory,
	// }
	// var subbmision entity.Submission
	return entity.Submission{}, nil
}
