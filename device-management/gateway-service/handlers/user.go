package handlers

import (
	"context"
	"encoding/json"
	"gateway-service/models"
	"log"
	"net/http"

	ocplugin "github.com/micro/plugins/v5/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v5/client"
)

func RegisterUserHandler(t opentracing.Tracer, w http.ResponseWriter, r *http.Request) {
	span := t.StartSpan(r.RequestURI)
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	c := client.NewClient(
		client.WrapCall(ocplugin.NewCallWrapper(t)),
	)
	var reqData models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if span != nil {
		span.LogKV("request", reqData)
	}
	req := c.NewRequest("user.service", "UserService.Register", reqData, client.WithContentType("application/json"))
	rsp := &models.RegisterResponse{}

	if err := c.Call(ctx, req, rsp); err != nil {
		log.Printf("Error calling User service: %v", err)
		http.Error(w, "Error calling remote service", http.StatusInternalServerError)
		return
	}
}
func ListUserHandler(t opentracing.Tracer, w http.ResponseWriter, r *http.Request) {
	span := t.StartSpan(r.RequestURI)
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	c := client.NewClient(
		client.WrapCall(ocplugin.NewCallWrapper(t)),
	)
	req := c.NewRequest("user.service", "UserService.List", &models.ListRequest{}, client.WithContentType("application/json"))
	rsp := &models.ListResponse{}

	if err := c.Call(ctx, req, rsp); err != nil {
		log.Printf("Error calling User service: %v", err)
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
func UpdateUserHandler(t opentracing.Tracer, w http.ResponseWriter, r *http.Request) {
	span := t.StartSpan(r.RequestURI)
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	var reqData models.UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if span != nil {
		span.LogKV("request", reqData)
	}
	c := client.NewClient(
		client.WrapCall(ocplugin.NewCallWrapper(t)),
	)
	req := c.NewRequest("user.service", "UserService.Update", reqData, client.WithContentType("application/json"))
	rsp := &models.UpdateResponse{}

	if err := c.Call(ctx, req, rsp); err != nil {
		log.Printf("Error calling User service: %v", err)
		http.Error(w, "Error calling remote service", http.StatusInternalServerError)
		return
	}
	if span != nil {
		span.LogKV("response", rsp)
	}
}
func LoginUserHandler(t opentracing.Tracer, w http.ResponseWriter, r *http.Request) {
	span := t.StartSpan(r.RequestURI)
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	var reqData models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}
	if span != nil {
		span.LogKV("request", reqData)
	}
	c := client.NewClient(
		client.WrapCall(ocplugin.NewCallWrapper(t)),
	)
	req := c.NewRequest("user.service", "UserService.Login", reqData, client.WithContentType("application/json"))
	rsp := &models.LoginResponse{}

	if err := c.Call(ctx, req, rsp); err != nil {
		http.Error(w, "Invalid Username/Password", http.StatusUnauthorized)
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
