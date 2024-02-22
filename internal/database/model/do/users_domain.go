// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UsersDomain is the golang structure of table users_domain for DAO operations like Where/Data.
type UsersDomain struct {
	g.Meta      `orm:"table:users_domain, do:true"`
	Id          interface{} // ID
	AccountNo   interface{} // 账号标识
	DomainNo    interface{} // 域名标识
	Domain      interface{} // 域名 不需要 http 等协议信息
	Memo        interface{} // 备注信息
	License     interface{} // icp 备案号
	SubLicense  interface{} // icp 备案号 带序号
	State       interface{} // 状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用
	CreateTime  *gtime.Time // 创建时间
	AuditTime   *gtime.Time // 审核时间
	DisableTime *gtime.Time // 禁用时间
	ModifyTime  *gtime.Time // 修改时间
}
