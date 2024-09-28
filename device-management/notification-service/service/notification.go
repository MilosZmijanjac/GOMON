package service

import (
	"context"
	"notification-service/models"

	"go-micro.dev/v5"
	"go-micro.dev/v5/errors"
	"gorm.io/gorm"
)

type NotificationService struct{ db *gorm.DB }
type NotificationRequest struct {
	DeviceID string
}
type NotificationResponse struct {
	Notifications []models.Notification
}

func (n *NotificationService) Check(ctx context.Context, req *NotificationRequest, rsp *NotificationResponse) error {

	if err := n.db.Table("notifications").Where("device_id = ?", req.DeviceID).Order("timestamp desc").Limit(5).Find(&rsp.Notifications).Error; err != nil {
		return errors.InternalServerError("NotificationService.Check", "Error saving user: %v", err)
	}
	return nil
}
func RegisterNotificationServiceHandler(service *micro.Service, db *gorm.DB) {
	micro.RegisterHandler((*service).Server(), &NotificationService{db: db})
}
