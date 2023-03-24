/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package controller

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

type cHome struct {
}

var (
	Home = cHome{}
)

// Index is the index page.
func (c *cHome) Index(ctx context.Context, req *v1.HomeIndexReq) (res *v1.HomeIndexRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-home-index")
	defer span.End()

	res = &v1.HomeIndexRes{}
	if res.HomeIndexOutput, err = service.Home().Index(ctx, req.HomeIndexInput); err != nil {
		err = gerror.Wrap(err, "home-index failed")
	}
	return
}
