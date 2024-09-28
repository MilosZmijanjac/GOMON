package service

import (
	"context"
	"device-service/models"

	"go-micro.dev/v5"
	"go-micro.dev/v5/errors"
	"gorm.io/gorm"
)

type DeviceService struct{ db *gorm.DB }
type RegisterRequest struct {
	DeviceId        string
	Name            string
	Status          string
	IP              string
	FirmwareVersion string
	Location        string
	OS              string
	UserIds         []string
}
type RegisterResponse struct{}
type ListDevicesRequest struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
}

type ListDevicesResponse struct {
	Devices []models.Device `json:"devices"`
}
type NewUserRequest struct {
	DeviceId string `json:"deviceId"`
	UserId   string `json:"userId"`
}
type NewUserResponse struct{}

func (s *DeviceService) Register(ctx context.Context, req *RegisterRequest, rsp *RegisterResponse) error {
	device := models.Device{
		DeviceID:        req.DeviceId,
		OS:              req.OS,
		Status:          req.Status,
		Name:            req.Name,
		Location:        req.Location,
		UserIds:         req.UserIds,
		FirmwareVersion: req.FirmwareVersion,
		IP:              req.IP,
	}
	if err := s.db.Create(&device).Error; err != nil {
		return errors.InternalServerError("DeviceService.Register", "Error saving device: %v", err)
	}

	return nil
}
func (s *DeviceService) Update(ctx context.Context, req *RegisterRequest, rsp *RegisterResponse) error {
	device := models.Device{
		DeviceID:        req.DeviceId,
		OS:              req.OS,
		Status:          req.Status,
		Name:            req.Name,
		Location:        req.Location,
		UserIds:         req.UserIds,
		FirmwareVersion: req.FirmwareVersion,
		IP:              req.IP,
	}
	if err := s.db.Save(&device).Error; err != nil {
		return errors.InternalServerError("DeviceService.Register", "Error saving device: %v", err)
	}

	return nil
}
func (s *DeviceService) List(ctx context.Context, req *ListDevicesRequest, rsp *ListDevicesResponse) error {
	var err error
	if req.IsAdmin {
		err = s.db.Find(&rsp.Devices).Error
	} else {
		err = s.db.Where("? = ANY(user_ids)", req.Username).Find(&rsp.Devices).Error
	}

	if err != nil {
		return errors.InternalServerError("DeviceService.ListDevices", "Error fetching devices: %v", err)
	}
	return nil
}
func (s *DeviceService) NewUser(ctx context.Context, req *NewUserRequest, rsp *NewUserResponse) error {
	var device models.Device

	if err := s.db.First(&device, req.DeviceId).Error; err != nil {
		return errors.InternalServerError("DeviceService.NewUser", "Error fetching device : %v", err)
	}
	device.UserIds = append(device.UserIds, req.UserId)

	if err := s.db.Save(&device).Error; err != nil {
		return errors.InternalServerError("DeviceService.NewUser", "Error adding new user : %v", err)
	}
	return nil
}
func (s *DeviceService) Get(ctx context.Context, req *NewUserRequest, rsp *models.Device) error {

	if err := s.db.First(rsp, req.DeviceId).Error; err != nil {
		return errors.InternalServerError("DeviceService.NewUser", "Error fetching device : %v", err)
	}
	return nil
}
func RegisterDeviceServiceHandler(service *micro.Service, db *gorm.DB) {
	micro.RegisterHandler((*service).Server(), &DeviceService{db: db})
}
