package env

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"
)

// TencentEnv .
type TencentEnv struct {
	secretID  string
	secretKey string
	region    string
	endpoint  string
	ctx       context.Context
}

// SecretID .
func (a *TencentEnv) SecretID(ctx context.Context) string {
	a.ctx = ctx
	return a.secretID
}

// SecretKey .
func (a *TencentEnv) SecretKey(ctx context.Context) string {
	a.ctx = ctx
	return a.secretKey
}

// Region .
func (a *TencentEnv) Region(ctx context.Context) string {
	a.ctx = ctx
	return a.region
}

// Endpoint .
func (a *TencentEnv) Endpoint(ctx context.Context) string {
	a.ctx = ctx
	return a.endpoint
}

// String .
func (a *TencentEnv) String(ctx context.Context) string {
	a.ctx = ctx
	return `{"secretId":"` + a.secretID + `","secretKey":"` + a.secretKey +
		`","region":"` + a.region + `","endpoint":"` + a.endpoint + `"}`
}

// NewTencentEnv .
func NewTencentEnv(ctx context.Context) (*TencentEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-NewTencentEnv")
	defer span.End()

	var (
		v, err = g.Cfg().Get(ctx, "tencent")
		logger = gconv.String(ctx.Value("logger"))
	)

	if err != nil {
		g.Log(logger).Error(ctx, " config tencent fail err:", err)
		err = gerror.Wrap(err, "config tencent get failed")
		return nil, err
	}
	if v.IsNil() || v.IsEmpty() {
		g.Log(logger).Info(ctx, " config tencent is empty")
		err = gerror.New("config tencent is empty")
		return nil, err
	}

	var (
		config = v.MapStrStr()
		env    = &TencentEnv{
			secretID:  config["secretID"],
			secretKey: config["secretKey"],
			region:    config["region"],
			endpoint:  config["endpoint"],
			ctx:       ctx,
		}
	)

	g.Log(logger).Info(ctx, " config app:", env.String(ctx))
	return env, nil
}
