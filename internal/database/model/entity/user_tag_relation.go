// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserTagRelation is the golang structure for table user_tag_relation.
type UserTagRelation struct {
	Id         uint64      `json:"id"         orm:"id"          description:"ID"`
	UserNo     uint64      `json:"userNo"     orm:"user_no"     description:"用户标识"`
	TagNo      uint64      `json:"tagNo"      orm:"tag_no"      description:"标签标识"`
	State      uint        `json:"state"      orm:"state"       description:"状态 0 默认 100 正常 200 失效"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`
	ModifyTime *gtime.Time `json:"modifyTime" orm:"modify_time" description:"更新时间"`
}
