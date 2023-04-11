// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

// Package main is the main package for the schedule service.
package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/houseme/url-shortenter/app/schedule/internal/cmd"
	_ "github.com/houseme/url-shortenter/app/schedule/internal/logic"
	_ "github.com/houseme/url-shortenter/app/schedule/internal/packed"
	"github.com/houseme/url-shortenter/internal/tracing"
)

func main() {
	ctx := gctx.New()
	tracing.InitTracer(ctx, "tracing-shortenter-schedule")
	cmd.Main.Run(ctx)
}
