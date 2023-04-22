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

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/home/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

type Controller struct {
}

// New is the constructor of Controller.
func New() *Controller {
	return &Controller{}
}

// Index is the index page.
func (c *Controller) Index(ctx context.Context, req *v1.HomeIndexReq) (res *v1.HomeIndexRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-home-index")
	defer span.End()

	res = &v1.HomeIndexRes{}
	if res.HomeIndexOutput, err = service.Home().Index(ctx, req.HomeIndexInput); err != nil {
		err = gerror.Wrap(err, "home-index failed")
	}
	return
}
