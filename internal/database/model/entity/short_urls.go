// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortUrls is the golang structure for table short_urls.
type ShortUrls struct {
	Id            uint64      `json:"id"            orm:"id"              description:"ID"`
	AccountNo     uint64      `json:"accountNo"     orm:"account_no"      description:"所属用户表示"`
	ShortNo       uint64      `json:"shortNo"       orm:"short_no"        description:"短链的唯一 ID"`
	ShortUrl      string      `json:"shortUrl"      orm:"short_url"       description:"短链内容"`
	ShortDomain   string      `json:"shortDomain"   orm:"short_domain"    description:"短链域名"`
	ShortDomainNo uint64      `json:"shortDomainNo" orm:"short_domain_no" description:"短链域名标识"`
	DestUrl       string      `json:"destUrl"       orm:"dest_url"        description:"原始链接"`
	Domain        string      `json:"domain"        orm:"domain"          description:"主域名"`
	IsValid       uint        `json:"isValid"       orm:"is_valid"        description:"是否可用 0 默认 100 正常 200 失效"`
	DisableTime   *gtime.Time `json:"disableTime"   orm:"disable_time"    description:"失效时间"`
	Memo          string      `json:"memo"          orm:"memo"            description:"备注"`
	RawState      uint        `json:"rawState"      orm:"raw_state"       description:"原始状态 0 默认 100 附带额外参数"`
	Sort          uint64      `json:"sort"          orm:"sort"            description:"排序字段"`
	CollectState  uint        `json:"collectState"  orm:"collect_state"   description:"镜像采集状态 0 默认 100 已采集 200 采集失败"`
	CollectTime   *gtime.Time `json:"collectTime"   orm:"collect_time"    description:"采集时间"`
	DelState      uint        `json:"delState"      orm:"del_state"       description:"是否删除 0 默认 200 删除"`
	DelTime       *gtime.Time `json:"delTime"       orm:"del_time"        description:"删除时间"`
	CreateTime    *gtime.Time `json:"createTime"    orm:"create_time"     description:"创建时间"`
	ModifyTime    *gtime.Time `json:"modifyTime"    orm:"modify_time"     description:"修改时间"`
}
