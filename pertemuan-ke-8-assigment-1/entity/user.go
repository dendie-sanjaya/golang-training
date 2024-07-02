package entity

import "time"

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`                                        // ID pengguna sebagai primary key
	Name      string    `gorm:"type:varchar;not null" json:"name" binding:"required"`        // Nama pengguna (wajib diisi)
	Email     string    `gorm:"type:varchar;not null" json:"email" binding:"required,email"` // Email pengguna (wajib diisi)
	CreatedAt time.Time `json:"created_at"`                                                  // Waktu pembuatan pengguna
	UpdatedAt time.Time `json:"updated_at"`                                                  // Waktu pembaruan terakhir pengguna
}

type UserDetail struct {
	ID             int       `gorm:"primaryKey" json:"id"`                                  // ID pengguna sebagai primary key
	Name           string    `gorm:"type:varchar;not null" json:"name" binding:"required"`  // Nama pengguna (wajib diisi)
	Email          string    `gorm:"type:varchar;not null" json:"email" binding:"required"` // Email pengguna (wajib diisi)
	RiskScore      int       `json:"risk_score"`                                            // Skor risiko pengguna
	RiskCategory   string    `json:"risk_category"`                                         // Kategori risiko pengguna
	RiskDefinition string    `json:"risk_definition"`                                       // Definisi risiko pengguna
	CreatedAt      time.Time `json:"created_at"`                                            // Waktu pembuatan pengguna
	UpdatedAt      time.Time `json:"updated_at"`                                            // Waktu pembaruan terakhir pengguna
}
