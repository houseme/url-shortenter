package model

// DomainAuthorizationInput is the input for DomainAuthorization
type DomainAuthorizationInput struct {
	*Base     `json:"-"`
	Domain    string `json:"domain" dc:"域名" v:"required|domain#域名不能为空|域名格式不正确"`
	Memo      string `json:"memo" dc:"备注"`
	ICPNumber string `json:"icpNumber" dc:"ICP备案号"`
}

// DomainAuthorizationOutput is the output for DomainAuthorization
type DomainAuthorizationOutput struct {
	State      uint   `json:"state" dc:"状态 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	DomainID   uint64 `json:"domainId" dc:"域名ID"`
	StateDesc  string `json:"stateDesc" dc:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	CreateTime string `json:"createTime" dc:"创建时间"`
}

// QueryDomainAuthorizationInput is the input for QueryDomainAuthorization
type QueryDomainAuthorizationInput struct {
	*Base    `json:"-"`
	DomainID uint64 `json:"domainId" dc:"域名ID" v:"required|integer#域名ID不能为空|域名ID格式不正确"`
}

// QueryDomainAuthorizationOutput is the output for QueryDomainAuthorization
type QueryDomainAuthorizationOutput struct {
	State      uint   `json:"state" dc:"状态 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	DomainID   uint64 `json:"domainId" dc:"域名ID"`
	StateDesc  string `json:"stateDesc" dc:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	CreateTime string `json:"createTime" dc:"创建时间"`
	AuditTime  string `json:"auditTime,omitempty" dc:"审核时间"`
	Reason     string `json:"reason,omitempty" dc:"审核失败原因"`
}

// DomainAuthorizationListInput is the input for QueryDomainAuthorizationList
type DomainAuthorizationListInput struct {
	*Base `json:"-"`
	State uint `json:"state" dc:"状态"`
	Page  int  `json:"page" dc:"页码" v:"required|numeric|min=1#页码必须大于0"`
	Size  int  `json:"size" dc:"每页数量" v:"required|numeric|min=1#每页数量必须大于0"`
}

// DomainAuthorizationListOutput is the output for QueryDomainAuthorizationList
type DomainAuthorizationListOutput struct {
	Total int                          `json:"total" dc:"总数"`
	List  []*DomainAuthorizationDetail `json:"list" dc:"列表"`
	Page  int                          `json:"page" dc:"页码"`
	Size  int                          `json:"size" dc:"每页数量"`
}

// DomainAuthorizationDetailInput is the input for QueryDomainAuthorizationDetail
type DomainAuthorizationDetailInput struct {
	*Base    `json:"-"`
	DomainID uint64 `json:"domainId" dc:"域名ID" v:"required|integer#域名ID不能为空|域名ID格式不正确"`
}

// DomainAuthorizationDetailOutput is the output for DomainAuthorizationDetail
type DomainAuthorizationDetailOutput struct {
	*DomainAuthorizationDetail
}

// DomainAuthorizationDetail query domain authorization detail
type DomainAuthorizationDetail struct {
	State      uint   `json:"state" dc:"状态"`
	DomainID   uint64 `json:"domainId" dc:"域名ID"`
	Domain     string `json:"domain" dc:"域名"`
	Memo       string `json:"memo" dc:"备注"`
	ICPNumber  string `json:"icpNumber" dc:"ICP备案号"`
	CreateTime string `json:"createTime" dc:"创建时间"`
	AuditTime  string `json:"auditTime,omitempty" dc:"审核时间"`
	StateDesc  string `json:"stateDesc" dc:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
}

// DomainAuthorizationDeleteInput is the input for DomainAuthorizationDelete
type DomainAuthorizationDeleteInput struct {
	*Base     `json:"-"`
	DomainIDs []uint64 `json:"domainId" dc:"域名ID" v:"required|integer#域名ID不能为空|域名ID格式不正确"`
}

// DomainAuthorizationDeleteOutput is the output for DomainAuthorizationDelete
type DomainAuthorizationDeleteOutput struct {
	State      uint     `json:"state" dc:"状态"`
	DomainIDs  []uint64 `json:"domainId" dc:"域名ID"`
	Domain     string   `json:"domain" dc:"域名"`
	DeleteTime string   `json:"deleteTime" dc:"删除时间"`
	StateDesc  string   `json:"stateDesc" dc:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
}

// DomainAuthorizationUpdateInput is the input for DomainAuthorizationUpdate
type DomainAuthorizationUpdateInput struct {
	*Base     `json:"-"`
	State     uint     `json:"state" dc:"状态 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	DomainIDs []uint64 `json:"domainId" dc:"域名ID" v:"required|integer#域名ID不能为空|域名ID格式不正确"`
}

// DomainAuthorizationUpdateOutput is the output for DomainAuthorizationUpdate
type DomainAuthorizationUpdateOutput struct {
	State      uint   `json:"state" dc:"状态"`
	StateDesc  string `json:"stateDesc" dc:"状态描述 0:未提交 10:审核中 20:审核通过 30:审核失败 40:已禁用"`
	ModifyTime string `json:"modifyTime" dc:"修改时间"`
	Reason     string `json:"reason,omitempty" dc:"修改失败原因"`
}
