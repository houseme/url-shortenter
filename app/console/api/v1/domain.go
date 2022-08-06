package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/houseme/url-shortenter/app/console/internal/model"
)

// DomainAuthorizationReq is the request for DomainAuthorization
type DomainAuthorizationReq struct {
	g.Meta `path:"/domain/create" tags:"Domain Service" method:"Post" summary:"create an domain authorization"`
	*model.DomainAuthorizationInput
}

// DomainAuthorizationRes is the response for DomainAuthorization
type DomainAuthorizationRes struct {
	*model.DomainAuthorizationOutput
}

// DomainAuthorizationListReq is the request for DomainAuthorizationList
type DomainAuthorizationListReq struct {
	g.Meta `path:"/domain/list" tags:"Domain Service" method:"Get" summary:"list domain authorization"`
	*model.DomainAuthorizationListInput
}

// DomainAuthorizationListRes is the response for DomainAuthorizationList
type DomainAuthorizationListRes struct {
	*model.DomainAuthorizationListOutput
}

// DomainAuthorizationDeleteReq is the request for DomainAuthorizationDelete
type DomainAuthorizationDeleteReq struct {
	g.Meta `path:"/domain/delete" tags:"Domain Service" method:"Post" summary:"delete domain authorization"`
	*model.DomainAuthorizationDeleteInput
}

// DomainAuthorizationDeleteRes is the response for DomainAuthorizationDelete
type DomainAuthorizationDeleteRes struct {
	*model.DomainAuthorizationDeleteOutput
}

// DomainAuthorizationUpdateReq is the request for DomainAuthorizationUpdate
type DomainAuthorizationUpdateReq struct {
	g.Meta `path:"/domain/update" tags:"Domain Service" method:"Post" summary:"update domain authorization"`
	*model.DomainAuthorizationUpdateInput
}

// DomainAuthorizationUpdateRes is the response for DomainAuthorizationUpdate
type DomainAuthorizationUpdateRes struct {
	*model.DomainAuthorizationUpdateOutput
}

// QueryDomainAuthorizationReq is the request for QueryDomainAuthorization
type QueryDomainAuthorizationReq struct {
	g.Meta `path:"/domain/query" tags:"Domain Service" method:"Get" summary:"query domain authorization"`
	*model.QueryDomainAuthorizationInput
}

// QueryDomainAuthorizationRes is the response for QueryDomainAuthorization
type QueryDomainAuthorizationRes struct {
	*model.QueryDomainAuthorizationOutput
}

// DomainAuthorizationDetailReq is the request for DomainAuthorizationDetail
type DomainAuthorizationDetailReq struct {
	g.Meta `path:"/domain/detail" tags:"Domain Service" method:"Get" summary:"detail domain authorization"`
	*model.DomainAuthorizationDetailInput
}

// DomainAuthorizationDetailRes is the response for DomainAuthorizationDetail
type DomainAuthorizationDetailRes struct {
	*model.DomainAuthorizationDetailOutput
}
