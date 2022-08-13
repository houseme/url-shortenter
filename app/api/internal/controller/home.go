package controller

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/text/gstr"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	v1 "github.com/houseme/url-shortenter/app/api/api/v1"
	"github.com/houseme/url-shortenter/app/api/internal/service"
	"github.com/houseme/url-shortenter/utility"
)

var (
	// Home is the controller for the home page.
	Home = cHome{}
)

type cHome struct{}

// Index is the controller for the home page.
func (c *cHome) Index(ctx context.Context, req *v1.HomeReq) (res *v1.HomeRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-Home-Index")
	defer span.End()

	defer func() {
		if err != nil {
			g.Log(utility.Helper().Logger(ctx)).Error(ctx, "home-index err:", err)
			span.RecordError(err, trace.WithAttributes(attribute.String("home-index-err", err.Error())))
		}
	}()
	var (
		out    string
		logger = utility.Helper().Logger(ctx)
		r      = g.RequestFromCtx(ctx)
	)
	g.Log(logger).Debug(ctx, "home-index in:", req)
	req.RawQuery = r.Request.URL.RawQuery
	req.ShortAll = r.Request.URL.String()
	req.ClientIP = r.GetClientIp()
	req.UserAgent = r.UserAgent()
	req.Referer = r.Referer()
	req.Host = r.Request.Host

	g.Log(logger).Debug(ctx, "home-index modify req:", req)
	if out, err = service.Home().ShortDetail(ctx, req.HomeInput); err != nil {
		g.Log(logger).Error(ctx, "home-index err:", err)
		return
	}

	if err == nil && gstr.Trim(out) == "" {
		r.Response.Status = http.StatusNotFound
		return
	}
	res = (*v1.HomeRes)(&out)
	g.Log(logger).Debug(ctx, "home-index res:", res, "url:", out)
	return
}
