// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta     `orm:"table:users, do:true"`
	Id         interface{} // ID
	UserNo     interface{} // 账号唯一标识
	Account    interface{} // 账号
	Password   interface{} // 密码
	State      interface{} // 状态 0 默认 100 正常 200 失效
	GroupLevel interface{} // 用户等级 0 默认超级 1000 商户管理员，10000 普通管理员
	AccountNo  interface{} // 用户关联企业 ID 同企业管理员 ID 一致
	CreateTime *gtime.Time // 创建时间
	ModifyTime *gtime.Time // 更新时间
}
