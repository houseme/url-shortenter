// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TotalCountTop25 is the golang structure for table total_count_top25.
type TotalCountTop25 struct {
	ShortUrl        string      `json:"shortUrl"        description:"短链内容"`
	TodayCount      int64       `json:"todayCount"      description:""`
	YesterdayCount  int64       `json:"yesterdayCount"  description:""`
	Last7DaysCount  int64       `json:"last7DaysCount"  description:""`
	MonthlyCount    int64       `json:"monthlyCount"    description:""`
	TotalCount      int64       `json:"totalCount"      description:""`
	DTodayCount     int64       `json:"dTodayCount"     description:""`
	DYesterdayCount int64       `json:"dYesterdayCount" description:""`
	DLast7DaysCount int64       `json:"dLast7DaysCount" description:""`
	DMonthlyCount   int64       `json:"dMonthlyCount"   description:""`
	DTotalCount     int64       `json:"dTotalCount"     description:""`
	Id              uint64      `json:"id"              description:"ID"`
	DestUrl         string      `json:"destUrl"         description:"原始链接"`
	CreateTime      *gtime.Time `json:"createTime"      description:"创建时间"`
	IsValid         uint        `json:"isValid"         description:"是否可用 0 默认 100 正常 200 失效"`
	Memo            string      `json:"memo"            description:"备注"`
}
