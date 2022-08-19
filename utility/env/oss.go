package env

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"
)

// ossEnv is the environment variable for oss
type ossEnv struct {
	accessKeyID     string
	accessKeySecret string
	endpoint        string
	bucket          string
	domain          string
	config          map[string]string
	ctx             context.Context
}

// NewOssEnv create a new oss env
func NewOssEnv(ctx context.Context, key string) (*ossEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-NewOssEnv")
	defer span.End()

	var (
		v, err = g.Cfg().Get(ctx, key)
		logger = gconv.String(ctx.Value("logger"))
	)

	if err != nil {
		err = gerror.Wrap(err, "config oss get failed")
		return nil, err
	}
	if v.IsNil() || v.IsEmpty() {
		g.Log(logger).Info(ctx, " config oss is empty")
		err = gerror.New("config oss is empty")
		return nil, err
	}
	var config = v.MapStrStr()
	return &ossEnv{
		domain:          config["domain"],
		accessKeyID:     config["accessKeyId"],
		accessKeySecret: config["accessKeySecret"],
		endpoint:        config["endpoint"],
		bucket:          config["bucket"],
		config:          config,
	}, nil
}

// AccessKeyID get access key id
func (o *ossEnv) AccessKeyID(ctx context.Context) string {
	o.ctx = ctx
	return o.accessKeyID
}

// AccessKeySecret get access key secret
func (o *ossEnv) AccessKeySecret(ctx context.Context) string {
	o.ctx = ctx
	return o.accessKeySecret
}

// Endpoint get endpoint
func (o *ossEnv) Endpoint(ctx context.Context) string {
	o.ctx = ctx
	return o.endpoint
}

// Bucket get bucket
func (o *ossEnv) Bucket(ctx context.Context) string {
	o.ctx = ctx
	return o.bucket
}

// Domain get domain
func (o *ossEnv) Domain(ctx context.Context) string {
	o.ctx = ctx
	return o.domain
}

// Config get config
func (o *ossEnv) Config(ctx context.Context) map[string]string {
	o.ctx = ctx
	return o.config
}

// String get string
func (o *ossEnv) String(ctx context.Context) string {
	o.ctx = ctx
	return `{"accessKeyID":` + o.accessKeyID + `,"accessKeySecret":` + o.accessKeySecret + `,"endpoint":` + o.endpoint +
		`,"bucket":` + o.bucket + `,"domain":` + o.domain + `}`
}
