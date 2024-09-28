package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
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
type SocketRequest struct {
	RequestType uint32 `json:"type"`
	Payload     string `json:"payload"`
}
type CommandRequest struct {
	DeviceID string
	Command  int
}

func CollectRealData(deviceID string) (TelemetryData, error) {

	cpuPercentages, err := cpu.Percent(0, false)
	if err != nil {
		return TelemetryData{}, err
	}
	cpuUsage := cpuPercentages[0]

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return TelemetryData{}, err
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		return TelemetryData{}, err
	}

	hostStat, err := host.SensorsTemperatures()
	if err != nil || len(hostStat) == 0 {
		return TelemetryData{}, err
	}
	temperature := hostStat[0].Temperature

	netStat, err := net.IOCounters(false)
	if err != nil {
		return TelemetryData{}, err
	}
	uptime, err := host.Uptime()
	if err != nil || len(hostStat) == 0 {
		return TelemetryData{}, err
	}
	return TelemetryData{
		DeviceID:             deviceID,
		Timestamp:            time.Now(),
		Uptime:               time.Duration(uptime),
		CPUUsage:             cpuUsage,
		MemoryUsage:          vmStat.Used,
		DiskUsage:            diskStat.Used,
		Temperature:          temperature,
		NetworkBytesSent:     netStat[0].BytesSent,
		NetworkBytesRecieved: netStat[0].BytesRecv,
	}, nil
}

func CollectDemoData(deviceID string) (TelemetryData, error) {
	return TelemetryData{
		DeviceID:             deviceID,
		Timestamp:            time.Now(),
		Uptime:               time.Duration(rand.Intn(100000)) * time.Second,
		CPUUsage:             rand.Float64()*50 + 20,
		MemoryUsage:          rand.Uint64()%8192 + 1024,
		DiskUsage:            rand.Uint64()%102400 + 51200,
		Temperature:          rand.Float64()*10 + 30,
		NetworkBytesSent:     rand.Uint64()%1000000 + 10000,
		NetworkBytesRecieved: rand.Uint64()%2000000 + 50000,
	}, nil
}
func main() {
	demo := flag.Bool("demo", false, "enable demo mode")
	flag.Parse()

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
	}

	deviceID := os.Getenv("DEVICE_ID")
	if deviceID == "" {
		fmt.Println("Device id is not set in .env file!")
		return
	}
	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		fmt.Println("Device id is not set in .env file!")
		return
	}
	conn, _, err := websocket.DefaultDialer.Dial(serverAddress+deviceID, nil)
	if err != nil {
		fmt.Println("Error while connecting to websocket server")
		conn.Close()
	}

	go func() {
		for {
			var data TelemetryData
			if *demo {
				data, err = CollectDemoData(deviceID)
			} else {
				data, err = CollectRealData(deviceID)
			}

			if err != nil {
				fmt.Println("Error collecting telemetry data:", err)
				time.Sleep(5 * time.Second)
				continue
			}

			jsonData, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				fmt.Println("Error marshalling telemetry data:", err)
				time.Sleep(5 * time.Second)
				continue
			}
			request := SocketRequest{
				RequestType: 2,
				Payload:     string(jsonData),
			}
			jsonData1, err := json.MarshalIndent(request, "", "  ")
			if err != nil {
				fmt.Println("Error marshalling telemetry data:", err)
				time.Sleep(5 * time.Second)
				continue
			}
			err = conn.WriteMessage(websocket.TextMessage, jsonData1)
			if err != nil {
				fmt.Println("Error sending message:", err)
				return
			}

			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message:", err)
				return
			}
			command := CommandRequest{}
			json.Unmarshal(message, &command)

			if *demo {
				command.Command = 2
			}
			if command.Command == 1 {
				cmd := exec.Command("shutdown", "/r", "/t", "0")
				err := cmd.Run()
				if err != nil {
					log.Fatalf("Failed to restart the computer: %v", err)
				}
			} else if command.Command == 2 {
				os.Exit(1)
			}

			fmt.Printf("Received message: %s\n", message)
		}
	}()

	select {}
}
