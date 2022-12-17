package env

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

// OSSEnv is the environment variable for oss
type OSSEnv struct {
	accessKeyID     string
	accessKeySecret string
	endpoint        string
	bucket          string
	domain          string
	config          map[string]string
}

// NewOssEnv create a new oss env
func NewOssEnv(ctx context.Context, key string) (*OSSEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-NewOssEnv")
	defer span.End()

	var v, err = g.Cfg().Get(ctx, key)
	if err != nil {
		err = gerror.Wrap(err, "config oss get failed")
		return nil, err
	}
	if v.IsNil() || v.IsEmpty() {
		err = gerror.New("config oss is empty")
		return nil, err
	}
	var config = v.MapStrStr()
	return &OSSEnv{
		domain:          config["domain"],
		accessKeyID:     config["accessKeyId"],
		accessKeySecret: config["accessKeySecret"],
		endpoint:        config["endpoint"],
		bucket:          config["bucket"],
		config:          config,
	}, nil
}

// AccessKeyID get access key id
func (o *OSSEnv) AccessKeyID(_ context.Context) string {
	return o.accessKeyID
}

// AccessKeySecret get access key secret
func (o *OSSEnv) AccessKeySecret(_ context.Context) string {
	return o.accessKeySecret
}

// Endpoint get endpoint
func (o *OSSEnv) Endpoint(_ context.Context) string {
	return o.endpoint
}

// Bucket get bucket
func (o *OSSEnv) Bucket(_ context.Context) string {
	return o.bucket
}

// Domain get domain
func (o *OSSEnv) Domain(_ context.Context) string {
	return o.domain
}

// Config get config
func (o *OSSEnv) Config(_ context.Context) map[string]string {
	return o.config
}

// String get string
func (o *OSSEnv) String(_ context.Context) string {
	return `{"accessKeyID":` + o.accessKeyID + `,"accessKeySecret":` + o.accessKeySecret + `,"endpoint":` + o.endpoint +
		`,"bucket":` + o.bucket + `,"domain":` + o.domain + `}`
}
