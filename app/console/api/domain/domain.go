// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package domain

import (
	"context"

	"github.com/houseme/url-shortenter/app/console/api/domain/v1"
)

type IDomainV1 interface {
	DomainAuthorization(ctx context.Context, req *v1.DomainAuthorizationReq) (res *v1.DomainAuthorizationRes, err error)
	DomainAuthorizationList(ctx context.Context, req *v1.DomainAuthorizationListReq) (res *v1.DomainAuthorizationListRes, err error)
	DomainAuthorizationDelete(ctx context.Context, req *v1.DomainAuthorizationDeleteReq) (res *v1.DomainAuthorizationDeleteRes, err error)
	DomainAuthorizationUpdate(ctx context.Context, req *v1.DomainAuthorizationUpdateReq) (res *v1.DomainAuthorizationUpdateRes, err error)
	QueryDomainAuthorization(ctx context.Context, req *v1.QueryDomainAuthorizationReq) (res *v1.QueryDomainAuthorizationRes, err error)
	DomainAuthorizationDetail(ctx context.Context, req *v1.DomainAuthorizationDetailReq) (res *v1.DomainAuthorizationDetailRes, err error)
}
