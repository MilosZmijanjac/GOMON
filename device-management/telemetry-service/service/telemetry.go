package service

import (
	"context"
	"telemetry-service/models"
	"time"

	"go-micro.dev/v5"
	"gorm.io/gorm"
)

type TelemetryService struct{ db *gorm.DB }
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

func (t *TelemetryService) Get(ctx context.Context, req *TelemetryRequest, rsp *TelemetryResponse) error {
	var telemetryData []models.TelemetryData
	t.db.Where("device_id = ? AND timestamp >= ?", req.DeviceID, time.Now().Add(-1*time.Minute)).
		Order("timestamp asc").Find(&telemetryData)

	for _, data := range telemetryData {
		rsp.Timestamps = append(rsp.Timestamps, data.Timestamp)
		rsp.CPUUsages = append(rsp.CPUUsages, data.CPUUsage)
		rsp.MemoryUsages = append(rsp.MemoryUsages, data.MemoryUsage)
		rsp.DiskUsages = append(rsp.DiskUsages, data.DiskUsage)
		rsp.Temperatures = append(rsp.Temperatures, data.Temperature)
		rsp.NetworkBytesSent = append(rsp.NetworkBytesSent, data.NetworkBytesSent)
		rsp.NetworkBytesReceived = append(rsp.NetworkBytesReceived, data.NetworkBytesRecieved)
		rsp.LastUptimeSeconds = data.Uptime.Seconds()
	}

	return nil
}
func RegisterTelemetryServiceHandler(service *micro.Service, db *gorm.DB) {
	micro.RegisterHandler((*service).Server(), &TelemetryService{db: db})
}
