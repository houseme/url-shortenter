// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortDomain is the golang structure for table short_domain.
type ShortDomain struct {
	Id         uint64      `json:"id"         orm:"id"          description:"ID"`
	DomainNo   uint64      `json:"domainNo"   orm:"domain_no"   description:"域名编号"`
	Domain     string      `json:"domain"     orm:"domain"      description:"短链域名"`
	Memo       string      `json:"memo"       orm:"memo"        description:"备注"`
	State      uint        `json:"state"      orm:"state"       description:"状态  0 默认 100 正常 200 失效"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`
	ModifyTime *gtime.Time `json:"modifyTime" orm:"modify_time" description:"更新时间"`
}
