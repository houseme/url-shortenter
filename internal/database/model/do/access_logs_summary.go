// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AccessLogsSummary is the golang structure of table access_logs_summary for DAO operations like Where/Data.
type AccessLogsSummary struct {
	g.Meta     `orm:"table:access_logs_summary, do:true"`
	Id         interface{} // ID
	AccountNo  interface{} // 账号标识
	ShortNo    interface{} // 短链标识
	ShortUrl   interface{} // 短链内容
	ShortAll   interface{} // 带参数 URL
	YearTime   interface{} // 年份
	MonthTime  interface{} // 月份
	DayTime    interface{} // 日期
	AccessDate *gtime.Time // 访问日期
	UserAgent  interface{} // 访问 user_agent
	Ip         interface{} // 访问 IP
	Summary    interface{} // 访问汇总
	SuccessSum interface{} // 成功
	FailSum    interface{} // 失败
	CreateTime *gtime.Time // 创建时间
	ModifyTime *gtime.Time // 修改时间
}
