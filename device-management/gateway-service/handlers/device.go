package handlers

import (
	"context"
	"encoding/json"
	"gateway-service/models"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	ocplugin "github.com/micro/plugins/v5/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v5/client"
)

func RegisterDeviceHandler(t opentracing.Tracer, w http.ResponseWriter, r *http.Request) {
	log.Println(r.Body)
	span := t.StartSpan(r.RequestURI)
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	c := client.NewClient(
		client.WrapCall(ocplugin.NewCallWrapper(t)),
	)
	var reqData models.RegisterDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if span != nil {
		span.LogKV("request", reqData)
	}
	req := c.NewRequest("device.service", "DeviceService.Register", reqData, client.WithContentType("application/json"))
	rsp := &models.RegisterDeviceResponse{}

	if err := c.Call(ctx, req, rsp); err != nil {
		log.Printf("Error calling Device service: %v", err)
		http.Error(w, "Error calling remote service", http.StatusInternalServerError)
		return
	}
}
func ListDeviceHandler(t opentracing.Tracer, w http.ResponseWriter, r *http.Request) {
	span := t.StartSpan(r.RequestURI)
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	c := client.NewClient(
		client.WrapCall(ocplugin.NewCallWrapper(t)),
	)
	var reqData models.ListDevicesRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if span != nil {
		span.LogKV("request", reqData)
	}
	req := c.NewRequest("device.service", "DeviceService.List", reqData, client.WithContentType("application/json"))
	rsp := &models.ListDevicesResponse{}

	if err := c.Call(ctx, req, rsp); err != nil {
		log.Printf("Error calling Device service: %v", err)
		http.Error(w, "Error calling remote service", http.StatusInternalServerError)
		return
	}
	if span != nil {
		span.LogKV("response", rsp)
	}
	byteData, err := json.Marshal(rsp)
	if err != nil {
		log.Print(err)
		return
	}
	w.Write(byteData)
}
func UpdateDeviceHandler(t opentracing.Tracer, w http.ResponseWriter, r *http.Request) {
	span := t.StartSpan(r.RequestURI)
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	c := client.NewClient(
		client.WrapCall(ocplugin.NewCallWrapper(t)),
	)
	var reqData models.RegisterDeviceRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if span != nil {
		span.LogKV("request", reqData)
	}
	req := c.NewRequest("device.service", "DeviceService.Update", reqData, client.WithContentType("application/json"))
	rsp := &models.RegisterDeviceResponse{}

	if err := c.Call(ctx, req, rsp); err != nil {
		log.Printf("Error calling Device service: %v", err)
		http.Error(w, "Error calling remote service", http.StatusInternalServerError)
		return
	}
}
func AddUserHandler(t opentracing.Tracer, w http.ResponseWriter, r *http.Request) {
	span := t.StartSpan(r.RequestURI)
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	c := client.NewClient(
		client.WrapCall(ocplugin.NewCallWrapper(t)),
	)
	var reqData models.NewUserRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if span != nil {
		span.LogKV("request", reqData)
	}
	req := c.NewRequest("device.service", "DeviceService.NewUser", reqData, client.WithContentType("application/json"))
	rsp := &models.NewUserResponse{}

	if err := c.Call(ctx, req, rsp); err != nil {
		log.Printf("Error calling Device service: %v", err)
		http.Error(w, "Error calling remote service", http.StatusInternalServerError)
		return
	}
}
func NotificationHandler(t opentracing.Tracer, w http.ResponseWriter, r *http.Request) {
	span := t.StartSpan(r.RequestURI)
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	c := client.NewClient(
		client.WrapCall(ocplugin.NewCallWrapper(t)),
	)
	var reqData models.NotificationRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if span != nil {
		span.LogKV("request", reqData)
	}
	req := c.NewRequest("notification.service", "NotificationService.Check", reqData, client.WithContentType("application/json"))
	rsp := &models.NotificationResponse{}

	if err := c.Call(ctx, req, rsp); err != nil {
		log.Printf("Error calling Notification service: %v", err)
		http.Error(w, "Error calling remote service", http.StatusInternalServerError)
		return
	}
	if span != nil {
		span.LogKV("response", rsp)
	}
	byteData, err := json.Marshal(rsp)
	if err != nil {
		log.Print(err)
		return
	}
	w.Write(byteData)
}
func TelemetryHandler(t opentracing.Tracer, w http.ResponseWriter, r *http.Request) {
	span := t.StartSpan(r.RequestURI)
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	c := client.NewClient(
		client.WrapCall(ocplugin.NewCallWrapper(t)),
	)

	var reqData models.TelemetryRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if span != nil {
		span.LogKV("request", reqData)
	}
	req := c.NewRequest("telemetry.service", "TelemetryService.Get", reqData, client.WithContentType("application/json"))
	rsp := &models.TelemetryResponse{}

	if err := c.Call(ctx, req, rsp); err != nil {
		log.Printf("Error calling Telemetry service: %v", err)
		http.Error(w, "Error calling remote service", http.StatusInternalServerError)
		return
	}
	if span != nil {
		span.LogKV("response", rsp)
	}
	byteData, err := json.Marshal(rsp)
	if err != nil {
		log.Print(err)
		return
	}
	w.Write(byteData)
}
func CommandHandler(t opentracing.Tracer, w http.ResponseWriter, r *http.Request) {
	span := t.StartSpan(r.RequestURI)
	defer span.Finish()

	opentracing.ContextWithSpan(context.Background(), span)
	var reqData models.CommandRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if span != nil {
		span.LogKV("request", reqData)
	}
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Error loading .env file: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		log.Println("SERVER_ADDRESS is not set in .env file!")
		http.Error(w, "Server address not configured", http.StatusInternalServerError)
		return
	}

	conn, _, err := websocket.DefaultDialer.Dial(serverAddress+"gateway-service", nil)
	if err != nil {
		log.Printf("Error connecting to WebSocket server: %v\n", err)
		http.Error(w, "Error connecting to WebSocket server", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	jsonData, err := json.MarshalIndent(reqData, "", "  ")
	if err != nil {
		log.Printf("Error marshaling request data: %v\n", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	socketRequest := models.SocketRequest{
		RequestType: 1,
		Payload:     string(jsonData),
	}

	jsonData1, err := json.MarshalIndent(socketRequest, "", "  ")
	if err != nil {
		log.Printf("Error marshaling socket request: %v\n", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	if err := conn.WriteMessage(websocket.TextMessage, jsonData1); err != nil {
		log.Printf("Error sending message to WebSocket server: %v\n", err)
		http.Error(w, "Error sending message to server", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Command successfully sent"))
}
func GetHandler(t opentracing.Tracer, w http.ResponseWriter, r *http.Request) {
	span := t.StartSpan(r.RequestURI)
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	c := client.NewClient(
		client.WrapCall(ocplugin.NewCallWrapper(t)),
	)
	deviceId := r.URL.Query().Get("id")
	req := c.NewRequest("device.service", "DeviceService.NewUser", models.GetDeviceRequest{DeviceID: deviceId}, client.WithContentType("application/json"))
	rsp := &models.Device{}

	if err := c.Call(ctx, req, rsp); err != nil {
		log.Printf("Error calling Device service: %v", err)
		http.Error(w, "Error calling remote service", http.StatusInternalServerError)
		return
	}
}
