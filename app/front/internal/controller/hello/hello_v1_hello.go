// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package hello

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "github.com/houseme/url-shortenter/app/front/api/hello/v1"
)

// Hello is a demo controller for output "Hello World!". It's used for test purpose.
func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.Log().Debug(ctx, "hello request params:", req)
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
