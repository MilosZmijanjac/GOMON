package main

import (
	"fmt"
	"gateway-service/api"
	"log"
	"net/http"

	"github.com/micro/plugins/v5/registry/consul"
	"github.com/opentracing/opentracing-go"

	"github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
	"go-micro.dev/v5/web"
)

type RegisterRequest struct {
	Username string
	IsAdmin  bool
}
type LoginResponse struct {
	Devices []Device
}
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

func main() {
	cfg := &config.Configuration{
		ServiceName: "api-gateway",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	defer closer.Close()
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)

	r := consul.NewRegistry()
	service := web.NewService(
		web.Registry(r),
		web.Name("gateway.service"),
		web.Address(":8000"),
	)

	service.Init()

	api.RegisterUserApi(service, tracer)
	api.RegisterDeviceApi(service, tracer)
	service.Handle("/", http.DefaultServeMux)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
