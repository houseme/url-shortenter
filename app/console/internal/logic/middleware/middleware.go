// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/houseme/url-shortenter/app/console/internal/consts"
	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
	"github.com/houseme/url-shortenter/internal/tracing"
	"github.com/houseme/url-shortenter/utility/cache"
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

// Logger Middleware Log
func (s *sMiddleware) Logger(r *ghttp.Request) {
	r.Middleware.Next()
	errStr := "success"
	if err := r.GetError(); err != nil {
		g.Log(r.GetCtxVar("logger").String()).Errorf(r.GetCtx(), "Server logger Error:%+v", err)
		errStr = err.Error()
	}
	g.Log(r.GetCtxVar("logger").String()).Debug(r.GetCtx(), "status: ", r.Response.Status, "path: ", r.URL.Path, "msg: ", errStr)
}

// HandlerResponse is a middleware handler for ghttp.Request.
func (s *sMiddleware) HandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	ctx, span := gtrace.NewSpan(r.GetCtx(), "tracing-service-middleware-HandlerResponse")
	r.SetCtx(ctx)
	defer span.End()

	// There's custom buffer content, it then exits the current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	// 设置公共参数
	tracing.SetAttributes(r, span)

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
		msg = http.StatusText(r.Response.Status)
		switch r.Response.Status {
		case http.StatusNotFound:
			code = gcode.CodeNotFound
		case http.StatusForbidden:
			code = gcode.CodeNotAuthorized
		case http.StatusBadGateway:
			code = gcode.CodeInvalidRequest
		case http.StatusInternalServerError:
			code = gcode.CodeNotSupported
		default:
			code = gcode.CodeUnknown
		}
		err = gerror.NewCode(code, msg)
		r.SetError(err)
	} else {
		code = gcode.New(http.StatusOK, "success", nil)
		msg = code.Message()
		g.Log(r.GetCtxVar("logger").String()).Debug(r.GetCtx(), "HandlerResponse body res:", res)
	}
	r.Response.WriteJson(&model.DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
		Time:    gtime.TimestampMicro(),
		TraceID: span.SpanContext().TraceID().String(),
	})
}

// AuthorizationForAPI is a middleware handler for ghttp.Request.
func (s *sMiddleware) AuthorizationForAPI(r *ghttp.Request) {
	if s.authorization(r, consts.AuthTypeAPIKey) {
		r.Middleware.Next()
	}
}

// AuthorizationForConsole is a middleware handler for ghttp.Request.
func (s *sMiddleware) AuthorizationForConsole(r *ghttp.Request) {
	if s.authorization(r, consts.AuthTypePassword) {
		r.Middleware.Next()
	}
}

// authorization is a middleware handler for ghttp.Request.
func (s *sMiddleware) authorization(r *ghttp.Request, authType string) bool {
	ctx, span := gtrace.NewSpan(r.Context(), "tracing-console-service-middleware-authorization")
	r.SetCtx(ctx)
	defer span.End()

	var (
		authHeader = gstr.Trim(r.GetHeader(consts.AuthorizationHeaderKey))
		logger     = g.Log(r.GetCtxVar("logger").String())
		resp       = &model.DefaultHandlerResponse{
			Code:    http.StatusMovedPermanently,
			Message: http.StatusText(http.StatusMovedPermanently),
			Data:    nil,
			Time:    gtime.TimestampMicro(),
			TraceID: span.SpanContext().TraceID().String(),
		}
	)
	logger.Debug(r.GetCtx(), "authorization authHeader: ", authHeader)
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		resp.Message = "Invalid authorization Header"
		return s.middlewareResponse(r, span, resp)
	}

	if fields[0] != consts.AuthorizationTypeBearer {
		resp.Message = "Unsupported authorization Type"
		return s.middlewareResponse(r, span, resp)
	}

	var res, err = validateToken(r.GetCtx(), fields[1], authType, logger)
	if err != nil {
		logger.Error(r.GetCtx(), "authorization failed: ", err)
		resp.Message = "authorization failed reason: " + err.Error()
		return s.middlewareResponse(r, span, resp)
	}

	if res == nil {
		logger.Debug(r.GetCtx(), "authorization failed")
		resp.Message = "authorization failed"
		return s.middlewareResponse(r, span, resp)
	}

	if res.AuthToken != fields[1] {
		logger.Debug(r.GetCtx(), "authorization token is Refresh new token:", res.AuthToken)
		resp.Code = http.StatusFound
		resp.Message = "token is Refresh"
		resp.Data = g.Map{
			"token": res.AuthToken,
		}
		return s.middlewareResponse(r, span, resp)
	}

	r.SetParam("authAccountNo", res.AuthAccountNo)
	r.SetParam("authAccountLevel", res.AuthAccountLevel)

	r.SetCtxVar("authAccountNo", res.AuthAccountNo)
	r.SetCtxVar("authAccountLevel", res.AuthAccountLevel)
	return true
}

