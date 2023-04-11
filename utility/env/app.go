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
	"github.com/gogf/gf/v2/net/gipv4"
	"github.com/gogf/gf/v2/net/gtrace"
)

// AppEnv .
type AppEnv struct {
	config         map[string]string
	env            string
	environment    string
	version        string
	jaegerEndpoint string
	hostIP         string
	uploadPath     string
	visitPath      string
	site           string
	roleModel      string
	frontSite      string
}

// Env .
func (e *AppEnv) Env(_ context.Context) string {
	return e.env
}

// Environment .
func (e *AppEnv) Environment(_ context.Context) string {
	return e.environment
}

// Version .
func (e *AppEnv) Version(_ context.Context) string {
	return e.version
}

// JaegerEndpoint .
func (e *AppEnv) JaegerEndpoint(_ context.Context) string {
	return e.jaegerEndpoint
}

// Config .获取配置信息
func (e *AppEnv) Config(_ context.Context) map[string]string {
	return e.config
}

// HostIP . 获取本机IP
func (e *AppEnv) HostIP(_ context.Context) string {
	return e.hostIP
}

// UploadPath .	上传路径
func (e *AppEnv) UploadPath(_ context.Context) string {
	return e.uploadPath
}

// VisitPath file server访问路径
func (e *AppEnv) VisitPath(_ context.Context) string {
	return e.visitPath
}

// Site .网站名称
func (e *AppEnv) Site(_ context.Context) string {
	return e.site
}

// RoleModel .
func (e *AppEnv) RoleModel(_ context.Context) string {
	return e.roleModel
}

// FrontSite .
func (e *AppEnv) FrontSite(_ context.Context) string {
	return e.frontSite
}

// String
func (e *AppEnv) String(_ context.Context) string {
	return `{"env":"` + e.env + `","environment":"` + e.environment + `","version":"` + e.version +
		`","jaegerEndpoint":"` + e.jaegerEndpoint + `","hostIP":"` + e.hostIP +
		`","uploadPath":"` + e.uploadPath + `","visitPath":"` + e.visitPath +
		`","site":"` + e.site + `","roleModel":"` + e.roleModel + `","frontSite":"` + e.frontSite + `"}`
}

// New create app environment
func New(ctx context.Context) (*AppEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-New")
	defer span.End()

	var v, err = g.Cfg().Get(ctx, "app")
	if err != nil {
		err = gerror.Wrap(err, "config app get failed")
		return nil, err
	}
	if v.IsNil() || v.IsEmpty() {
		err = gerror.New("config app is empty")
		return nil, err
	}
	var (
		config    = v.MapStrStr()
		hostIP, _ = gipv4.GetIntranetIp()
	)

	config["hostIP"] = hostIP
	return &AppEnv{
		config:         config,
		env:            config["env"],
		environment:    config["environment"],
		version:        config["version"],
		jaegerEndpoint: config["jaegerEndpoint"],
		hostIP:         hostIP,
		uploadPath:     config["uploadPath"],
		visitPath:      config["visitPath"],
		site:           config["site"],
		roleModel:      config["roleModel"],
		frontSite:      config["frontSite"],
	}, nil
}
