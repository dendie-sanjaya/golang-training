package postgres_gorm

import (
	"context"
	"fmt"
	"praisindo_consumer_1/entity"

	"gorm.io/gorm"
)

type UserHandler struct {
	Db *gorm.DB
}

// SaveNotification saves a notification to the database
func (handler *UserHandler) SaveNotification(ctx context.Context, message *entity.NotificationLog) error {
	if err := handler.Db.WithContext(ctx).Create(&message).Error; err != nil {
		return fmt.Errorf("failed to save notification: %w", err)
	}
	return nil
}

func (handler *UserHandler) GetAllNotifications(ctx context.Context) ([]entity.NotificationLog, error) {
	var notifications []entity.NotificationLog
	if err := handler.Db.WithContext(ctx).Find(&notifications).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve notifications: %w", err)
	}
	return notifications, nil
}
