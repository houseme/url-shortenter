// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UrlIpCountStats is the golang structure for table url_ip_count_stats.
type UrlIpCountStats struct {
	ShortUrl        string `json:"shortUrl"        orm:"short_url"           description:"短链内容"`
	ShortNo         uint64 `json:"shortNo"         orm:"short_no"            description:"短链的唯一 ID"`
	TodayCount      int64  `json:"todayCount"      orm:"today_count"         description:""`
	YesterdayCount  int64  `json:"yesterdayCount"  orm:"yesterday_count"     description:""`
	Last7DaysCount  int64  `json:"last7DaysCount"  orm:"last_7_days_count"   description:""`
	MonthlyCount    int64  `json:"monthlyCount"    orm:"monthly_count"       description:""`
	TotalCount      int64  `json:"totalCount"      orm:"total_count"         description:""`
	DTodayCount     int64  `json:"dTodayCount"     orm:"d_today_count"       description:""`
	DYesterdayCount int64  `json:"dYesterdayCount" orm:"d_yesterday_count"   description:""`
	DLast7DaysCount int64  `json:"dLast7DaysCount" orm:"d_last_7_days_count" description:""`
	DMonthlyCount   int64  `json:"dMonthlyCount"   orm:"d_monthly_count"     description:""`
	DTotalCount     int64  `json:"dTotalCount"     orm:"d_total_count"       description:""`
}
