package env

import (
	"context"
)

// AppEnv .
type appEnv struct {
	env            string                 `json:"env"`
	environment    string                 `json:"environment"`
	version        string                 `json:"version"`
	datacenterID   int64                  `json:"datacenterID"`
	workerID       int64                  `json:"workerID"`
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
func (a *appEnv) Env(ctx context.Context) string {
	return a.env
}

// Environment .
func (a *appEnv) Environment(ctx context.Context) string {
	return a.environment
}

// Version .
func (a *appEnv) Version(ctx context.Context) string {
	return a.version
}

// DatacenterID .
func (a *appEnv) DatacenterID(ctx context.Context) int64 {
	return a.datacenterID
}

// WorkerID .工作机器ID
func (a *appEnv) WorkerID(ctx context.Context) int64 {
	return a.workerID
}

// JaegerEndpoint .
func (a *appEnv) JaegerEndpoint(ctx context.Context) string {
	return a.jaegerEndpoint
}

// Config .获取配置信息
func (a *appEnv) Config(ctx context.Context) map[string]interface{} {
	return a.config
}

// HostIP . 获取本机IP
func (a *appEnv) HostIP(ctx context.Context) string {
	return a.hostIP
}

// UploadPath .	上传路径
func (a *appEnv) UploadPath(ctx context.Context) string {
	return a.uploadPath
}

// VisitPath file server访问路径
func (a *appEnv) VisitPath(ctx context.Context) string {
	return a.visitPath
}

// Site .网站名称
func (a *appEnv) Site(ctx context.Context) string {
	return a.site
}

// RoleModel .
func (a *appEnv) RoleModel(ctx context.Context) string {
	return a.roleModel
}

// FrontSite .
func (a *appEnv) FrontSite(ctx context.Context) string {
	return a.frontSite
}
