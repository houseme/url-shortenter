// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package model

// Base is the base model
type Base struct {
	AuthAccountNo    uint64 `json:"authAccountNo" dc:"认证账号"`
	AuthUserNo       uint64 `json:"authUserNo" dc:"认证用户"`
	AuthAccountLevel uint   `json:"authAccountLevel" dc:"账号级别"`
}

// DefaultHandlerResponse .
type DefaultHandlerResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Time    int64       `json:"time,string"` // 返回当前响应时间
	TraceID string      `json:"traceID"`     // 请求唯一标识
}

// PageInfoOutput is the golang structure for PageInfo.
type PageInfoOutput struct {
	Page     int    `json:"page" dc:"页码" v:"required|integer#页码不能为空 | 页码必须为数字" default:"1"`
	PageSize int    `json:"pageSize" dc:"每页数量" v:"required|integer#每页数量不能为空 | 每页数量必须为数字" default:"20"`
	Total    int    `json:"total" dc:"总数" default:"0"`
	LastID   uint64 `json:"lastId,string" dc:"最后一条数据的 ID"`
}

// PageInfoInput is the golang structure for PageInfo.
type PageInfoInput struct {
	Page     int    `json:"page" v:"required|integer|min:1#页码不能为空 | 页码必须为数字 | 页码最小值是 1" dc:"页码"  default:"1"`
	PageSize int    `json:"pageSize" v:"required|integer|min:10#每页数量不能为空 | 每页数量必须为数字 | 每页数量的最小值是 10" dc:"每页数量" default:"20"`
	LastID   uint64 `json:"lastId,string" dc:"最后一条数据的 ID"`
}
