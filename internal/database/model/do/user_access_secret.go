// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserAccessSecret is the golang structure of table user_access_secret for DAO operations like Where/Data.
type UserAccessSecret struct {
	g.Meta     `orm:"table:user_access_secret, do:true"`
	Id         interface{} // ID
	AccountNo  interface{} // 企业管理员 ID 一致
	SecretId   interface{} // 授权应用 ID
	SecretKey  interface{} // 授权应用 key
	Salt       interface{} // 盐值
	SaltKey    interface{} // 盐值 key
	GrantType  interface{} // 授权类型：默认空，API token 授权:client_credentials
	State      interface{} // 状态 0 默认 100 正常 200 禁用
	CreateTime *gtime.Time // 创建时间
	ModifyTime *gtime.Time // 修改时间
}
