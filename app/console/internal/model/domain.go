// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package model

// DomainAuthorizeInput is the input for DomainAuthorization
// 域名信息提交
type DomainAuthorizeInput struct {
	*Base     `json:"-"`
	Domain    string `json:"domain" dc:"域名" v:"required|domain#域名不能为空 | 域名格式不正确"`
	Memo      string `json:"memo" dc:"备注"`
	ICPNumber string `json:"icpNumber" dc:"ICP 备案号"`
}

// DomainAuthorizeOutput is the output for DomainAuthorization
type DomainAuthorizeOutput struct {
	State      uint   `json:"state" dc:"状态 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	DomainNo   uint64 `json:"domainNo" dc:"域名 ID"`
	StateDesc  string `json:"stateDesc" dc:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	CreateTime string `json:"createTime" dc:"创建时间"`
}

// QueryDomainAuthorizeInput is the input for QueryDomainAuthorize
// 查询域名审核情况
type QueryDomainAuthorizeInput struct {
	*Base    `json:"-"`
	DomainNo uint64 `json:"domainNo" dc:"域名 ID" v:"required|integer#域名 ID 不能为空 | 域名 ID 格式不正确"`
}

// QueryDomainAuthorizeOutput is the output for QueryDomainAuthorize
type QueryDomainAuthorizeOutput struct {
	State      uint   `json:"state" dc:"状态 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	DomainNo   uint64 `json:"domainNo" dc:"域名 ID"`
	StateDesc  string `json:"stateDesc" dc:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	CreateTime string `json:"createTime" dc:"创建时间"`
	AuditTime  string `json:"auditTime,omitempty" dc:"审核时间"`
	Reason     string `json:"reason,omitempty" dc:"审核失败原因"`
}

// DomainAuthorizeListInput is the input for QueryDomainAuthorizationList
// 查询域名列表
type DomainAuthorizeListInput struct {
	*Base `json:"-"`
	*PageInfoInput
	State uint `json:"state" dc:"状态"`
}

// DomainAuthorizeListOutput is the output for QueryDomainAuthorizationList
type DomainAuthorizeListOutput struct {
	List     []*DomainAuthorizationDetail `json:"list" dc:"列表"`
	Paginate *PageInfoOutput              `json:"paginate" description:"分页信息"`
}

// DomainAuthorizeDetailInput is the input for QueryDomainAuthorizationDetail
// 授权域名请求参数
type DomainAuthorizeDetailInput struct {
	*Base    `json:"-"`
	DomainNo uint64 `json:"domainNo" dc:"域名 ID" v:"required|integer#域名 ID 不能为空 | 域名 ID 格式不正确"`
}

// DomainAuthorizeDetailOutput is the output for DomainAuthorizationDetail
// 响应结果
type DomainAuthorizeDetailOutput struct {
	*DomainAuthorizationDetail
}

// DomainAuthorizationDetail query domain authorization detail
type DomainAuthorizationDetail struct {
	State      uint   `json:"state" dc:"状态"`
	DomainNo   uint64 `json:"domainNo" dc:"域名 ID"`
	Domain     string `json:"domain" dc:"域名"`
	Memo       string `json:"memo" dc:"备注"`
	ICPNumber  string `json:"icpNumber" dc:"ICP 备案号"`
	CreateTime string `json:"createTime" dc:"创建时间"`
	AuditTime  string `json:"auditTime,omitempty" dc:"审核时间"`
	StateDesc  string `json:"stateDesc" dc:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
}

// DomainAuthorizeDeleteInput is the input for DomainAuthorizationDelete
// 删除域名请求参数
type DomainAuthorizeDeleteInput struct {
	*Base     `json:"-"`
	DomainNos []uint64 `json:"domainNos" dc:"域名 ID" v:"required|integer#域名 ID 不能为空 | 域名 ID 格式不正确"`
}

// DomainAuthorizeDeleteOutput is the output for DomainAuthorizationDelete
type DomainAuthorizeDeleteOutput struct {
	State      uint     `json:"state" dc:"状态"`
	DomainNos  []uint64 `json:"domainNos" dc:"域名 ID"`
	Domain     string   `json:"domain" dc:"域名"`
	DeleteTime string   `json:"deleteTime" dc:"删除时间"`
	StateDesc  string   `json:"stateDesc" dc:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
}

// DomainAuthorizeUpdateInput is the input for DomainAuthorizationUpdate
// 更新域名请求参数 批量更新
type DomainAuthorizeUpdateInput struct {
	*Base     `json:"-"`
	State     uint     `json:"state" dc:"状态 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	DomainNos []uint64 `json:"domainNos" dc:"域名 ID" v:"required|integer#域名 ID 不能为空 | 域名 ID 格式不正确"`
}

// DomainAuthorizeUpdateOutput is the output for DomainAuthorizationUpdate
// 更新域名响应结果
type DomainAuthorizeUpdateOutput struct {
	State      uint   `json:"state" dc:"状态"`
	StateDesc  string `json:"stateDesc" dc:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	ModifyTime string `json:"modifyTime" dc:"修改时间"`
	Reason     string `json:"reason,omitempty" dc:"修改失败原因"`
}
