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
	// Id is the id of access.
}

// AccessDetailInput is the input of AccessDetail.
type AccessDetailInput struct {
	// Id is the id of access.
	*Base `json:"-"`
	ID    uint64 `json:"id,string" description:"访问 log ID" v:"required|integer#访问 log ID 不能为空 | 访问 log ID 必须为数字"`
}

// AccessDetailOutput is the output of AccessDetail.
type AccessDetailOutput struct {
	// Access is the access.
}
