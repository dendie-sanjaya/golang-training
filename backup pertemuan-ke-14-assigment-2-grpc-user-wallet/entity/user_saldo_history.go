package entity

import "time"

type UserSaldoHistory struct {
	ID              int       `gorm:"primaryKey" json:"id"`
	UserIdFrom      int       `gorm:"type:int;not null" json:"user_id_from" binding:"required"`
	UserIdTo        int       `gorm:"type:int;not null" json:"user_id_to" binding:"required"`
	TypeTransaction string    `gorm:"type:string;not null" json:"type_transcation" binding:"required"`
	TypeCredit      string    `gorm:"type:string;not null" json:"type_credit" binding:"required"`
	Total           float64   `gorm:"type:float" json:"total" binding:"required`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
