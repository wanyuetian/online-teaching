package server

import (
	"github.com/go-kratos/grpc-gateway/v2/protoc-gen-openapiv2/generator"
	v1 "online-teaching/api/teaching/v1"
	"online-teaching/internal/conf"
	"online-teaching/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, course *service.CourseService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	openAPIHandler := openapiv2.NewHandler(openapiv2.WithGeneratorOptions(generator.UseJSONNamesForFields(true), generator.EnumsAsInts(true)))
	srv.HandlePrefix("/q/", openAPIHandler)
	v1.RegisterCourseHTTPServer(srv, course)
	return srv
}
