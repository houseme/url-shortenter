// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/houseme/url-shortenter/app/console/internal/controller/account"
	"github.com/houseme/url-shortenter/app/console/internal/controller/auth"
	"github.com/houseme/url-shortenter/app/console/internal/controller/domain"
	"github.com/houseme/url-shortenter/app/console/internal/controller/echo"
	"github.com/houseme/url-shortenter/app/console/internal/controller/home"
	"github.com/houseme/url-shortenter/app/console/internal/controller/short"
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
				group.Middleware(ghttp.MiddlewareCORS, service.Middleware().ConsoleLogger, service.Middleware().Logger, service.Middleware().HandlerResponse)
				group.Group("/console", func(group *ghttp.RouterGroup) {
					group.Bind(
						auth.New(),
					)
				})

				group.Group("/console", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().AuthorizationForAPI)
					group.Bind(
						echo.New(),
						account.New(),
						short.New(),
					)
				})
				group.Group("/console", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().AuthorizationForConsole)
					group.Bind(
						home.New(),
						// controller.Account,
						// controller.Short,
						domain.New(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
