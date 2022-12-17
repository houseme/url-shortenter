// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package controller

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
	"github.com/houseme/url-shortenter/utility/helper"
)

type cShort struct {
}

// Short is the controller for the short page.
var Short = cShort{}

// CreateShort is the handler for CreateShort
func (c *cShort) CreateShort(ctx context.Context, req *v1.CreateShortReq) (res *v1.CreateShortRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-CreateShort")
	defer span.End()

	var logger = helper.Helper().Logger(ctx)
	res = &v1.CreateShortRes{}
	if res.CreateShortOutput, err = service.Short().CreateShort(ctx, req.CreateShortInput); err != nil {
		g.Log(logger).Error(ctx, "short-CreateShort err:", err)
		err = gerror.Wrap(err, "account-ModifyAccount err:")
	}
	return
}

// ModifyShort is the handler for ModifyShort
func (c *cShort) ModifyShort(ctx context.Context, req *v1.ModifyShortReq) (res *v1.ModifyShortRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-ModifyShort")
	defer span.End()

	var logger = helper.Helper().Logger(ctx)
	res = &v1.ModifyShortRes{}
	if res.ModifyShortOutput, err = service.Short().ModifyShort(ctx, req.ModifyShortInput); err != nil {
		g.Log(logger).Error(ctx, "short-ModifyShort err:", err)
		err = gerror.Wrap(err, "short-ModifyShort err:")
	}
	return
}

// QueryShort is the handler for QueryShort
func (c *cShort) QueryShort(ctx context.Context, req *v1.QueryShortReq) (res *v1.QueryShortRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-QueryShort")
	defer span.End()

	var logger = helper.Helper().Logger(ctx)
	res = &v1.QueryShortRes{}
	if res.QueryShortOutput, err = service.Short().QueryShort(ctx, req.QueryShortInput); err != nil {
		g.Log(logger).Error(ctx, "short-QueryShort err:", err)
		err = gerror.Wrap(err, "short-QueryShort err:")
	}
	return
}

// QueryStat is the handler for QueryStat
func (c *cShort) QueryStat(ctx context.Context, req *v1.QueryStatReq) (res *v1.QueryStatRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-QueryStat")
	defer span.End()

	var logger = helper.Helper().Logger(ctx)
	res = &v1.QueryStatRes{}
	if res.QueryStatOutput, err = service.Short().QueryStat(ctx, req.QueryStatInput); err != nil {
		g.Log(logger).Error(ctx, "short-QueryStat err:", err)
		err = gerror.Wrap(err, "short-QueryStat err:")
	}
	return
}
