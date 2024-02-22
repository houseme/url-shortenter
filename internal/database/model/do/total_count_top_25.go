// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TotalCountTop25 is the golang structure of table total_count_top25 for DAO operations like Where/Data.
type TotalCountTop25 struct {
	g.Meta          `orm:"table:total_count_top25, do:true"`
	ShortUrl        interface{} // 短链内容
	TodayCount      interface{} //
	YesterdayCount  interface{} //
	Last7DaysCount  interface{} //
	MonthlyCount    interface{} //
	TotalCount      interface{} //
	DTodayCount     interface{} //
	DYesterdayCount interface{} //
	DLast7DaysCount interface{} //
	DMonthlyCount   interface{} //
	DTotalCount     interface{} //
	Id              interface{} // ID
	DestUrl         interface{} // 原始链接
	CreateTime      *gtime.Time // 创建时间
	IsValid         interface{} // 是否可用 0 默认 100 正常 200 失效
	Memo            interface{} // 备注
}
