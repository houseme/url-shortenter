// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id         uint64      `json:"id"         orm:"id"          description:"ID"`
	UserNo     uint64      `json:"userNo"     orm:"user_no"     description:"账号唯一标识"`
	Username   string      `json:"username"   orm:"username"    description:"用户名称"`
	Account    string      `json:"account"    orm:"account"     description:"账号"`
	Password   string      `json:"password"   orm:"password"    description:"密码"`
	State      uint        `json:"state"      orm:"state"       description:"状态 0 默认 100 正常 200 失效"`
	Avatar     string      `json:"avatar"     orm:"avatar"      description:"头像"`
	GroupLevel uint        `json:"groupLevel" orm:"group_level" description:"用户等级 0 默认超级 1000 企业管理员，10000 普通管理员"`
	AccountNo  uint64      `json:"accountNo"  orm:"account_no"  description:"用户关联企业 ID 同企业管理员 ID 一致"`
	Deadline   uint64      `json:"deadline"   orm:"deadline"    description:"截止时间"`
	Salt       string      `json:"salt"       orm:"salt"        description:"盐值"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`
	ModifyTime *gtime.Time `json:"modifyTime" orm:"modify_time" description:"更新时间"`
}
