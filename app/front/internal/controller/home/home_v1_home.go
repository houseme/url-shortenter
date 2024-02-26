// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package home

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/text/gstr"

	v1 "github.com/houseme/url-shortenter/app/front/api/home/v1"
	"github.com/houseme/url-shortenter/app/front/internal/service"
	"github.com/houseme/url-shortenter/utility/helper"
)

// Home is the controller for the API.
func (c *ControllerV1) Home(ctx context.Context, req *v1.HomeReq) (res *v1.HomeRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-Home-Index")
	defer span.End()

	logger := g.Log(helper.Helper().Logger(ctx))
	logger.Debug(ctx, "home index in:", req)
	var out string
	if out, err = service.Home().ShortDetail(ctx, req.HomeInput); err != nil {
		err = gerror.NewCode(gcode.CodeNotFound, "The short link does not exist")
		return
	}

	if err == nil && gstr.Trim(out) == "" {
		g.RequestFromCtx(ctx).Response.Status = http.StatusNotFound
		return
	}
	res = (*v1.HomeRes)(&out)
	logger.Debug(ctx, "home-index res:", res, "url:", out)
	return
}
