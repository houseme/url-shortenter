package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/houseme/url-shortenter/app/api/internal/controller"
	"github.com/houseme/url-shortenter/app/api/internal/service"
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
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().MiddlewareHandlerResponse)
				group.Bind(
					controller.Home,
				)
			})
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,
				)
			})
			s.Run()
			return nil
		},
	}
)
