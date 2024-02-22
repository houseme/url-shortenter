// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UsersDomain is the golang structure for table users_domain.
type UsersDomain struct {
	Id          uint64      `json:"id"          description:"ID"`
	AccountNo   uint64      `json:"accountNo"   description:"账号标识"`
	DomainNo    uint64      `json:"domainNo"    description:"域名标识"`
	Domain      string      `json:"domain"      description:"域名 不需要 http 等协议信息"`
	Memo        string      `json:"memo"        description:"备注信息"`
	License     string      `json:"license"     description:"icp 备案号"`
	SubLicense  string      `json:"subLicense"  description:"icp 备案号 带序号"`
	State       uint        `json:"state"       description:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	CreateTime  *gtime.Time `json:"createTime"  description:"创建时间"`
	AuditTime   *gtime.Time `json:"auditTime"   description:"审核时间"`
	DisableTime *gtime.Time `json:"disableTime" description:"禁用时间"`
	ModifyTime  *gtime.Time `json:"modifyTime"  description:"修改时间"`
}
