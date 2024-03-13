// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

// Package consts defines the constants of the application.
package consts

const (
	// DefaultLogger 默认日志
	DefaultLogger = "default"

	// ShortValid 有效状态 0 默认 100 正常 200 失效
	ShortValid = 100
	// ShortInvalid 失效状态
	ShortInvalid = 200

	// AccessLogsPrefix 日志前缀
	AccessLogsPrefix = "OH_ACCESS_LOGS#"

	// VisitState 访问状态 0 默认 100 正常 200 失效
	VisitState = 0
	// VisitStateNormal 访问状态 0 默认 100 正常 200 失效
	VisitStateNormal = 100
	// VisitStateInvalid 失效状态
	VisitStateInvalid = 200
)
