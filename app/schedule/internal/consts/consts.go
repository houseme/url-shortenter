// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package consts

const (
	// Logger is the name of the logger.
	Logger = "schedule"
	// ShortValid valid status 0 default 100 ok 200 invalid
	ShortValid = 100
	// ShortInvalid failed state
	ShortInvalid = 200

	// ShortCollectStateSuccess 收集状态 0 默认 100 成功 200 失败
	ShortCollectStateSuccess = 100
	// ShortCollectStateFailed 失败状态
	ShortCollectStateFailed = 200
	// ShortCollectStateProcessing 正在收集状态
	ShortCollectStateProcessing = 0

	// VisitState 访问状态 0 默认 100 正常 200 失效
	VisitState = 0
	// VisitStateNormal 访问状态 0 默认 100 正常 200 失效
	VisitStateNormal = 100
	// VisitStateInvalid 失效状态
	VisitStateInvalid = 200

	// ShortAccountNo 帐号
	ShortAccountNo = "short_url_account_no_"
	// ShortShortNo 短网址
	ShortShortNo = "short_url_short_no_"

	// ContentTypeMirror 内容类型 0 默认 100 镜像，200 审核
	ContentTypeMirror = 100
	// ContentTypeAudit 内容类型 0 默认 100 镜像，200 审核
	ContentTypeAudit = 200
	// ContentTypeDefault 内容类型 0 默认 100 镜像，200 审核
	ContentTypeDefault = 0
)
