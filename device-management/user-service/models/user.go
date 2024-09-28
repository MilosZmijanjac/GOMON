package models

type User struct {
	Username string `gorm:"primaryKey"`
	Password string `gorm:"not null"`
	Role     uint8
	IsActive bool
}
