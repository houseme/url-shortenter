package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/houseme/url-shortenter/app/api/internal/packed"

	_ "github.com/houseme/url-shortenter/app/api/internal/logic"

	"github.com/houseme/url-shortenter/internal/tracing"
	"github.com/houseme/url-shortenter/utility/env"

	"github.com/houseme/url-shortenter/app/api/internal/cmd"
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
