// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserDomainAudit is the golang structure for table user_domain_audit.
type UserDomainAudit struct {
	Id             uint64      `json:"id"             orm:"id"               description:"ID"`
	AccountNo      uint64      `json:"accountNo"      orm:"account_no"       description:"用户标识"`
	UserNo         uint64      `json:"userNo"         orm:"user_no"          description:"用户标识"`
	DomainNo       uint64      `json:"domainNo"       orm:"domain_no"        description:"用户认证域名 ID"`
	Icp            string      `json:"icp"            orm:"icp"              description:"icp 备案号"`
	QueryResult    string      `json:"queryResult"    orm:"query_result"     description:"查询结果"`
	AuditAccountNo uint64      `json:"auditAccountNo" orm:"audit_account_no" description:"审核用户 ID"`
	AuditTime      *gtime.Time `json:"auditTime"      orm:"audit_time"       description:"审核时间"`
	AuditState     uint        `json:"auditState"     orm:"audit_state"      description:"审核状态 0 默认 20:审核通过 30:审核失败"`
	CreateTime     *gtime.Time `json:"createTime"     orm:"create_time"      description:"创建时间"`
	ModifyTime     *gtime.Time `json:"modifyTime"     orm:"modify_time"      description:"修改时间"`
}
