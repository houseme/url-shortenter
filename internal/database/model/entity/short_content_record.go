// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortContentRecord is the golang structure for table short_content_record.
type ShortContentRecord struct {
	Id          uint64      `json:"id"          orm:"id"           description:"ID"`
	ShortNo     uint64      `json:"shortNo"     orm:"short_no"     description:"短链标识"`
	TrxId       uint64      `json:"trxId"       orm:"trx_id"       description:"唯一标识"`
	ContentType uint        `json:"contentType" orm:"content_type" description:"内容类型 0 默认，100 镜像内容 200 审核内容"`
	Content     string      `json:"content"     orm:"content"      description:"网页内容"`
	HashContent string      `json:"hashContent" orm:"hash_content" description:"hash 值 sha256"`
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time"  description:"创建时间"`
	ModifyTime  *gtime.Time `json:"modifyTime"  orm:"modify_time"  description:"修改时间"`
}
