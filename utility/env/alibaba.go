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

// AlibabaEnv .
type AlibabaEnv struct {
	accessKeyID     string
	accessKeySecret string
	bucketName      string
	endpoint        string
	region          string
	config          map[string]string
}

// AccessKeyID access key id
func (e *AlibabaEnv) AccessKeyID(_ context.Context) string {
	return e.accessKeyID
}

// AccessKeySecret access key secret
func (e *AlibabaEnv) AccessKeySecret(_ context.Context) string {
	return e.accessKeySecret
}

// BucketName bucket name
func (e *AlibabaEnv) BucketName(_ context.Context) string {
	return e.bucketName
}

// Endpoint endpoint
func (e *AlibabaEnv) Endpoint(_ context.Context) string {
	return e.endpoint
}

// Region .
func (e *AlibabaEnv) Region(_ context.Context) string {
	return e.region
}

// Config .
func (e *AlibabaEnv) Config(_ context.Context) map[string]string {
	return e.config
}

// String .
func (e *AlibabaEnv) String(_ context.Context) string {
	return `{"accessKeyId":"` + e.accessKeyID + `","accessKeySecret":"` + e.accessKeySecret +
		`","bucketName":"` + e.bucketName + `","endpoint":"` + e.endpoint + `","region":"` + e.region + `"}`
}

// NewAlibabaEnv new alibaba env return alibaba env
func NewAlibabaEnv(ctx context.Context) (*AlibabaEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-NewAlibabaEnv")
	defer span.End()

	v, err := g.Cfg().Get(ctx, "alibaba")
	if err != nil {
		return nil, gerror.Wrap(err, "config alibaba get failed")
	}
	if v.IsNil() || v.IsEmpty() {
		return nil, gerror.New("config alibaba is empty")
	}
	config := v.MapStrStr()
	return &AlibabaEnv{
		accessKeyID:     config["accessKeyID"],
		accessKeySecret: config["accessKeySecret"],
		region:          config["region"],
		endpoint:        config["endpoint"],
		bucketName:      config["bucketName"],
		config:          config,
	}, nil
}
