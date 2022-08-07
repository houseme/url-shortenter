package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/houseme/url-shortenter/app/console/internal/model"
)

// DomainAuthorizationReq is the request for DomainAuthorization
type DomainAuthorizationReq struct {
	g.Meta `path:"/domain/create" tags:"Domain Service" method:"Post" summary:"create an domain authorization"`
	*model.DomainAuthorizeInput
}

// DomainAuthorizationRes is the response for DomainAuthorization
type DomainAuthorizationRes struct {
	*model.DomainAuthorizeOutput
}

// DomainAuthorizationListReq is the request for DomainAuthorizationList
type DomainAuthorizationListReq struct {
	g.Meta `path:"/domain/list" tags:"Domain Service" method:"Get" summary:"list domain authorization"`
	*model.DomainAuthorizeListInput
}

// DomainAuthorizationListRes is the response for DomainAuthorizationList
type DomainAuthorizationListRes struct {
	*model.DomainAuthorizeListOutput
}

// DomainAuthorizationDeleteReq is the request for DomainAuthorizationDelete
type DomainAuthorizationDeleteReq struct {
	g.Meta `path:"/domain/delete" tags:"Domain Service" method:"Post" summary:"delete domain authorization"`
	*model.DomainAuthorizeDeleteInput
}

// DomainAuthorizationDeleteRes is the response for DomainAuthorizationDelete
type DomainAuthorizationDeleteRes struct {
	*model.DomainAuthorizeDeleteOutput
}

// DomainAuthorizationUpdateReq is the request for DomainAuthorizationUpdate
type DomainAuthorizationUpdateReq struct {
	g.Meta `path:"/domain/update" tags:"Domain Service" method:"Post" summary:"update domain authorization"`
	*model.DomainAuthorizeUpdateInput
}

// DomainAuthorizationUpdateRes is the response for DomainAuthorizationUpdate
type DomainAuthorizationUpdateRes struct {
	*model.DomainAuthorizeUpdateOutput
}

// QueryDomainAuthorizationReq is the request for QueryDomainAuthorize
type QueryDomainAuthorizationReq struct {
	g.Meta `path:"/domain/query" tags:"Domain Service" method:"Get" summary:"query domain authorization"`
	*model.QueryDomainAuthorizeInput
}

// QueryDomainAuthorizationRes is the response for QueryDomainAuthorize
type QueryDomainAuthorizationRes struct {
	*model.QueryDomainAuthorizeOutput
}

// DomainAuthorizationDetailReq is the request for DomainAuthorizeDetail
type DomainAuthorizationDetailReq struct {
	g.Meta `path:"/domain/detail" tags:"Domain Service" method:"Get" summary:"detail domain authorization"`
	*model.DomainAuthorizeDetailInput
}

// DomainAuthorizationDetailRes is the response for DomainAuthorizeDetail
type DomainAuthorizationDetailRes struct {
	*model.DomainAuthorizeDetailOutput
}
