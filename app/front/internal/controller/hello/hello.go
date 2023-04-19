/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package hello

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/houseme/url-shortenter/app/front/api/v1/hello"
)

func New() *cHello {
	return &cHello{}
}

type cHello struct{}

// Hello is a hello world.
func (c *cHello) Hello(ctx context.Context, req *v1.Req) (res *v1.Res, err error) {
	g.Log().Debug(ctx, "hello request params:", req)
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
