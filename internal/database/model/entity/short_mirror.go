// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortMirror is the golang structure for table short_mirror.
type ShortMirror struct {
	Id             uint64      `json:"id"             description:"ID"`
	ShortNo        uint64      `json:"shortNo"        description:"短链标识"`
	DestUrl        string      `json:"destUrl"        description:"原始 URL"`
	FullScreenshot string      `json:"fullScreenshot" description:"整屏镜像"`
	ContentPath    string      `json:"contentPath"    description:"内容文件地址"`
	Content        string      `json:"content"        description:"网页内容"`
	HashContent    string      `json:"hashContent"    description:"hash 值 sha256"`
	CreateTime     *gtime.Time `json:"createTime"     description:"创建时间"`
	ModifyTime     *gtime.Time `json:"modifyTime"     description:"修改时间"`
}
