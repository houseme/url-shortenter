/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package model

// AccessListInput is the input of AccessList.
type AccessListInput struct {
	// Page is the page number.
	*Base `json:"-"`
	*PageInfoInput
}

// AccessListOutput is the output of AccessList.
type AccessListOutput struct {
	// List is the list of access.
	List     []*AccessItem   `json:"list" description:"访问 log 列表"`
	Paginate *PageInfoOutput `json:"paginate" description:"分页信息"`
}

// AccessItem is the item of access.
type AccessItem struct {
	// ID is the id of access.
	ID         uint64 `json:"id,string" description:"访问 log ID"`
	ShortNo    uint64 `json:"short_no,string" description:"短链 ID"`
	ShortUrl   string `json:"short_url" description:"短链地址"`
	ShortAll   string `json:"short_all" description:"短链完整地址"`
	YearTime   uint   `json:"year_time" description:"访问时间 年份"`
	MonthTime  uint   `json:"month_time" description:"访问时间 月份"`
	DayTime    uint   `json:"day_time" description:"访问时间 日期"`
	AccessTime string `json:"access_time" description:"访问时间"`
	AccessDate string `json:"access_date" description:"访问日期"`
	UserAgent  string `json:"user_agent" description:"用户代理"`
	IP         string `json:"ip" description:"访问 IP"`
	TraceID    string `json:"trace_id" description:"链路追踪 ID"`
	VisitState uint   `json:"visit_state" description:"访问状态 0 默认，100 正常 200 失效"`
	ServerIP   string `json:"server_ip" description:"服务端 IP"`
}

// AccessDetailInput is the input of AccessDetail.
type AccessDetailInput struct {
	// ID is the id of access.
	*Base `json:"-"`
	ID    uint64 `json:"id,string" description:"访问 log ID" v:"required|integer#访问 log ID 不能为空 | 访问 log ID 必须为数字"`
}

// AccessDetailOutput is the output of AccessDetail.
type AccessDetailOutput struct {
	// Access is the access.
	Access *AccessItem `json:"access" description:"访问 log 详情"`
}
