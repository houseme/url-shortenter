/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

// Package hello is a hello world.
package hello

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/houseme/url-shortenter/app/front/api/hello/v1"
)

// New is a hello world.
func New() *Controller {
	return &Controller{}
}

// Controller .
type Controller struct{}

// Hello is a hello world.
func (c *Controller) Hello(ctx context.Context, req *v1.Req) (res *v1.Res, err error) {
	g.Log().Debug(ctx, "hello request params:", req)
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
