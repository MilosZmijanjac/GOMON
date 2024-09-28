package api

import (
	"gateway-service/auth"
	"gateway-service/handlers"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v5/web"
)

func RegisterUserApi(service web.Service, t opentracing.Tracer) {
	//POST
	service.Handle("/user/register", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterUserHandler(t, w, r)
	})))
	//GET
	service.Handle("/user/list", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.ListUserHandler(t, w, r)
	})))
	//POST
	service.Handle("/user/update", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateUserHandler(t, w, r)
	})))
	//POST
	service.Handle("/user/login", auth.CorsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.LoginUserHandler(t, w, r)
	})))
}
