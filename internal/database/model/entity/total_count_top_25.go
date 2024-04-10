// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TotalCountTop25 is the golang structure for table total_count_top25.
type TotalCountTop25 struct {
	ShortUrl        string      `json:"shortUrl"        orm:"short_url"           description:"短链内容"`
	TodayCount      int64       `json:"todayCount"      orm:"today_count"         description:""`
	YesterdayCount  int64       `json:"yesterdayCount"  orm:"yesterday_count"     description:""`
	Last7DaysCount  int64       `json:"last7DaysCount"  orm:"last_7_days_count"   description:""`
	MonthlyCount    int64       `json:"monthlyCount"    orm:"monthly_count"       description:""`
	TotalCount      int64       `json:"totalCount"      orm:"total_count"         description:""`
	DTodayCount     int64       `json:"dTodayCount"     orm:"d_today_count"       description:""`
	DYesterdayCount int64       `json:"dYesterdayCount" orm:"d_yesterday_count"   description:""`
	DLast7DaysCount int64       `json:"dLast7DaysCount" orm:"d_last_7_days_count" description:""`
	DMonthlyCount   int64       `json:"dMonthlyCount"   orm:"d_monthly_count"     description:""`
	DTotalCount     int64       `json:"dTotalCount"     orm:"d_total_count"       description:""`
	Id              uint64      `json:"id"              orm:"id"                  description:"ID"`
	DestUrl         string      `json:"destUrl"         orm:"dest_url"            description:"原始链接"`
	CreateTime      *gtime.Time `json:"createTime"      orm:"create_time"         description:"创建时间"`
	IsValid         uint        `json:"isValid"         orm:"is_valid"            description:"是否可用 0 默认 100 正常 200 失效"`
	Memo            string      `json:"memo"            orm:"memo"                description:"备注"`
}
