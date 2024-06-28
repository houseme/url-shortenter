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

// AppEnv .
type AppEnv struct {
	config      map[string]string
	env         string
	environment string
	version     string
	useTrace    string
	traceType   string
	endpoint    string
	traceToken  string
	service     string
	application string
	uploadPath  string
	visitPath   string
	site        string
	roleModel   string
	frontSite   string
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

// UseTrace .
func (e *AppEnv) UseTrace(_ context.Context) string {
	return e.useTrace
}

// TraceType .
func (e *AppEnv) TraceType(_ context.Context) string {
	return e.traceType
}

// Endpoint .
func (e *AppEnv) Endpoint(_ context.Context) string {
	return e.endpoint
}

// TraceToken .
func (e *AppEnv) TraceToken(_ context.Context) string {
	return e.traceToken
}

// Service .
func (e *AppEnv) Service(_ context.Context) string {
	return e.service
}

// Application .
func (e *AppEnv) Application(_ context.Context) string {
	return e.application
}

// Config .获取配置信息
func (e *AppEnv) Config(_ context.Context) map[string]string {
	return e.config
}

// UploadPath .	上传路径
func (e *AppEnv) UploadPath(_ context.Context) string {
	return e.uploadPath
}

// VisitPath file server 访问路径
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

// ApplicationService .
func (e *AppEnv) ApplicationService() string {
	return e.application + "-" + e.service
}

// String
func (e *AppEnv) String(_ context.Context) string {
	return `{"env":"` + e.env + `","environment":"` + e.environment + `","version":"` + e.version + `","useTrace":"` + e.useTrace +
		`","endpoint":"` + e.endpoint + `","traceToken":"` + e.traceToken + `","traceType":"` + e.traceType + `","uploadPath":"` + e.uploadPath +
		`","visitPath":"` + e.visitPath + `","service":"` + e.service + `","application":"` + e.application +
		`","site":"` + e.site + `","roleModel":"` + e.roleModel + `","frontSite":"` + e.frontSite + `"}`
}

// New create app environment
func New(ctx context.Context) (*AppEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-env-New")
	defer span.End()

	value, err := g.Cfg().Get(ctx, "app")
	if err != nil {
		return nil, gerror.Wrap(err, "config app get failed")
	}
	if value.IsNil() || value.IsEmpty() {
		return nil, gerror.New("config app is empty")
	}
	config := value.MapStrStr()
	return &AppEnv{
		config:      config,
		env:         config["env"],
		environment: config["environment"],
		version:     config["version"],
		useTrace:    config["useTrace"],
		traceType:   config["traceType"],
		endpoint:    config["endpoint"],
		traceToken:  config["traceToken"],
		service:     config["service"],
		application: config["application"],
		uploadPath:  config["uploadPath"],
		visitPath:   config["visitPath"],
		site:        config["site"],
		roleModel:   config["roleModel"],
		frontSite:   config["frontSite"],
	}, nil
}
