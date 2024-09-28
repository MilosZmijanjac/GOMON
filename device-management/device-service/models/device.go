package models

import "github.com/lib/pq"

type Device struct {
	DeviceID        string         `gorm:"primaryKey;not null"`
	Name            string         `gorm:"not null"`
	Status          string         `gorm:"not null"`
	IP              string         `gorm:"not null"`
	FirmwareVersion string         `gorm:"not null"`
	Location        string         `gorm:"not null"`
	OS              string         `gorm:"not null"`
	UserIds         pq.StringArray `gorm:"type:text[];column:user_ids"`
}
