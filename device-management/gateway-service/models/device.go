package models

import "time"

type Device struct {
	DeviceID        string
	Name            string
	Status          string
	IP              string
	FirmwareVersion string
	Location        string
	OS              string
	UserIds         []string
}
type RegisterDeviceRequest struct {
	DeviceId        string   `json:"deviceId"`
	Name            string   `json:"name"`
	Status          string   `json:"status"`
	IP              string   `json:"ip"`
	FirmwareVersion string   `json:"firmwareVersion"`
	Location        string   `json:"location"`
	OS              string   `json:"os"`
	UserIds         []string `json:"userIds"`
}
type RegisterDeviceResponse struct{}
type ListDevicesRequest struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
}
type ListDevicesResponse struct {
	Devices []Device `json:"devices"`
}
type NewUserRequest struct {
	DeviceId string `json:"deviceId"`
	UserId   string `json:"userId"`
}
type NewUserResponse struct{}
type Notification struct {
	DeviceID  string    `json:"deviceId"`
	Timestamp time.Time `json:"timestamp"`
	Code      int32     `json:"code"`
}
type NotificationRequest struct {
	DeviceID string
}
type NotificationResponse struct {
	Notifications []Notification
}
type TelemetryRequest struct {
	DeviceID string
}
type TelemetryResponse struct {
	LastUptimeSeconds    float64     `json:"lastUptimeSeconds"`
	CPUUsages            []float64   `json:"cpuUsages"`
	MemoryUsages         []uint64    `json:"memoryUsages"`
	DiskUsages           []uint64    `json:"diskUsages"`
	Temperatures         []float64   `json:"temperatures"`
	NetworkBytesSent     []uint64    `json:"networkBytesSent"`
	NetworkBytesReceived []uint64    `json:"networkBytesReceived"`
	Timestamps           []time.Time `json:"timestamps"`
}
type CommandRequest struct {
	DeviceID string
	Command  uint32
}
type SocketRequest struct {
	RequestType uint32 `json:"type"`
	Payload     string `json:"payload"`
}
type GetDeviceRequest struct {
	DeviceID string
}
