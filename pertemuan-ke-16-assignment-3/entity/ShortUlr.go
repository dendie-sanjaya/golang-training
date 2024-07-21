package entity

import "time"

type ShortUrl struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	UrlLong   string    `gorm:"type:varchar;not null" json:"name" binding:"required"`
	UrlShort  string    `gorm:"type:varchar;uniqueIndex;not null" binding:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
