// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortUrls is the golang structure for table short_urls.
type ShortUrls struct {
	Id            uint64      `json:"id"            description:"ID"`
	AccountNo     uint64      `json:"accountNo"     description:"所属用户表示"`
	ShortNo       uint64      `json:"shortNo"       description:"短链的唯一 ID"`
	ShortUrl      string      `json:"shortUrl"      description:"短链内容"`
	ShortDomain   string      `json:"shortDomain"   description:"短链域名"`
	ShortDomainNo uint64      `json:"shortDomainNo" description:"短链域名标识"`
	DestUrl       string      `json:"destUrl"       description:"原始链接"`
	Domain        string      `json:"domain"        description:"主域名"`
	IsValid       uint        `json:"isValid"       description:"是否可用 0 默认 100 正常 200 失效"`
	DisableTime   *gtime.Time `json:"disableTime"   description:"失效时间"`
	Memo          string      `json:"memo"          description:"备注"`
	RawState      uint        `json:"rawState"      description:"原始状态 0 默认 100 附带额外参数"`
	Sort          uint64      `json:"sort"          description:"排序字段"`
	CollectState  uint        `json:"collectState"  description:"镜像采集状态 0 默认 100 已采集 200 采集失败"`
	CollectTime   *gtime.Time `json:"collectTime"   description:"采集时间"`
	DelState      uint        `json:"delState"      description:"是否删除 0 默认 200 删除"`
	DelTime       *gtime.Time `json:"delTime"       description:"删除时间"`
	CreateTime    *gtime.Time `json:"createTime"    description:"创建时间"`
	ModifyTime    *gtime.Time `json:"modifyTime"    description:"修改时间"`
}
