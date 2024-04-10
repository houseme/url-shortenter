// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortMirror is the golang structure for table short_mirror.
type ShortMirror struct {
	Id             uint64      `json:"id"             orm:"id"              description:"ID"`
	ShortNo        uint64      `json:"shortNo"        orm:"short_no"        description:"短链标识"`
	DestUrl        string      `json:"destUrl"        orm:"dest_url"        description:"原始 URL"`
	FullScreenshot string      `json:"fullScreenshot" orm:"full_screenshot" description:"整屏镜像"`
	ContentPath    string      `json:"contentPath"    orm:"content_path"    description:"内容文件地址"`
	Content        string      `json:"content"        orm:"content"         description:"网页内容"`
	HashContent    string      `json:"hashContent"    orm:"hash_content"    description:"hash 值 sha256"`
	CreateTime     *gtime.Time `json:"createTime"     orm:"create_time"     description:"创建时间"`
	ModifyTime     *gtime.Time `json:"modifyTime"     orm:"modify_time"     description:"修改时间"`
}
