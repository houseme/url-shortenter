// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortDomain is the golang structure for table short_domain.
type ShortDomain struct {
	Id         uint64      `json:"id"         description:"ID"`
	DomainNo   uint64      `json:"domainNo"   description:"域名编号"`
	Domain     string      `json:"domain"     description:"短链域名"`
	Memo       string      `json:"memo"       description:"备注"`
	State      uint        `json:"state"      description:"状态  0 默认 100 正常 200 失效"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	ModifyTime *gtime.Time `json:"modifyTime" description:"更新时间"`
}
