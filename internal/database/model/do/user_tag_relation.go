// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserTagRelation is the golang structure of table user_tag_relation for DAO operations like Where/Data.
type UserTagRelation struct {
	g.Meta     `orm:"table:user_tag_relation, do:true"`
	Id         interface{} // ID
	UserNo     interface{} // 用户标识
	TagNo      interface{} // 标签标识
	State      interface{} // 状态 0 默认 100 正常 200 失效
	CreateTime *gtime.Time // 创建时间
	ModifyTime *gtime.Time // 更新时间
}
