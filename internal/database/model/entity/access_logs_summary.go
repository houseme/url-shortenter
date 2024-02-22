// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AccessLogsSummary is the golang structure for table access_logs_summary.
type AccessLogsSummary struct {
	Id         uint64      `json:"id"         description:"ID"`
	AccountNo  uint64      `json:"accountNo"  description:"账号标识"`
	ShortNo    uint64      `json:"shortNo"    description:"短链标识"`
	ShortUrl   string      `json:"shortUrl"   description:"短链内容"`
	ShortAll   string      `json:"shortAll"   description:"带参数 URL"`
	YearTime   uint        `json:"yearTime"   description:"年份"`
	MonthTime  uint        `json:"monthTime"  description:"月份"`
	DayTime    uint        `json:"dayTime"    description:"日期"`
	AccessDate *gtime.Time `json:"accessDate" description:"访问日期"`
	UserAgent  string      `json:"userAgent"  description:"访问 user_agent"`
	Ip         string      `json:"ip"         description:"访问 IP"`
	Summary    uint64      `json:"summary"    description:"访问汇总"`
	SuccessSum uint64      `json:"successSum" description:"成功"`
	FailSum    uint64      `json:"failSum"    description:"失败"`
	CreateTime *gtime.Time `json:"createTime" description:"创建时间"`
	ModifyTime *gtime.Time `json:"modifyTime" description:"修改时间"`
}
