// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortAuditLog is the golang structure for table short_audit_log.
type ShortAuditLog struct {
	Id                 uint64      `json:"id"                 orm:"id"                   description:"ID"`
	ShortNo            uint64      `json:"shortNo"            orm:"short_no"             description:"短链标识"`
	TrxId              uint64      `json:"trxId"              orm:"trx_id"               description:"记录 ID"`
	FullScreenshot     string      `json:"fullScreenshot"     orm:"full_screenshot"      description:"整屏镜像"`
	Content            string      `json:"content"            orm:"content"              description:"网页内容"`
	HashContent        string      `json:"hashContent"        orm:"hash_content"         description:"hash 值 sha256"`
	ContentPath        string      `json:"contentPath"        orm:"content_path"         description:"内容文件地址"`
	SafetyAuditAlibaba string      `json:"safetyAuditAlibaba" orm:"safety_audit_alibaba" description:"阿里内容安全审核"`
	SafetyAuditTencent string      `json:"safetyAuditTencent" orm:"safety_audit_tencent" description:"腾讯内容审核"`
	AuditState         uint        `json:"auditState"         orm:"audit_state"          description:"审核状态 0 默认 10000 阿里异常，20000 腾讯异常，100000 俩者都异常"`
	HashState          uint        `json:"hashState"          orm:"hash_state"           description:"hash 状态 0 默认 100 正常 200 失效"`
	RedirectState      uint        `json:"redirectState"      orm:"redirect_state"       description:"跳转状态 0 默认 100 正常 200 异常"`
	ModifyTime         *gtime.Time `json:"modifyTime"         orm:"modify_time"          description:"修改时间"`
	CreateTime         *gtime.Time `json:"createTime"         orm:"create_time"          description:"创建时间"`
}
