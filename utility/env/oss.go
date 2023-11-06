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

// OSSEnv is the environment variable for oss
type OSSEnv struct {
	accessKeyID     string
	accessKeySecret string
	endpoint        string
	bucket          string
	domain          string
	config          map[string]string
}

// AccessKeyID get access key id
func (e *OSSEnv) AccessKeyID(_ context.Context) string {
	return e.accessKeyID
}

// AccessKeySecret get access key secret
func (e *OSSEnv) AccessKeySecret(_ context.Context) string {
	return e.accessKeySecret
}

// Endpoint get endpoint
func (e *OSSEnv) Endpoint(_ context.Context) string {
	return e.endpoint
}

// Bucket get bucket
func (e *OSSEnv) Bucket(_ context.Context) string {
	return e.bucket
}

// Domain get domain
func (e *OSSEnv) Domain(_ context.Context) string {
	return e.domain
}

// Config get config
func (e *OSSEnv) Config(_ context.Context) map[string]string {
	return e.config
}

// String get string
func (e *OSSEnv) String(_ context.Context) string {
	return `{"accessKeyID":` + e.accessKeyID + `,"accessKeySecret":` + e.accessKeySecret + `,"endpoint":` + e.endpoint +
		`,"bucket":` + e.bucket + `,"domain":` + e.domain + `}`
}

// NewOssEnv create a new oss env
func NewOssEnv(ctx context.Context, key string) (*OSSEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-NewOssEnv")
	defer span.End()

	var v, err = g.Cfg().Get(ctx, key)
	if err != nil {
		return nil, gerror.Wrap(err, "config oss get failed")
	}
	if v.IsNil() || v.IsEmpty() {
		return nil, gerror.New("config oss is empty")
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
