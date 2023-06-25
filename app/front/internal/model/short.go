// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package model

// ShortRedis is a struct for redis data
type ShortRedis struct {
	ShortNo   uint64 `json:"shortNo" dc:"shortNo 短链编号"`
	AccountNo uint64 `json:"accountNo" dc:"accountNo 账户编号"`
	DestURL   string `json:"destUrl" dc:"destUrl 目标链接"`
	IsValid   uint   `json:"isValid" dc:"isValid 是否可用 0 默认 100 正常 200 失效"`
}
