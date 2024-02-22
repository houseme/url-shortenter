// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortDomain is the golang structure of table short_domain for DAO operations like Where/Data.
type ShortDomain struct {
	g.Meta     `orm:"table:short_domain, do:true"`
	Id         interface{} // ID
	DomainNo   interface{} // 域名编号
	Domain     interface{} // 短链域名
	Memo       interface{} // 备注
	State      interface{} // 状态  0 默认 100 正常 200 失效
	CreateTime *gtime.Time // 创建时间
	ModifyTime *gtime.Time // 更新时间
}
