// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserDomain is the golang structure for table user_domain.
type UserDomain struct {
	Id          uint64      `json:"id"          orm:"id"           description:"ID"`
	AccountNo   uint64      `json:"accountNo"   orm:"account_no"   description:"账号标识"`
	UserNo      uint64      `json:"userNo"      orm:"user_no"      description:"用户标识"`
	DomainNo    uint64      `json:"domainNo"    orm:"domain_no"    description:"域名标识"`
	Domain      string      `json:"domain"      orm:"domain"       description:"域名 不需要 http 等协议信息"`
	Memo        string      `json:"memo"        orm:"memo"         description:"备注信息"`
	License     string      `json:"license"     orm:"license"      description:"icp 备案号"`
	SubLicense  string      `json:"subLicense"  orm:"sub_license"  description:"icp 备案号 带序号"`
	State       uint        `json:"state"       orm:"state"        description:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"  description:"创建时间"`
	AuditTime   *gtime.Time `json:"auditTime"   orm:"audit_time"   description:"审核时间"`
	DisableTime *gtime.Time `json:"disableTime" orm:"disable_time" description:"禁用时间"`
	ModifyTime  *gtime.Time `json:"modifyTime"  orm:"modify_time"  description:"修改时间"`
}
