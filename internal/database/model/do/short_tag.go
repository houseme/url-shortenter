// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortTag is the golang structure of table short_tag for DAO operations like Where/Data.
type ShortTag struct {
	g.Meta     `orm:"table:short_tag, do:true"`
	Id         interface{} // ID
	TagNo      interface{} // 唯一标识
	TagName    interface{} // 标签名称
	TagPinyin  interface{} // 名称拼音 中划线链接
	TagEn      interface{} // 名称英文 中划线链接
	State      interface{} // 状态 0 默认 100 正常 200 失效
	CreateTime *gtime.Time // 创建时间
	ModifyTime *gtime.Time // 修改时间
}
