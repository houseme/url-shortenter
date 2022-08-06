package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CreateShortInput is the input for creating a short
type CreateShortInput struct {
	*Base   `json:"-"`
	Memo    string `json:"memo" dc:"备注"`
	DestURL string `json:"destUrl" dc:"目标链接"`
}

// CreateShortOutput is the output of CreateShort
type CreateShortOutput struct {
	ShortURL string `json:"shortUrl" dc:"短链"`
}

// QueryShortInput is the input for QueryShort
type QueryShortInput struct {
	*Base    `json:"-"`
	ShortURL string `json:"shortUrl" dc:"短链"`
}

// QueryShortOutput is the output of QueryShort
type QueryShortOutput struct {
	DestURL     string      `json:"destUrl" dc:"destUrl 原始url"`
	Memo        string      `json:"memo,omitempty" dc:"备注"`
	State       uint        `json:"state" dc:"状态"`
	StateMsg    string      `json:"stateMsg" dc:"状态描述"`
	DisableTime *gtime.Time `json:"disableTime,omitempty" dc:"禁用时间"`
	CreateTime  *gtime.Time `json:"createTime" dc:"创建时间"`
}

// QueryStatInput is the input for QueryStat
type QueryStatInput struct {
	*Base    `json:"-"`
	ShortURL string `json:"shortUrl"`
}

// QueryStatOutput is the output of QueryStat
type QueryStatOutput struct {
	ShortURL                string      `json:"shortUrl" dc:"短链"`
	DestURL                 string      `json:"destUrl" dc:"原始url"`
	Memo                    string      `json:"memo,omitempty" dc:"备注"`
	State                   uint        `json:"state" dc:"状态"`
	StateMsg                string      `json:"stateMsg" dc:"状态描述"`
	CreateTime              *gtime.Time `json:"createTime" dc:"创建时间"`
	DisableTime             *gtime.Time `json:"disableTime,omitempty" dc:"禁用时间"`
	TodayCount              uint64      `json:"todayCount" dc:"今日访问次数"`
	YesterdayCount          uint64      `json:"yesterdayCount" dc:"昨日访问次数"`
	LastSevenDaysCount      uint64      `json:"lastSevenDaysCount" dc:"近7日访问次数"`
	MonthlyCount            uint64      `json:"monthlyCount" dc:"本月访问次数"`
	LastMonthlyCount        uint64      `json:"lastMonthlyCount" dc:"上月访问次数"`
	TotalCount              uint64      `json:"totalCount" dc:"总访问次数"`
	ProtoTodayCount         uint64      `json:"protoTodayCount" dc:"今日独立IP访问次数"`
	ProtoYesterdayCount     uint64      `json:"protoYesterdayCount" dc:"昨日独立IP访问次数"`
	ProtoLastSevenDaysCount uint64      `json:"protoLastSevenDaysCount" dc:"近7日独立IP访问次数"`
	ProtoMonthlyCount       uint64      `json:"protoMonthlyCount" dc:"本月独立IP访问次数"`
	ProtoLastMonthlyCount   uint64      `json:"protoLastMonthlyCount" dc:"上月独立IP访问次数"`
	ProtoTotalCount         uint64      `json:"protoTotalCount" dc:"总独立IP访问次数"`
}

// ModifyShortInput is the input for ModifyShort
type ModifyShortInput struct {
	*Base    `json:"-"`
	ShortURL string `json:"shortUrl" dc:"短链"`
	Enable   bool   `json:"enable" dc:"是否启用"`
}

// ModifyShortOutput is the output of ModifyShort
type ModifyShortOutput bool
