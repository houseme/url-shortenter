// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortRaw is the golang structure of table short_raw for DAO operations like Where/Data.
type ShortRaw struct {
	g.Meta     `orm:"table:short_raw, do:true"`
	Id         interface{} // ID
	ShortNo    interface{} // 链接标识
	RawNo      interface{} // 唯一标识
	ShortRaw   interface{} // 额外参数标识
	ShortValue interface{} // 额外参数对应 value
	ShortKey   interface{} // 链接参数 key
	State      interface{} // 0 默认 100 正常 200 失效
	CreateTime *gtime.Time // 创建时间
	ModifyTime *gtime.Time // 更新时间
}
