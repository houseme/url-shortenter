// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UrlIpCountStats is the golang structure for table url_ip_count_stats.
type UrlIpCountStats struct {
	ShortUrl        string `json:"shortUrl"        description:"短链内容"`
	ShortNo         uint64 `json:"shortNo"         description:"短链的唯一 ID"`
	TodayCount      int64  `json:"todayCount"      description:""`
	YesterdayCount  int64  `json:"yesterdayCount"  description:""`
	Last7DaysCount  int64  `json:"last7DaysCount"  description:""`
	MonthlyCount    int64  `json:"monthlyCount"    description:""`
	TotalCount      int64  `json:"totalCount"      description:""`
	DTodayCount     int64  `json:"dTodayCount"     description:""`
	DYesterdayCount int64  `json:"dYesterdayCount" description:""`
	DLast7DaysCount int64  `json:"dLast7DaysCount" description:""`
	DMonthlyCount   int64  `json:"dMonthlyCount"   description:""`
	DTotalCount     int64  `json:"dTotalCount"     description:""`
}
