package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/houseme/url-shortenter/app/console/internal/controller"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

var (
	// Main command.
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")
			s.Group("/api.v1", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().ConsoleLogger, service.Middleware().Logger, service.Middleware().HandlerResponse)
				group.Group("/console", func(group *ghttp.RouterGroup) {
					group.Bind(
						controller.Auth,
					)
				})

				group.Group("/console", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().AuthorizationForAPI)
					group.Bind(
						controller.Echo,
						controller.Account,
						controller.Short,
					)
				})
				group.Group("/console", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().AuthorizationForConsole)
					group.Bind(
						// controller.Echo,
						// controller.Account,
						// controller.Short,
						controller.Domain,
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
