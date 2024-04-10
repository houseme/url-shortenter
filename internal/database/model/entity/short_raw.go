// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ShortRaw is the golang structure for table short_raw.
type ShortRaw struct {
	Id         uint64      `json:"id"         orm:"id"          description:"ID"`
	ShortNo    uint64      `json:"shortNo"    orm:"short_no"    description:"链接标识"`
	RawNo      uint64      `json:"rawNo"      orm:"raw_no"      description:"唯一标识"`
	ShortRaw   string      `json:"shortRaw"   orm:"short_raw"   description:"额外参数标识"`
	ShortValue string      `json:"shortValue" orm:"short_value" description:"额外参数对应 value"`
	ShortKey   string      `json:"shortKey"   orm:"short_key"   description:"链接参数 key"`
	State      int         `json:"state"      orm:"state"       description:"0 默认 100 正常 200 失效"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`
	ModifyTime *gtime.Time `json:"modifyTime" orm:"modify_time" description:"更新时间"`
}
