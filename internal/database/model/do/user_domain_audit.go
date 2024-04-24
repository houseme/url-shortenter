// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserDomainAudit is the golang structure of table user_domain_audit for DAO operations like Where/Data.
type UserDomainAudit struct {
	g.Meta         `orm:"table:user_domain_audit, do:true"`
	Id             interface{} // ID
	AccountNo      interface{} // 用户标识
	UserNo         interface{} // 用户标识
	DomainNo       interface{} // 用户认证域名 ID
	Icp            interface{} // icp 备案号
	QueryResult    interface{} // 查询结果
	AuditAccountNo interface{} // 审核用户 ID
	AuditTime      *gtime.Time // 审核时间
	AuditState     interface{} // 审核状态 0 默认 20:审核通过 30:审核失败
	CreateTime     *gtime.Time // 创建时间
	ModifyTime     *gtime.Time // 修改时间
}
