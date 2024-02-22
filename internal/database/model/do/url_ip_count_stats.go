// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UrlIpCountStats is the golang structure of table url_ip_count_stats for DAO operations like Where/Data.
type UrlIpCountStats struct {
	g.Meta          `orm:"table:url_ip_count_stats, do:true"`
	ShortUrl        interface{} // 短链内容
	ShortNo         interface{} // 短链的唯一 ID
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
}
