// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

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
	DestURL     string      `json:"destUrl" dc:"destUrl 原始 url"`
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
	DestURL                 string      `json:"destUrl" dc:"原始 url"`
	Memo                    string      `json:"memo,omitempty" dc:"备注"`
	State                   uint        `json:"state" dc:"状态"`
	StateMsg                string      `json:"stateMsg" dc:"状态描述"`
	CreateTime              *gtime.Time `json:"createTime" dc:"创建时间"`
	DisableTime             *gtime.Time `json:"disableTime,omitempty" dc:"禁用时间"`
	TodayCount              uint64      `json:"todayCount" dc:"今日访问次数"`
	YesterdayCount          uint64      `json:"yesterdayCount" dc:"昨日访问次数"`
	LastSevenDaysCount      uint64      `json:"lastSevenDaysCount" dc:"近 7 日访问次数"`
	MonthlyCount            uint64      `json:"monthlyCount" dc:"本月访问次数"`
	LastMonthlyCount        uint64      `json:"lastMonthlyCount" dc:"上月访问次数"`
	TotalCount              uint64      `json:"totalCount" dc:"总访问次数"`
	ProtoTodayCount         uint64      `json:"protoTodayCount" dc:"今日独立 IP 访问次数"`
	ProtoYesterdayCount     uint64      `json:"protoYesterdayCount" dc:"昨日独立 IP 访问次数"`
	ProtoLastSevenDaysCount uint64      `json:"protoLastSevenDaysCount" dc:"近 7 日独立 IP 访问次数"`
	ProtoMonthlyCount       uint64      `json:"protoMonthlyCount" dc:"本月独立 IP 访问次数"`
	ProtoLastMonthlyCount   uint64      `json:"protoLastMonthlyCount" dc:"上月独立 IP 访问次数"`
	ProtoTotalCount         uint64      `json:"protoTotalCount" dc:"总独立 IP 访问次数"`
}

// ModifyShortInput is the input for ModifyShort
type ModifyShortInput struct {
	*Base    `json:"-"`
	ShortURL string `json:"shortUrl" dc:"短链"`
	Enable   bool   `json:"enable" dc:"是否启用"`
}

// ModifyShortOutput is the output of ModifyShort
type ModifyShortOutput bool

// ShortDomainInput is the input for ShortDomain
type ShortDomainInput struct {
	*Base `json:"-"`
}

// ShortDomainOutput is the output of ShortDomain
type ShortDomainOutput struct {
	List []*ShortDomainItem `json:"list" dc:"短域名列表"`
}

// ShortDomainItem is the item of ShortDomain
type ShortDomainItem struct {
	Domain   string `json:"shortDomain" dc:"短域名"`
	DomainNo uint64 `json:"domainNo,string" dc:"短域名编号"`
	Memo     string `json:"memo" dc:"备注"`
}
