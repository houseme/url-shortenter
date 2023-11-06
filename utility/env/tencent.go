// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

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
	config    map[string]string
}

// SecretID .
func (e *TencentEnv) SecretID(_ context.Context) string {
	return e.secretID
}

// SecretKey .
func (e *TencentEnv) SecretKey(_ context.Context) string {
	return e.secretKey
}

// Region .
func (e *TencentEnv) Region(_ context.Context) string {
	return e.region
}

// Endpoint .
func (e *TencentEnv) Endpoint(_ context.Context) string {
	return e.endpoint
}

// Config .
func (e *TencentEnv) Config(_ context.Context) map[string]string {
	return e.config
}

// String .
func (e *TencentEnv) String(_ context.Context) string {
	return `{"secretId":"` + e.secretID + `","secretKey":"` + e.secretKey +
		`","region":"` + e.region + `","endpoint":"` + e.endpoint + `"}`
}

// NewTencentEnv .
func NewTencentEnv(ctx context.Context) (*TencentEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-NewTencentEnv")
	defer span.End()

	var v, err = g.Cfg().Get(ctx, "tencent")
	if err != nil {
		return nil, gerror.Wrap(err, "config tencent get failed")
	}
	if v.IsNil() || v.IsEmpty() {
		return nil, gerror.New("config tencent is empty")
	}

	var config = v.MapStrStr()
	return &TencentEnv{
		secretID:  config["secretID"],
		secretKey: config["secretKey"],
		region:    config["region"],
		endpoint:  config["endpoint"],
		config:    config,
	}, nil
}
