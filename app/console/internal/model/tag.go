/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package model

// CreateTagInput is the input structure for the CreateTag API.
type CreateTagInput struct {
	*Base `json:"-"` // Base structure
	Name  string     `json:"name"  v:"required|length:1,100#请输入标签名称 | 标签名称长度为:min 到:max 位"`
}

// CreateTagOutput is the output structure for the CreateTag API.
type CreateTagOutput struct {
	TagNo uint64 `json:"tagNo,string"` // Tag ID
}

// DelTagInput is the input structure for the DelTag API.
type DelTagInput struct {
	*Base   `json:"-"` // Base structure
	TagNo   uint64     `json:"tagNo,string"  v:"required#请输入标签编号"`   // Tag ID
	ShortNo uint64     `json:"shortNo,string"  v:"required#请输入短链编号"` // Short ID
}

// DelTagOutput is the output structure for the DelTag API.
type DelTagOutput struct {
	ResultCode int    `json:"resultCode" dc:"响应 code 值，100 成功，200 失败"` // Result code
	ResultMsg  string `json:"resultMsg" dc:"返回描述信息"`                   // Result message
}

// AddTagInput is the input structure for the AddTag API.
type AddTagInput struct {
	*Base   `json:"-"` // Base structure
	TagNo   uint64     `json:"tagNo,string"  v:"required#请输入标签编号"`   // Tag ID
	ShortNo uint64     `json:"shortNo,string"  v:"required#请输入短链编号"` // Short ID
}

// AddTagOutput is the output structure for the AddTag API.
type AddTagOutput struct {
	ResultCode int    `json:"resultCode" dc:"响应 code 值，100 成功，200 失败"` // Result code
	ResultMsg  string `json:"resultMsg" dc:"返回描述信息"`                   // Result message
}
