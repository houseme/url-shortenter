// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package model

// CreateMerchantInput is the input of CreateMerchant.
type CreateMerchantInput struct {
	*Base     `json:"-"`
	MerName   string `json:"merName" desc:"商户名称" v:"required#商户名称不能为空"`
	MerAvatar string `json:"merAvatar" desc:"商户头像" v:"required#商户头像不能为空"`
	MerMobile string `json:"merMobile" desc:"商户手机号" v:"required#商户手机号不能为空"`
}

// CreateMerchantOutput is the output of CreateMerchant.
type CreateMerchantOutput struct {
}
