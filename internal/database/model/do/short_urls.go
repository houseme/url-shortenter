// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortUrls is the golang structure of table short_urls for DAO operations like Where/Data.
type ShortUrls struct {
	g.Meta        `orm:"table:short_urls, do:true"`
	Id            interface{} // ID
	AccountNo     interface{} // 所属用户表示
	UserNo        interface{} // 用户表示
	ShortNo       interface{} // 短链的唯一 ID
	ShortUrl      interface{} // 短链内容
	ShortDomain   interface{} // 短链域名
	ShortDomainNo interface{} // 短链域名标识
	DestUrl       interface{} // 原始链接
	Domain        interface{} // 主域名
	IsValid       interface{} // 是否可用 0 默认 100 正常 200 失效
	DisableTime   *gtime.Time // 失效时间
	Memo          interface{} // 备注
	RawState      interface{} // 原始状态 0 默认 100 附带额外参数
	Sort          interface{} // 排序字段
	CollectState  interface{} // 镜像采集状态 0 默认 100 已采集 200 采集失败
	CollectTime   *gtime.Time // 采集时间
	DelState      interface{} // 是否删除 0 默认 200 删除
	DelTime       *gtime.Time // 删除时间
	CreateTime    *gtime.Time // 创建时间
	ModifyTime    *gtime.Time // 修改时间
}
