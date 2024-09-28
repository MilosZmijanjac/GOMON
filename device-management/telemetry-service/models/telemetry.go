package models

import (
	"time"
)

type TelemetryData struct {
	DeviceID             string        `json:"deviceId"`
	Timestamp            time.Time     `json:"timestamp"`
	Uptime               time.Duration `json:"uptime"`
	CPUUsage             float64       `json:"cpuUsage"`
	MemoryUsage          uint64        `json:"memoryUsage"`
	DiskUsage            uint64        `json:"diskUsage"`
	Temperature          float64       `json:"temperature"`
	NetworkBytesSent     uint64        `json:"networkBytesSent"`
	NetworkBytesRecieved uint64        `json:"networkBytesRecieved"`
}
