// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AccessLogs is the golang structure for table access_logs.
type AccessLogs struct {
	Id         uint64      `json:"id"         orm:"id"          description:"ID"`
	AccountNo  uint64      `json:"accountNo"  orm:"account_no"  description:"账号标识"`
	ShortNo    uint64      `json:"shortNo"    orm:"short_no"    description:"短链标识"`
	ShortUrl   string      `json:"shortUrl"   orm:"short_url"   description:"短链内容"`
	ShortAll   string      `json:"shortAll"   orm:"short_all"   description:"带参数 URL"`
	YearTime   uint        `json:"yearTime"   orm:"year_time"   description:"年份"`
	MonthTime  uint        `json:"monthTime"  orm:"month_time"  description:"月份"`
	DayTime    uint        `json:"dayTime"    orm:"day_time"    description:"日期"`
	AccessDate *gtime.Time `json:"accessDate" orm:"access_date" description:"访问日期"`
	AccessTime *gtime.Time `json:"accessTime" orm:"access_time" description:"访问时间"`
	UserAgent  string      `json:"userAgent"  orm:"user_agent"  description:"访问 user_agent"`
	Ip         string      `json:"ip"         orm:"ip"          description:"访问 IP"`
	TraceId    string      `json:"traceId"    orm:"trace_id"    description:"链路追踪标识"`
	VisitState uint        `json:"visitState" orm:"visit_state" description:"访问状态 0 默认，100 正常 200 失效"`
	ServerIp   string      `json:"serverIp"   orm:"server_ip"   description:"服务器 IP"`
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`
	ModifyTime *gtime.Time `json:"modifyTime" orm:"modify_time" description:"修改时间"`
}
