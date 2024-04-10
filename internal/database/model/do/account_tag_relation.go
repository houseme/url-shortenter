// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AccountTagRelation is the golang structure of table account_tag_relation for DAO operations like Where/Data.
type AccountTagRelation struct {
	g.Meta     `orm:"table:account_tag_relation, do:true"`
	Id         interface{} // ID
	AccountNo  interface{} // 用户标识
	TagNo      interface{} // 标签标识
	State      interface{} // 状态 0 默认 100 正常 200 失效
	CreateTime *gtime.Time // 创建时间
	ModifyTime *gtime.Time // 更新时间
}
