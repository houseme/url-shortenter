// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortContentRecord is the golang structure of table short_content_record for DAO operations like Where/Data.
type ShortContentRecord struct {
	g.Meta      `orm:"table:short_content_record, do:true"`
	Id          interface{} // ID
	ShortNo     interface{} // 短链标识
	TrxId       interface{} // 唯一标识
	ContentType interface{} // 内容类型 0 默认，100 镜像内容 200 审核内容
	Content     interface{} // 网页内容
	HashContent interface{} // hash 值 sha256
	CreateTime  *gtime.Time // 创建时间
	ModifyTime  *gtime.Time // 修改时间
}
