package env

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gipv4"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"
)

// New  创建APP环境
func New(ctx context.Context) (*AppEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-New")
	defer span.End()

	var v, err = g.Cfg().Get(ctx, "app")
	defer func() {
		span.RecordError(err)
	}()

	if err != nil {
		err = gerror.Wrap(err, "config app get failed")
		return nil, err
	}
	if v.IsNil() || v.IsEmpty() {
		err = gerror.New("config app is empty")
		return nil, err
	}
	var config = v.MapStrAny()

	hostIP, _ := gipv4.GetIntranetIp()
	config["hostIP"] = hostIP
	return &AppEnv{
		config:         config,
		env:            gconv.String(config["env"]),
		environment:    gconv.String(config["environment"]),
		version:        gconv.String(config["version"]),
		jaegerEndpoint: gconv.String(config["jaegerEndpoint"]),
		hostIP:         hostIP,
		uploadPath:     gconv.String(config["uploadPath"]),
		visitPath:      gconv.String(config["visitPath"]),
		site:           gconv.String(config["site"]),
		roleModel:      gconv.String(config["roleModel"]),
		frontSite:      gconv.String(config["frontSite"]),
	}, nil
}
