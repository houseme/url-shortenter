// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserAccessSecret is the golang structure for table user_access_secret.
type UserAccessSecret struct {
	Id         uint64      `json:"id"         orm:"id"          description:"ID"`
	AccountNo  uint64      `json:"accountNo"  orm:"account_no"  description:"企业管理员 ID 一致"`
	SecretId   string      `json:"secretId"   orm:"secret_id"   description:"授权应用 ID"`
	SecretKey  string      `json:"secretKey"  orm:"secret_key"  description:"授权应用 key"`
	Salt       string      `json:"salt"       orm:"salt"        description:"盐值"`
	SaltKey    string      `json:"saltKey"    orm:"salt_key"    description:"盐值 key"`
	GrantType  string      `json:"grantType"  orm:"grant_type"  description:"授权类型：默认空，API token 授权:client_credentials"`
	State      uint        `json:"state"      orm:"state"       description:"状态 0 默认 100 正常 200 禁用"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`
	ModifyTime *gtime.Time `json:"modifyTime" orm:"modify_time" description:"修改时间"`
}
