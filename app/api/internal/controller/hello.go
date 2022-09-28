// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/houseme/url-shortenter/app/api/api/v1"
)

var (
	// Hello is a hello world.
	Hello = cHello{}
)

type cHello struct{}

// Hello is a hello world.
func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.Log().Debug(ctx, "hello request params:", req)
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
