// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AccessLogsSummary is the golang structure for table access_logs_summary.
type AccessLogsSummary struct {
	Id         uint64      `json:"id"         orm:"id"          description:"ID"`
	SummaryNo  uint64      `json:"summaryNo"  orm:"summary_no"  description:"统计记录唯一标识"`
	AccountNo  uint64      `json:"accountNo"  orm:"account_no"  description:"账号标识"`
	UserNo     uint64      `json:"userNo"     orm:"user_no"     description:"用户标识"`
	ShortNo    uint64      `json:"shortNo"    orm:"short_no"    description:"短链标识"`
	ShortUrl   string      `json:"shortUrl"   orm:"short_url"   description:"短链内容"`
	ShortAll   string      `json:"shortAll"   orm:"short_all"   description:"带参数 URL"`
	YearTime   uint        `json:"yearTime"   orm:"year_time"   description:"年份"`
	MonthTime  uint        `json:"monthTime"  orm:"month_time"  description:"月份"`
	DayTime    uint        `json:"dayTime"    orm:"day_time"    description:"日期"`
	AccessDate *gtime.Time `json:"accessDate" orm:"access_date" description:"访问日期"`
	UserAgent  string      `json:"userAgent"  orm:"user_agent"  description:"访问 user_agent"`
	Ip         string      `json:"ip"         orm:"ip"          description:"访问 IP"`
	Summary    uint64      `json:"summary"    orm:"summary"     description:"访问汇总"`
	SuccessSum uint64      `json:"successSum" orm:"success_sum" description:"成功"`
	FailSum    uint64      `json:"failSum"    orm:"fail_sum"    description:"失败"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`
	ModifyTime *gtime.Time `json:"modifyTime" orm:"modify_time" description:"修改时间"`
}
