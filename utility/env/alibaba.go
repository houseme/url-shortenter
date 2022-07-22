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
	accessKeyID     string `json:"accessKeyId"`
	accessKeySecret string `json:"accessKeySecret"`
	bucketName      string `json:"bucketName"`
	endpoint        string `json:"endpoint"`
	region          string `json:"region"`
}

// AccessKeyID .
func (a *AlibabaEnv) AccessKeyID(ctx context.Context) string {
	return a.accessKeyID
}

// AccessKeySecret .
func (a *AlibabaEnv) AccessKeySecret(ctx context.Context) string {
	return a.accessKeySecret
}

// BucketName .
func (a *AlibabaEnv) BucketName(ctx context.Context) string {
	return a.bucketName
}

// Endpoint .
func (a *AlibabaEnv) Endpoint(ctx context.Context) string {
	return a.endpoint
}

// Region .
func (a *AlibabaEnv) Region(ctx context.Context) string {
	return a.region
}

// String .
func (a *AlibabaEnv) String() string {
	return `{"accessKeyId":"` + a.accessKeyID + `","accessKeySecret":"` + a.accessKeySecret + `","bucketName":"` + a.bucketName + `","endpoint":"` + a.endpoint + `","region":"` + a.region + `"}`
}

// NewAlibabaEnv .
func NewAlibabaEnv(ctx context.Context) (*AlibabaEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-NewAlibabaEnv")
	defer span.End()

	var (
		v, err = g.Cfg().Get(ctx, "alibaba")
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
		err = gerror.New("config app is empty")
		return nil, err
	}

	var env *AlibabaEnv
	if err = v.Scan(&env); err != nil {
		g.Log(logger).Error(ctx, " config app scan fail err:", err)
		err = gerror.Wrap(err, "config app scan failed")
		return nil, err
	}
	g.Log(logger).Info(ctx, " config app:", env)
	return env, nil
}
