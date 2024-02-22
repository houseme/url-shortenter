// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AccessLogs is the golang structure of table access_logs for DAO operations like Where/Data.
type AccessLogs struct {
	g.Meta     `orm:"table:access_logs, do:true"`
	Id         interface{} // ID
	AccountNo  interface{} // 账号标识
	ShortNo    interface{} // 短链标识
	ShortUrl   interface{} // 短链内容
	ShortAll   interface{} // 带参数 URL
	YearTime   interface{} // 年份
	MonthTime  interface{} // 月份
	DayTime    interface{} // 日期
	AccessDate *gtime.Time // 访问日期
	AccessTime *gtime.Time // 访问时间
	UserAgent  interface{} // 访问 user_agent
	Ip         interface{} // 访问 IP
	TraceId    interface{} // 链路追踪标识
	VisitState interface{} // 访问状态 0 默认，100 正常 200 失效
	ServerIp   interface{} // 服务器 IP
	CreateTime *gtime.Time // 创建时间
	ModifyTime *gtime.Time // 修改时间
}
