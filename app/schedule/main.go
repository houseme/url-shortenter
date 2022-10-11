package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/houseme/url-shortenter/app/schedule/internal/cmd"
	_ "github.com/houseme/url-shortenter/app/schedule/internal/logic"
	_ "github.com/houseme/url-shortenter/app/schedule/internal/packed"
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
	_, err = tracing.InitJaeger("tracing-shortenter-schedule", appEnv.JaegerEndpoint(ctx), appEnv.Version(ctx), appEnv.Environment(ctx), appEnv.HostIP(ctx))
	if err != nil {
		g.Log().Fatal(ctx, err)
	}
	cmd.Main.Run(ctx)
}
