// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	Id         uint64      `json:"id"         description:"ID"`
	UserNo     uint64      `json:"userNo"     description:"账号唯一标识"`
	Account    string      `json:"account"    description:"账号"`
	Password   string      `json:"password"   description:"密码"`
	State      uint        `json:"state"      description:"状态 0 默认 100正常 200失效"`
	GroupLevel uint        `json:"groupLevel" description:"用户等级 0 默认超级 1000商户管理员，10000普通管理员"`
	AccountNo  uint64      `json:"accountNo"  description:"用户关联企业ID 同企业管理员ID一致"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	ModifyTime *gtime.Time `json:"modifyTime" description:"更新时间"`
}
