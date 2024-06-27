package entity

import (
	"encoding/json"
	"time"
)

type Submission struct {
	ID           int       `gorm:"primaryKey" json:"id"`                                 // ID subbmision sebagai primary key
	UserId       int       `gorm:"type:int;not null" json:"user_id" binding:"required"`  // ID pengguna
	Answers      []Answer  `gorm:"type:json;not null" json:"answers" binding:"required"` // jawaban pengguna
	RiskScore    int       ``                                                            // Menyimpan risk profile pengguna
	RiskCategory string    ``                                                            // Menyimpan kategori risk profile pengguna
	CreatedAt    time.Time `json:"created_at"`                                           // Waktu pembuatan pengguna
	UpdatedAt    time.Time `json:"updated_at"`                                           // Waktu pembaruan terakhir pengguna
}

type Answer struct {
	QuestionID int    `json:"question_id"`
	Answer     string `json:"answer"`
}

type SubmissionData struct {
	ID           int             `gorm:"primaryKey" json:"id"`                                  // ID subbmision sebagai primary key
	UserId       int             `gorm:"type:int;not null" json:"user_id" binding:"required"`   // ID pengguna
	Answers      json.RawMessage `gorm:"type:jsonb;not null" json:"answers" binding:"required"` // jawaban pengguna
	RiskScore    int             `gorm:"type:int;not null" json:"risk_score" binding:"required"`
	RiskCategory string          `gorm:"type:string;" json:"risk_category"` // Menyimpan kategori risk profile pengguna
	CreatedAt    time.Time       `json:"created_at"`                        // Waktu pembuatan pengguna
	UpdatedAt    time.Time       `json:"updated_at"`                        // Waktu pembaruan terakhir pengguna
}

type ProfileRiskCategory string
