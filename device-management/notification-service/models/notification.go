package models

import "time"

type Notification struct {
	DeviceID  string    `json:"deviceId"`
	Timestamp time.Time `json:"timestamp"`
	Code      int32     `json:"code"`
}
