// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/houseme/url-shortenter/app/api/internal/cmd"
	_ "github.com/houseme/url-shortenter/app/api/internal/logic"
	_ "github.com/houseme/url-shortenter/app/api/internal/packed"
	"github.com/houseme/url-shortenter/internal/tracing"
	"github.com/houseme/url-shortenter/utility/env"
)

func main() {
	var (
		ctx         = gctx.New()
		appEnv, err = env.New(ctx)
	)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	_, err = tracing.InitJaeger("tracing-shortenter-api", appEnv.JaegerEndpoint(ctx), appEnv.Version(ctx), appEnv.Environment(ctx), appEnv.HostIP(ctx))
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	cmd.Main.Run(ctx)
}
