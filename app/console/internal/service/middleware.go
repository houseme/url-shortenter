// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type IMiddleware interface {
	ConsoleLogger(r *ghttp.Request)
	Logger(r *ghttp.Request)
	HandlerResponse(r *ghttp.Request)
	AuthorizationForAPI(r *ghttp.Request)
	AuthorizationForConsole(r *ghttp.Request)
}

var localMiddleware IMiddleware

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
