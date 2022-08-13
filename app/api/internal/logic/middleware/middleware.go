package middleware

import (
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/api/api/v1"
	"github.com/houseme/url-shortenter/app/api/internal/service"
)

type sMiddleware struct {
}

func init() {
	service.RegisterMiddleware(new())
}

// new 中间件
func new() *sMiddleware {
	return &sMiddleware{}
}

// MiddlewareHandlerResponse 响应处理
func (s *sMiddleware) MiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	ctx, span := gtrace.NewSpan(r.GetCtx(), "tracing-service-new-MiddlewareHandlerResponse")
	r.SetCtx(ctx)
	defer span.End()
	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err = r.GetError()
		res = r.GetHandlerResponse()
	)
	g.Log().Info(ctx, "MiddlewareHandlerResponse response:", res)
	if g.IsNil(res) || g.IsEmpty(res) {
		g.Log().Debug(ctx, "MiddlewareHandlerResponse statusCode:", r.Response.Status)
		r.Response.Status = http.StatusNotFound
	}
	if err != nil {
		g.Log().Error(ctx, "MiddlewareHandlerResponse err:", err)
		r.Response.Status = http.StatusInternalServerError
		if internalErr := r.Response.WriteTpl("error.html", g.Map{
			"title":   "内部错误 - 懒人科技短链平台",
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"label":   "Error",
		}); internalErr != nil {
			g.Log().Errorf(ctx, `r.Response.WriteTpl internalErr %+v`, internalErr)
		}
	}
	if r.Response.Status > 0 && r.Response.Status != http.StatusOK && r.Response.Status != http.StatusFound {
		g.Log().Info(ctx, "response:", res, "statusCode:", r.Response.Status)
		if internalErr := r.Response.WriteTpl("error.html", g.Map{
			"title":   "404 - 懒人科技短链平台",
			"code":    r.Response.Status,
			"message": "您访问的页面已失效",
			"label":   http.StatusText(r.Response.Status),
		}); internalErr != nil {
			g.Log().Errorf(ctx, `r.Response.WriteTpl 404 err: %+v`, internalErr)
		}
	}

	str := res.(*v1.HomeRes)
	g.Log().Debug(r.GetCtx(), "MiddlewareHandlerResponse end")
	if !g.IsNil(res) && !g.IsEmpty(res) {
		g.Log().Debug(r.GetCtx(), "MiddlewareHandlerResponse redirect url:", res)
		r.Response.RedirectTo(string(*str), http.StatusFound)
	}
}
