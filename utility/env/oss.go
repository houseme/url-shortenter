package env

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"
)

// ossEnv is the environment variable for oss
type ossEnv struct {
	accessKeyID     string                 `json:"accessKeyId"`
	accessKeySecret string                 `json:"accessKeySecret"`
	endpoint        string                 `json:"endpoint"`
	bucket          string                 `json:"bucket"`
	domain          string                 `json:"domain"`
	config          map[string]interface{} `json:"config"`
}

// NewOssEnv create a new oss env
func NewOssEnv(ctx context.Context, key string) (*ossEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-NewOssEnv")
	defer span.End()

	var (
		v, err = g.Cfg().Get(ctx, key)
		logger = gconv.String(ctx.Value("logger"))
	)

	defer func() {
		span.RecordError(err)
	}()

	if err != nil {
		g.Log(logger).Error(ctx, " config oss fail err:", err)
		err = gerror.Wrap(err, "config oss get failed")
		return nil, err
	}
	if v.IsNil() || v.IsEmpty() {
		g.Log(logger).Info(ctx, " config oss is empty")
		err = gerror.Wrap(errors.New("config oss is empty"), "config oss is empty")
		return nil, err
	}
	var config = v.MapStrAny()
	return &ossEnv{
		domain:          gconv.String(config["domain"]),
		accessKeyID:     gconv.String(config["accessKeyId"]),
		accessKeySecret: gconv.String(config["accessKeySecret"]),
		endpoint:        gconv.String(config["endpoint"]),
		bucket:          gconv.String(config["bucket"]),
		config:          config,
	}, nil
}

// AccessKeyID get access key id
func (o *ossEnv) AccessKeyID(ctx context.Context) string {
	return o.accessKeyID
}

// AccessKeySecret get access key secret
func (o *ossEnv) AccessKeySecret(ctx context.Context) string {
	return o.accessKeySecret
}

// Endpoint get endpoint
func (o *ossEnv) Endpoint(ctx context.Context) string {
	return o.endpoint
}

// Bucket get bucket
func (o *ossEnv) Bucket(ctx context.Context) string {
	return o.bucket
}

// Domain get domain
func (o *ossEnv) Domain(ctx context.Context) string {
	return o.domain
}

// Config get config
func (o *ossEnv) Config(ctx context.Context) map[string]interface{} {
	return o.config
}
