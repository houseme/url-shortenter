// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortMirror is the golang structure of table short_mirror for DAO operations like Where/Data.
type ShortMirror struct {
	g.Meta         `orm:"table:short_mirror, do:true"`
	Id             interface{} // ID
	ShortNo        interface{} // 短链标识
	DestUrl        interface{} // 原始 URL
	FullScreenshot interface{} // 整屏镜像
	ContentPath    interface{} // 内容文件地址
	Content        interface{} // 网页内容
	HashContent    interface{} // hash 值 sha256
	CreateTime     *gtime.Time // 创建时间
	ModifyTime     *gtime.Time // 修改时间
}
