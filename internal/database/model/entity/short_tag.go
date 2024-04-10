// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortTag is the golang structure for table short_tag.
type ShortTag struct {
	Id         uint64      `json:"id"         orm:"id"          description:"ID"`
	TagNo      uint64      `json:"tagNo"      orm:"tag_no"      description:"唯一标识"`
	TagName    string      `json:"tagName"    orm:"tag_name"    description:"标签名称"`
	TagPinyin  string      `json:"tagPinyin"  orm:"tag_pinyin"  description:"名称拼音 中划线链接"`
	TagEn      string      `json:"tagEn"      orm:"tag_en"      description:"名称英文 中划线链接"`
	State      uint        `json:"state"      orm:"state"       description:"状态 0 默认 100 正常 200 失效"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`
	ModifyTime *gtime.Time `json:"modifyTime" orm:"modify_time" description:"修改时间"`
}
