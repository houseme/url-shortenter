// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UsersDomainAudit is the golang structure for table users_domain_audit.
type UsersDomainAudit struct {
	Id             uint64      `json:"id"             description:"ID"`
	AccountNo      uint64      `json:"accountNo"      description:"用户标识"`
	DomainNo       uint64      `json:"domainNo"       description:"用户认证域名 ID"`
	Icp            string      `json:"icp"            description:"icp 备案号"`
	QueryResult    string      `json:"queryResult"    description:"查询结果"`
	AuditAccountNo uint64      `json:"auditAccountNo" description:"审核用户 ID"`
	AuditTime      *gtime.Time `json:"auditTime"      description:"审核时间"`
	AuditState     uint        `json:"auditState"     description:"审核状态 0 默认 20:审核通过 30:审核失败"`
	CreateTime     *gtime.Time `json:"createTime"     description:"创建时间"`
	ModifyTime     *gtime.Time `json:"modifyTime"     description:"修改时间"`
}
