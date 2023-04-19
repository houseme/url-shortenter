/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package home

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/text/gstr"

	v1 "github.com/houseme/url-shortenter/app/front/api/v1/home"
	"github.com/houseme/url-shortenter/app/front/internal/service"
	"github.com/houseme/url-shortenter/utility/helper"
)

// New is the constructor for the home page controller.
func New() *cHome {
	return &cHome{}
}

type cHome struct{}

// Index is the controller for the home page.
// is the handler for the home page GET "/:short"
func (c *cHome) Index(ctx context.Context, req *v1.Req) (res *v1.Res, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-Home-Index")
	defer span.End()

	log := g.Log(helper.Helper().Logger(ctx))
	log.Debug(ctx, "home-index in:", req)
	defer func() {
		if err != nil {
			log.Error(ctx, "home-index err:", err)
		}
	}()
	var out string
	if out, err = service.Home().ShortDetail(ctx, req.HomeInput); err != nil {
		err = gerror.NewCode(gcode.CodeNotFound, "短链接不存在")
		return
	}

	if err == nil && gstr.Trim(out) == "" {
		g.RequestFromCtx(ctx).Response.Status = http.StatusNotFound
		return
	}
	res = (*v1.Res)(&out)
	log.Debug(ctx, "home-index res:", res, "url:", out)
	return
}
