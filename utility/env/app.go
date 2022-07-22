package env

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gipv4"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"
)

// New  创建APP环境
func New(ctx context.Context) (*appEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-New")
	defer span.End()

	var (
		v, err = g.Cfg().Get(ctx, "app")
		logger = gconv.String(ctx.Value("logger"))
	)

	defer func() {
		span.RecordError(err)
	}()

	if err != nil {
		g.Log(logger).Error(ctx, " config app fail err:", err)
		err = gerror.Wrap(err, "config app get failed")
		return nil, err
	}
	if v.IsNil() || v.IsEmpty() {
		g.Log(logger).Info(ctx, " config app is empty")
		err = gerror.Wrap(errors.New("config app is empty"), "config app is empty")
		return nil, err
	}
	var config = v.MapStrAny()

	hostIP, _ := gipv4.GetIntranetIp()
	config["hostIP"] = hostIP
	return &appEnv{
		config:         config,
		env:            gconv.String(config["env"]),
		environment:    gconv.String(config["environment"]),
		version:        gconv.String(config["version"]),
		datacenterID:   gconv.Int64(config["datacenterID"]),
		workerID:       gconv.Int64(config["workerID"]),
		jaegerEndpoint: gconv.String(config["jaegerEndpoint"]),
		hostIP:         hostIP,
		uploadPath:     gconv.String(config["uploadPath"]),
		visitPath:      gconv.String(config["visitPath"]),
		site:           gconv.String(config["site"]),
		roleModel:      gconv.String(config["roleModel"]),
		frontSite:      gconv.String(config["frontSite"]),
	}, nil
}
