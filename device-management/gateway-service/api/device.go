package api

import (
	"gateway-service/auth"
	"gateway-service/handlers"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v5/web"
)

func RegisterDeviceApi(service web.Service, t opentracing.Tracer) {
	//POST
	service.Handle("/device/register", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterDeviceHandler(t, w, r)
	})))
	//POST
	service.Handle("/device/list", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.ListDeviceHandler(t, w, r)
	})))
	//POST
	service.Handle("/device/update", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateDeviceHandler(t, w, r)
	})))
	//POST
	service.Handle("/device/user", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.AddUserHandler(t, w, r)
	})))
	//POST
	service.Handle("/device/notification", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.NotificationHandler(t, w, r)
	})))
	//POST
	service.Handle("/device/telemetry", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.TelemetryHandler(t, w, r)
	})))
	//POST
	service.Handle("/device/command", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.CommandHandler(t, w, r)
	})))
	//POST
	service.Handle("/device", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetHandler(t, w, r)
	})))
}