// validateToken is a middleware handler for ghttp.Request.
func validateToken(ctx context.Context, token, authType string, logger glog.ILogger) (authToken *model.AuthorizationToken, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-console-service-middleware-validateToken")
	defer span.End()

	defer func() {
		if err != nil {
			logger.Errorf(ctx, "validateToken failed error:%+v", err)
		}
		logger.Debug(ctx, "validateToken end")
	}()
	var conn gredis.Conn
	if conn, err = g.Redis(cache.RedisCache().ShortAccessTokenConn(ctx)).Conn(ctx); err != nil {
		err = gerror.Wrap(err, "validateToken Redis conn failed")
		return nil, err
	}
	defer func() {
		_ = conn.Close(ctx)
	}()
	var (
		redisKey       = cache.RedisCache().ShortAccessTokenKey(ctx, token)
		isAuthPassword = false
	)
	if authType == consts.AuthTypePassword {
		isAuthPassword = true
		redisKey = cache.RedisCache().ShortAuthorizationKey(ctx, token)
	}
	var val *gvar.Var
	if val, err = conn.Do(ctx, "GET", redisKey); err != nil {
		err = gerror.Wrap(err, "validateToken Redis get failed(001)")
		return nil, err
	}

	if val.IsNil() || val.IsEmpty() {
		err = gerror.New("validateToken auth token not found")
		return nil, err
	}

	if err = val.Scan(&authToken); err != nil {
		err = gerror.Wrap(err, "validateToken Redis scan failed")
		return nil, err
	}

	if authToken == nil {
		err = gerror.New("validateToken Redis get failed(002)")
		return nil, err
	}

	if authToken.AuthType != authType {
		err = gerror.New("validateToken auth token type not match")
		return nil, err
	}
	var (
		authTime = authToken.AuthTime
		now      = gtime.Now()
	)
	// 刷新 token 过期时间
	if isAuthPassword {
		if now.Unix()-consts.RefreshTokenExpireTime > authTime {
			err = gerror.New("validateToken auth token expired")
			return nil, err
		}

		authToken.AuthTime = now.Unix()
		if now.Unix()-consts.PasswordExpireTime > authTime {
			logger.Debug(ctx, "validateToken auth token password expired 2 hours")
			if token, err = helper.Helper().CreateAccessToken(ctx, authToken.AuthAccountNo); err != nil {
				err = gerror.Wrap(err, "validateToken CreateAccessToken failed")
				return nil, err
			}
			authToken.AuthToken = token
			redisKey = cache.RedisCache().ShortAuthorizationKey(ctx, token)
		}
		logger.Debug(ctx, "validateToken auth token authTime:", authTime, "now:", now.Unix(), " authToken:", authToken)
		if val, err = conn.Do(ctx, "SETEX", redisKey, consts.TokenExpireTime, authToken); err != nil {
			err = gerror.Wrap(err, "validateToken Redis set failed")
			return nil, err
		}
		logger.Debug(ctx, "validateToken auth token set Redis value:", val)
		return authToken, nil
	}
	if now.Unix()-consts.APIKeyExpireTime > authTime {
		err = gerror.New("validateToken auth token expired")
		return nil, err
	}

	return authToken, nil
}

// middlewareResponse intercept the response
func (s *sMiddleware) middlewareResponse(r *ghttp.Request, span *gtrace.Span, resp *model.DefaultHandlerResponse) bool {
	g.Log(r.GetCtxVar("logger").String()).Debug(r.GetCtx(), "middlewareResponse body resp:", resp)
	// 设置公共参数
	tracing.SetAttributes(r, span)
	r.Response.WriteJson(resp)
	return false
}
