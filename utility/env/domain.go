package env

import (
	"context"
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
func (a *AppEnv) Env(_ context.Context) string {
	return a.env
}

// Environment .
func (a *AppEnv) Environment(_ context.Context) string {
	return a.environment
}

// Version .
func (a *AppEnv) Version(_ context.Context) string {
	return a.version
}

// JaegerEndpoint .
func (a *AppEnv) JaegerEndpoint(_ context.Context) string {
	return a.jaegerEndpoint
}

// Config .获取配置信息
func (a *AppEnv) Config(_ context.Context) map[string]string {
	return a.config
}

// HostIP . 获取本机IP
func (a *AppEnv) HostIP(_ context.Context) string {
	return a.hostIP
}

// UploadPath .	上传路径
func (a *AppEnv) UploadPath(_ context.Context) string {
	return a.uploadPath
}

// VisitPath file server访问路径
func (a *AppEnv) VisitPath(_ context.Context) string {
	return a.visitPath
}

// Site .网站名称
func (a *AppEnv) Site(_ context.Context) string {
	return a.site
}

// RoleModel .
func (a *AppEnv) RoleModel(_ context.Context) string {
	return a.roleModel
}

// FrontSite .
func (a *AppEnv) FrontSite(_ context.Context) string {
	return a.frontSite
}

// String
func (a *AppEnv) String(_ context.Context) string {
	return `{"env":"` + a.env + `","environment":"` + a.environment + `","version":"` + a.version +
		`","jaegerEndpoint":"` + a.jaegerEndpoint + `","hostIP":"` + a.hostIP +
		`","uploadPath":"` + a.uploadPath + `","visitPath":"` + a.visitPath +
		`","site":"` + a.site + `","roleModel":"` + a.roleModel + `","frontSite":"` + a.frontSite + `"}`
}
