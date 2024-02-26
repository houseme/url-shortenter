// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// Initializer is a middleware handler for ghttp.Request.
		Initializer(r *ghttp.Request)
		// ClientIP sets the client ip to the context.
		ClientIP(r *ghttp.Request)
		// Logger Middleware Log
		Logger(r *ghttp.Request)
		// MiddlewareHandlerResponse 响应处理
		MiddlewareHandlerResponse(r *ghttp.Request)
		// MiddlewareHandlerRequest 请求处理
		MiddlewareHandlerRequest(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
