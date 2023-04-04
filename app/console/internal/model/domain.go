// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package model

// DomainAuthorizeInput is the input for DomainAuthorization
type DomainAuthorizeInput struct {
	*Base     `json:"-"`
	Domain    string `json:"domain" dc:"域名" v1:"required|domain#域名不能为空|域名格式不正确"`
	Memo      string `json:"memo" dc:"备注"`
	ICPNumber string `json:"icpNumber" dc:"ICP备案号"`
}

// DomainAuthorizeOutput is the output for DomainAuthorization
type DomainAuthorizeOutput struct {
	State      uint   `json:"state" dc:"状态 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	DomainNo   uint64 `json:"domainNo" dc:"域名ID"`
	StateDesc  string `json:"stateDesc" dc:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	CreateTime string `json:"createTime" dc:"创建时间"`
}

// QueryDomainAuthorizeInput is the input for QueryDomainAuthorize
type QueryDomainAuthorizeInput struct {
	*Base    `json:"-"`
	DomainNo uint64 `json:"domainNo" dc:"域名ID" v1:"required|integer#域名ID不能为空|域名ID格式不正确"`
}

// QueryDomainAuthorizeOutput is the output for QueryDomainAuthorize
type QueryDomainAuthorizeOutput struct {
	State      uint   `json:"state" dc:"状态 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	DomainNo   uint64 `json:"domainNo" dc:"域名ID"`
	StateDesc  string `json:"stateDesc" dc:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	CreateTime string `json:"createTime" dc:"创建时间"`
	AuditTime  string `json:"auditTime,omitempty" dc:"审核时间"`
	Reason     string `json:"reason,omitempty" dc:"审核失败原因"`
}

// DomainAuthorizeListInput is the input for QueryDomainAuthorizationList
type DomainAuthorizeListInput struct {
	*Base `json:"-"`
	State uint `json:"state" dc:"状态"`
	Page  int  `json:"page" dc:"页码" v1:"required|numeric|min=1#页码必须大于0"`
	Size  int  `json:"size" dc:"每页数量" v1:"required|numeric|min=1#每页数量必须大于0"`
}

// DomainAuthorizeListOutput is the output for QueryDomainAuthorizationList
type DomainAuthorizeListOutput struct {
	Total int                          `json:"total" dc:"总数"`
	List  []*DomainAuthorizationDetail `json:"list" dc:"列表"`
	Page  int                          `json:"page" dc:"页码"`
	Size  int                          `json:"size" dc:"每页数量"`
}

// DomainAuthorizeDetailInput is the input for QueryDomainAuthorizationDetail
type DomainAuthorizeDetailInput struct {
	*Base    `json:"-"`
	DomainNo uint64 `json:"domainNo" dc:"域名ID" v1:"required|integer#域名ID不能为空|域名ID格式不正确"`
}

// DomainAuthorizeDetailOutput is the output for DomainAuthorizationDetail
type DomainAuthorizeDetailOutput struct {
	*DomainAuthorizationDetail
}

// DomainAuthorizationDetail query domain authorization detail
type DomainAuthorizationDetail struct {
	State      uint   `json:"state" dc:"状态"`
	DomainNo   uint64 `json:"domainNo" dc:"域名ID"`
	Domain     string `json:"domain" dc:"域名"`
	Memo       string `json:"memo" dc:"备注"`
	ICPNumber  string `json:"icpNumber" dc:"ICP备案号"`
	CreateTime string `json:"createTime" dc:"创建时间"`
	AuditTime  string `json:"auditTime,omitempty" dc:"审核时间"`
	StateDesc  string `json:"stateDesc" dc:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
}

// DomainAuthorizeDeleteInput is the input for DomainAuthorizationDelete
type DomainAuthorizeDeleteInput struct {
	*Base     `json:"-"`
	DomainNos []uint64 `json:"domainNos" dc:"域名ID" v1:"required|integer#域名ID不能为空|域名ID格式不正确"`
}

// DomainAuthorizeDeleteOutput is the output for DomainAuthorizationDelete
type DomainAuthorizeDeleteOutput struct {
	State      uint     `json:"state" dc:"状态"`
	DomainNos  []uint64 `json:"domainNos" dc:"域名ID"`
	Domain     string   `json:"domain" dc:"域名"`
	DeleteTime string   `json:"deleteTime" dc:"删除时间"`
	StateDesc  string   `json:"stateDesc" dc:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
}

// DomainAuthorizeUpdateInput is the input for DomainAuthorizationUpdate
type DomainAuthorizeUpdateInput struct {
	*Base     `json:"-"`
	State     uint     `json:"state" dc:"状态 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	DomainNos []uint64 `json:"domainNos" dc:"域名ID" v1:"required|integer#域名ID不能为空|域名ID格式不正确"`
}

// DomainAuthorizeUpdateOutput is the output for DomainAuthorizationUpdate
type DomainAuthorizeUpdateOutput struct {
	State      uint   `json:"state" dc:"状态"`
	StateDesc  string `json:"stateDesc" dc:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	ModifyTime string `json:"modifyTime" dc:"修改时间"`
	Reason     string `json:"reason,omitempty" dc:"修改失败原因"`
}
