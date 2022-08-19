package env

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"
)

// AlibabaEnv .
type AlibabaEnv struct {
	accessKeyID     string
	accessKeySecret string
	bucketName      string
	endpoint        string
	region          string
	ctx             context.Context
}

// AccessKeyID .
func (a *AlibabaEnv) AccessKeyID(ctx context.Context) string {
	a.ctx = ctx
	return a.accessKeyID
}

// AccessKeySecret .
func (a *AlibabaEnv) AccessKeySecret(ctx context.Context) string {
	a.ctx = ctx
	return a.accessKeySecret
}

// BucketName .
func (a *AlibabaEnv) BucketName(ctx context.Context) string {
	a.ctx = ctx
	return a.bucketName
}

// Endpoint .
func (a *AlibabaEnv) Endpoint(ctx context.Context) string {
	a.ctx = ctx
	return a.endpoint
}

// Region .
func (a *AlibabaEnv) Region(ctx context.Context) string {
	a.ctx = ctx
	return a.region
}

// String .
func (a *AlibabaEnv) String(ctx context.Context) string {
	a.ctx = ctx
	return `{"accessKeyId":"` + a.accessKeyID + `","accessKeySecret":"` + a.accessKeySecret +
		`","bucketName":"` + a.bucketName + `","endpoint":"` + a.endpoint + `","region":"` + a.region + `"}`
}

// NewAlibabaEnv .
func NewAlibabaEnv(ctx context.Context) (*AlibabaEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-NewAlibabaEnv")
	defer span.End()

	var (
		v, err = g.Cfg().Get(ctx, "alibaba")
		logger = gconv.String(ctx.Value("logger"))
	)

	if err != nil {
		err = gerror.Wrap(err, "config alibaba get failed")
		return nil, err
	}
	if v.IsNil() || v.IsEmpty() {
		g.Log(logger).Info(ctx, " config alibaba is empty")
		err = gerror.New("config alibaba is empty")
		return nil, err
	}
	g.Log(logger).Debug(ctx, " config alibaba:", v)
	var (
		config = v.MapStrStr()
		env    = &AlibabaEnv{
			accessKeyID:     config["accessKeyID"],
			accessKeySecret: config["accessKeySecret"],
			region:          config["region"],
			endpoint:        config["endpoint"],
			bucketName:      config["bucketName"],
			ctx:             ctx,
		}
	)
	g.Log(logger).Debug(ctx, " config alibaba:", env.String(ctx))
	return env, nil
}
