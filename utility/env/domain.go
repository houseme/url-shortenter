package env

import (
	"context"
)

// AppEnv .
type AppEnv struct {
	env            string                 `json:"env"`
	environment    string                 `json:"environment"`
	version        string                 `json:"version"`
	jaegerEndpoint string                 `json:"jaegerEndpoint"`
	hostIP         string                 `json:"hostIp"`
	config         map[string]interface{} `json:"config"`
	uploadPath     string                 `json:"uploadPath"`
	visitPath      string                 `json:"visitPath"`
	site           string                 `json:"site"`
	roleModel      string                 `json:"roleModel"`
	frontSite      string                 `json:"frontSite"`
}

// Env .
func (a *AppEnv) Env(ctx context.Context) string {
	return a.env
}

// Environment .
func (a *AppEnv) Environment(ctx context.Context) string {
	return a.environment
}

// Version .
func (a *AppEnv) Version(ctx context.Context) string {
	return a.version
}

// JaegerEndpoint .
func (a *AppEnv) JaegerEndpoint(ctx context.Context) string {
	return a.jaegerEndpoint
}

// Config .获取配置信息
func (a *AppEnv) Config(ctx context.Context) map[string]interface{} {
	return a.config
}

// HostIP . 获取本机IP
func (a *AppEnv) HostIP(ctx context.Context) string {
	return a.hostIP
}

// UploadPath .	上传路径
func (a *AppEnv) UploadPath(ctx context.Context) string {
	return a.uploadPath
}

// VisitPath file server访问路径
func (a *AppEnv) VisitPath(ctx context.Context) string {
	return a.visitPath
}

// Site .网站名称
func (a *AppEnv) Site(ctx context.Context) string {
	return a.site
}

// RoleModel .
func (a *AppEnv) RoleModel(ctx context.Context) string {
	return a.roleModel
}

// FrontSite .
func (a *AppEnv) FrontSite(ctx context.Context) string {
	return a.frontSite
}
