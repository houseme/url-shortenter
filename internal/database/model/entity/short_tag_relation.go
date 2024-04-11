// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortTagRelation is the golang structure for table short_tag_relation.
type ShortTagRelation struct {
	Id         uint64      `json:"id"         orm:"id"          description:"ID"`
	UserNo     uint64      `json:"userNo"     orm:"user_no"     description:"用户标识"`
	TagNo      uint64      `json:"tagNo"      orm:"tag_no"      description:"标签标识"`
	ShortNo    uint64      `json:"shortNo"    orm:"short_no"    description:"短链的唯一 ID"`
	State      uint        `json:"state"      orm:"state"       description:"状态 0 默认 100 正常 200 失效"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`
	ModifyTime *gtime.Time `json:"modifyTime" orm:"modify_time" description:"更新时间"`
}
