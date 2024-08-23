package entity

import "time"

type Notification struct {
	From    User   `json:"from"`
	To      User   `json:"to"`
	Message string `json:"message"`
}

type NotificationLog struct {
	FromId    string    `gorm:"type:string;not null"`
	ToId      string    `gorm:"type:string;not null"`
	Message   string    `gorm:"type:string;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
