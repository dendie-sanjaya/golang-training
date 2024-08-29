package entity

import "time"

type UserWalet struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	UserId    int       `gorm:"type:int;not null" json:"name" binding:"required"`
	Type      string    `gorm:"type:string;not null" json:"type" binding:"required"`
	Name      string    `gorm:"type:string;not null" json:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
