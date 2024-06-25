package entity

import (
	"encoding/json"
	"time"
)

type Submission struct {
	ID           int             `gorm:"primaryKey" json:"id"`                                 // ID subbmision sebagai primary key
	UserId       int             `gorm:"type:int;not null" json:"user_id" binding:"required"`  // ID pengguna
	Answers      json.RawMessage `gorm:"type:json;not null" json:"answers" binding:"required"` // jawaban pengguna
	RiskScore    int             ``                                                            // Menyimpan risk profile pengguna
	RiskCategory string          ``                                                            // Menyimpan kategori risk profile pengguna
	CreatedAt    time.Time       `json:"created_at"`                                           // Waktu pembuatan pengguna
	UpdatedAt    time.Time       `json:"updated_at"`                                           // Waktu pembaruan terakhir pengguna
}
