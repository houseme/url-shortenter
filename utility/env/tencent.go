package env

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

// TencentEnv .
type TencentEnv struct {
	secretID  string
	secretKey string
	region    string
	endpoint  string
}

// SecretID .
func (t *TencentEnv) SecretID(_ context.Context) string {
	return t.secretID
}

// SecretKey .
func (t *TencentEnv) SecretKey(_ context.Context) string {
	return t.secretKey
}

// Region .
func (t *TencentEnv) Region(_ context.Context) string {
	return t.region
}

// Endpoint .
func (t *TencentEnv) Endpoint(_ context.Context) string {
	return t.endpoint
}

// String .
func (t *TencentEnv) String(_ context.Context) string {
	return `{"secretId":"` + t.secretID + `","secretKey":"` + t.secretKey +
		`","region":"` + t.region + `","endpoint":"` + t.endpoint + `"}`
}

// NewTencentEnv .
func NewTencentEnv(ctx context.Context) (*TencentEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-NewTencentEnv")
	defer span.End()

	var v, err = g.Cfg().Get(ctx, "tencent")
	if err != nil {
		err = gerror.Wrap(err, "config tencent get failed")
		return nil, err
	}
	if v.IsNil() || v.IsEmpty() {
		err = gerror.New("config tencent is empty")
		return nil, err
	}

	var config = v.MapStrStr()

	return &TencentEnv{
		secretID:  config["secretID"],
		secretKey: config["secretKey"],
		region:    config["region"],
		endpoint:  config["endpoint"],
	}, nil
}
