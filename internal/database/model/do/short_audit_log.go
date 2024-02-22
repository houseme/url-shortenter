// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortAuditLog is the golang structure of table short_audit_log for DAO operations like Where/Data.
type ShortAuditLog struct {
	g.Meta             `orm:"table:short_audit_log, do:true"`
	Id                 interface{} // ID
	ShortNo            interface{} // 短链标识
	TrxId              interface{} // 记录 ID
	FullScreenshot     interface{} // 整屏镜像
	Content            interface{} // 网页内容
	HashContent        interface{} // hash 值 sha256
	ContentPath        interface{} // 内容文件地址
	SafetyAuditAlibaba interface{} // 阿里内容安全审核
	SafetyAuditTencent interface{} // 腾讯内容审核
	AuditState         interface{} // 审核状态 0 默认 10000 阿里异常，20000 腾讯异常，100000 俩者都异常
	HashState          interface{} // hash 状态 0 默认 100 正常 200 失效
	RedirectState      interface{} // 跳转状态 0 默认 100 正常 200 异常
	ModifyTime         *gtime.Time // 修改时间
	CreateTime         *gtime.Time // 创建时间
}
