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
	"github.com/houseme/url-shortenter/app/console/internal/controller/user"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// Main command.
var Main = gcmd.Command{
	Name:  "main",
	Usage: "main",
	Brief: "start HTTP server",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		s := g.Server()
		s.SetRewrite("/favicon.ico", "/resource/image/favicon.ico")
		s.Group("/api.v1", func(group *ghttp.RouterGroup) {
			group.Middleware(ghttp.MiddlewareCORS, service.Middleware().Initializer, service.Middleware().Logger, service.Middleware().HandlerResponse)
			group.Group("/console", func(group *ghttp.RouterGroup) {
				group.Bind(
					auth.NewV1(),
				)
			})

			group.Group("/console", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().AuthorizationForAPI)
				group.Bind(
					echo.NewV1(),
					account.NewV1(),
					short.NewV1(),
				)
			})
			group.Group("/console", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Middleware().AuthorizationForConsole)
				group.Bind(
					home.NewV1(),
					domain.NewV1(),
					user.NewV1(),
				)
			})
		})
		s.Run()
		return nil
	},
}
