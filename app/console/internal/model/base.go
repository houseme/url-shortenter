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
