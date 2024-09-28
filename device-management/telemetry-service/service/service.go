package service

import (
	"log"

	"github.com/micro/plugins/v5/registry/consul"
	ocplugin "github.com/micro/plugins/v5/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v5"
)

func InitService(service *micro.Service) {
	(*service).Init()
}
func NewService(t opentracing.Tracer) *micro.Service {
	r := consul.NewRegistry()
	srv := micro.NewService(micro.Registry(r), micro.Name("telemetry.service"), micro.WrapHandler(ocplugin.NewHandlerWrapper(t)))
	return &srv
}
func RunService(service *micro.Service) {
	if err := (*service).Run(); err != nil {
		log.Fatal(err)
	}
}
