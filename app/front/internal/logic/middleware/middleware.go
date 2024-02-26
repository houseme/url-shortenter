// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

// Package middleware is a middleware package for middleware.
package middleware

import (
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/front/api/home/v1"
	"github.com/houseme/url-shortenter/app/front/internal/consts"
	"github.com/houseme/url-shortenter/app/front/internal/service"
	"github.com/houseme/url-shortenter/utility/helper"
)

type sMiddleware struct {
}

func init() {
	service.RegisterMiddleware(&sMiddleware{})
}

// Initializer is a middleware handler for ghttp.Request.
func (s *sMiddleware) Initializer(r *ghttp.Request) {
	r.SetCtxVar("logger", consts.DefaultLogger)
	r.Middleware.Next()
}

// ClientIP sets the client ip to the context.
func (s *sMiddleware) ClientIP(r *ghttp.Request) {
	r.SetParam("clientIP", r.GetClientIp())
	r.SetParam("userAgent", r.UserAgent())
	r.SetParam("header", r.Header)
	r.Middleware.Next()
}

// Logger Middleware Log
func (s *sMiddleware) Logger(r *ghttp.Request) {
	r.Middleware.Next()
	errStr := "success"
	if err := r.GetError(); err != nil {
		errStr = err.Error()
		g.Log(r.GetCtxVar("logger").String()).Errorf(r.GetCtx(), "Server logger Error:%+v", err)
	}
	g.Log(r.GetCtxVar("logger").String()).Debug(r.GetCtx(), "status: ", r.Response.Status, "path: ", r.URL.Path, "msg: ", errStr)
}

// MiddlewareHandlerResponse 响应处理
func (s *sMiddleware) MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	ctx, span := gtrace.NewSpan(r.GetCtx(), "tracing-service-new-MiddlewareHandlerResponse")
	r.SetCtx(ctx)
	defer span.End()
	// There's custom buffer content, it then exits the current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err    = r.GetError()
		res    = r.GetHandlerResponse()
		logger = g.Log(helper.Helper().Logger(ctx))
	)
	logger.Info(ctx, "MiddlewareHandlerResponse response:", res, " statusCode:", r.Response.Status)
	if g.IsNil(res) || g.IsEmpty(res) {
		r.Response.Status = http.StatusNotFound
	}
	if err != nil {
		logger.Errorf(ctx, "MiddlewareHandlerResponse err:%+v", err)
		r.Response.Status = http.StatusInternalServerError
		if internalErr := r.Response.WriteTpl("error.html", g.Map{
			"title":   "内部错误 - 短链平台",
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"label":   "Error",
		}); internalErr != nil {
			logger.Errorf(ctx, `r.Response.WriteTpl internalErr %+v`, internalErr)
		}
	}
	if r.Response.Status > 0 && r.Response.Status != http.StatusOK && r.Response.Status != http.StatusFound {
		if internalErr := r.Response.WriteTpl("error.html", g.Map{
			"title":   "404 - 短链平台",
			"code":    r.Response.Status,
			"message": "您访问的页面已失效",
			"label":   http.StatusText(r.Response.Status),
		}); internalErr != nil {
			logger.Errorf(ctx, `r.Response.WriteTpl 404 err: %+v`, internalErr)
		}
	}

	logger.Debug(r.GetCtx(), "MiddlewareHandlerResponse end")
	if !g.IsNil(res) && !g.IsEmpty(res) {
		logger.Debug(r.GetCtx(), "MiddlewareHandlerResponse redirect url:", res)
		str := res.(*v1.HomeRes)
		r.Response.RedirectTo(string(*str), http.StatusFound)
	}
}

// MiddlewareHandlerRequest 请求处理
func (s *sMiddleware) MiddlewareHandlerRequest(r *ghttp.Request) {
	ctx, span := gtrace.NewSpan(r.GetCtx(), "tracing-service-new-MiddlewareHandlerRequest")
	r.SetCtx(ctx)
	defer span.End()

	g.Log().Debug(ctx, "MiddlewareHandlerRequest start")
	r.SetParam("rawQuery", r.Request.URL.RawQuery)
	r.SetParam("shortAll", r.Request.URL.String())
	r.SetParam("referer", r.Referer())
	r.SetParam("path", r.URL.Path)
	r.SetParam("origin", r.GetHeader("Origin"))
	r.SetParam("host", r.Request.Host)
	r.Middleware.Next()
}
