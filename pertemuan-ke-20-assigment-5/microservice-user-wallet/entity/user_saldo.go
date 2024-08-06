package entity

import "time"

type UserSaldo struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	UserId    int       `gorm:"type:int;not null" json:"name" binding:"required"`
	WalletId  int       `gorm:"type:int;not null" json:"wallet_id" binding:"required"`
	Saldo     float32   `gorm:"type:float"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
