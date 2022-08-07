package env

import (
	"context"
)

// AppEnv .
type AppEnv struct {
	env            string
	environment    string
	version        string
	jaegerEndpoint string
	hostIP         string
	config         map[string]interface{}
	uploadPath     string
	visitPath      string
	site           string
	roleModel      string
	frontSite      string
	ctx            context.Context
}

// Env .
func (a *AppEnv) Env(ctx context.Context) string {
	a.ctx = ctx
	return a.env
}

// Environment .
func (a *AppEnv) Environment(ctx context.Context) string {
	a.ctx = ctx
	return a.environment
}

// Version .
func (a *AppEnv) Version(ctx context.Context) string {
	a.ctx = ctx
	return a.version
}

// JaegerEndpoint .
func (a *AppEnv) JaegerEndpoint(ctx context.Context) string {
	a.ctx = ctx
	return a.jaegerEndpoint
}

// Config .获取配置信息
func (a *AppEnv) Config(ctx context.Context) map[string]interface{} {
	a.ctx = ctx
	return a.config
}

// HostIP . 获取本机IP
func (a *AppEnv) HostIP(ctx context.Context) string {
	a.ctx = ctx
	return a.hostIP
}

// UploadPath .	上传路径
func (a *AppEnv) UploadPath(ctx context.Context) string {
	a.ctx = ctx
	return a.uploadPath
}

// VisitPath file server访问路径
func (a *AppEnv) VisitPath(ctx context.Context) string {
	a.ctx = ctx
	return a.visitPath
}

// Site .网站名称
func (a *AppEnv) Site(ctx context.Context) string {
	a.ctx = ctx
	return a.site
}

// RoleModel .
func (a *AppEnv) RoleModel(ctx context.Context) string {
	a.ctx = ctx
	return a.roleModel
}

// FrontSite .
func (a *AppEnv) FrontSite(ctx context.Context) string {
	a.ctx = ctx
	return a.frontSite
}
