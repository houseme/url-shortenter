package env

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gipv4"
	"github.com/gogf/gf/v2/net/gtrace"
)

// NewLark create app environment
func New(ctx context.Context) (*AppEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-NewLark")
	defer span.End()

	var v, err = g.Cfg().Get(ctx, "app")
	if err != nil {
		err = gerror.Wrap(err, "config app get failed")
		return nil, err
	}
	if v.IsNil() || v.IsEmpty() {
		err = gerror.New("config app is empty")
		return nil, err
	}
	var (
		config    = v.MapStrStr()
		hostIP, _ = gipv4.GetIntranetIp()
	)

	config["hostIP"] = hostIP
	return &AppEnv{
		config:         config,
		env:            config["env"],
		environment:    config["environment"],
		version:        config["version"],
		jaegerEndpoint: config["jaegerEndpoint"],
		hostIP:         hostIP,
		uploadPath:     config["uploadPath"],
		visitPath:      config["visitPath"],
		site:           config["site"],
		roleModel:      config["roleModel"],
		frontSite:      config["frontSite"],
	}, nil
}
