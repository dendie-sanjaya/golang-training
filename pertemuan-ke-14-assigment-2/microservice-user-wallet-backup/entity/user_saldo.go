package entity

import "time"

type UserSaldo struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	UserId    int       `gorm:"type:int;not null" json:"name" binding:"required"`
	Saldo     float64   `gorm:"type:float"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
