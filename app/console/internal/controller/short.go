package controller

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	v1 "github.com/houseme/url-shortenter/app/console/api/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
	"github.com/houseme/url-shortenter/utility"
)

type cShort struct {
}

// Short is the controller for the short page.
var Short = cShort{}

// CreateShort is the handler for CreateShort
func (c *cShort) CreateShort(ctx context.Context, req *v1.CreateShortReq) (res *v1.CreateShortRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-CreateShort")
	defer span.End()

	var logger = utility.Helper().Logger(ctx)

	defer func() {
		if err != nil {
			g.Log(logger).Error(ctx, "short-CreateShort err:", err)
			span.RecordError(err, trace.WithAttributes(attribute.String("short-CreateShort-err", err.Error())))
		}
	}()
	res = &v1.CreateShortRes{}
	if res.CreateShortOutput, err = service.Short().CreateShort(ctx, req.CreateShortInput); err != nil {
		err = gerror.Wrap(err, "account-ModifyAccount err:")
	}
	return
}

// ModifyShort is the handler for ModifyShort
func (c *cShort) ModifyShort(ctx context.Context, req *v1.ModifyShortReq) (res *v1.ModifyShortRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-ModifyShort")
	defer span.End()

	var logger = utility.Helper().Logger(ctx)

	defer func() {
		if err != nil {
			g.Log(logger).Error(ctx, "short-ModifyShort err:", err)
			span.RecordError(err, trace.WithAttributes(attribute.String("short-ModifyShort-err", err.Error())))
		}
	}()
	res = &v1.ModifyShortRes{}
	if res.ModifyShortOutput, err = service.Short().ModifyShort(ctx, req.ModifyShortInput); err != nil {
		err = gerror.Wrap(err, "short-ModifyShort err:")
	}
	return
}

// QueryShort is the handler for QueryShort
func (c *cShort) QueryShort(ctx context.Context, req *v1.QueryShortReq) (res *v1.QueryShortRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-QueryShort")
	defer span.End()

	var logger = utility.Helper().Logger(ctx)

	defer func() {
		if err != nil {
			g.Log(logger).Error(ctx, "short-QueryShort err:", err)
			span.RecordError(err, trace.WithAttributes(attribute.String("short-QueryShort-err", err.Error())))
		}
	}()
	res = &v1.QueryShortRes{}
	if res.QueryShortOutput, err = service.Short().QueryShort(ctx, req.QueryShortInput); err != nil {
		err = gerror.Wrap(err, "short-QueryShort err:")
	}
	return
}

// QueryStat is the handler for QueryStat
func (c *cShort) QueryStat(ctx context.Context, req *v1.QueryStatReq) (res *v1.QueryStatRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-QueryStat")
	defer span.End()

	var logger = utility.Helper().Logger(ctx)

	defer func() {
		if err != nil {
			g.Log(logger).Error(ctx, "short-QueryStat err:", err)
			span.RecordError(err, trace.WithAttributes(attribute.String("short-QueryStat-err", err.Error())))
		}
	}()
	res = &v1.QueryStatRes{}
	if res.QueryStatOutput, err = service.Short().QueryStat(ctx, req.QueryStatInput); err != nil {
		err = gerror.Wrap(err, "short-QueryStat err:")
	}
	return
}
