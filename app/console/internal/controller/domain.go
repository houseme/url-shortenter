package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

type cDomain struct {
}

// Domain authorization
var Domain = &cDomain{}

// CreateDomainAuthorization creates a new domain.
func (c *cDomain) CreateDomainAuthorization(ctx context.Context, req *v1.DomainAuthorizationReq) (res *v1.DomainAuthorizationRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-CreateDomainAuthorization")
	defer span.End()

	res = &v1.DomainAuthorizationRes{}
	if res.DomainAuthorizationOutput, err = service.Domain().CreateDomainAuthorization(ctx, req.DomainAuthorizationInput); err != nil {
		g.Log().Error(ctx, "CreateDomainAuthorization failed err:", err)
	}
	return
}

// QueryDomainAuthorization queries the domain.
func (c *cDomain) QueryDomainAuthorization(ctx context.Context, req *v1.QueryDomainAuthorizationReq) (res *v1.QueryDomainAuthorizationRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-QueryDomainAuthorization")
	defer span.End()

	res = &v1.QueryDomainAuthorizationRes{}
	if res.QueryDomainAuthorizationOutput, err = service.Domain().QueryDomainAuthorization(ctx, req.QueryDomainAuthorizationInput); err != nil {
		g.Log().Error(ctx, "QueryDomainAuthorization failed err:", err)
	}
	return
}

// UpdateDomainAuthorization updates the domain.
func (c *cDomain) UpdateDomainAuthorization(ctx context.Context, req *v1.DomainAuthorizationUpdateReq) (res *v1.DomainAuthorizationUpdateRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-UpdateDomainAuthorization")
	defer span.End()

	res = &v1.DomainAuthorizationUpdateRes{}
	if res.DomainAuthorizationUpdateOutput, err = service.Domain().UpdateDomainAuthorization(ctx, req.DomainAuthorizationUpdateInput); err != nil {
		g.Log().Error(ctx, "UpdateDomainAuthorization failed err:", err)
	}
	return
}

// DeleteDomainAuthorization deletes the domain.
func (c *cDomain) DeleteDomainAuthorization(ctx context.Context, req *v1.DomainAuthorizationDeleteReq) (res *v1.DomainAuthorizationDeleteRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DeleteDomainAuthorization")
	defer span.End()

	res = &v1.DomainAuthorizationDeleteRes{}
	if res.DomainAuthorizationDeleteOutput, err = service.Domain().DeleteDomainAuthorization(ctx, req.DomainAuthorizationDeleteInput); err != nil {
		g.Log().Error(ctx, "DeleteDomainAuthorization failed err:", err)
	}
	return
}

// DomainAuthorizationList lists the domain.
func (c *cDomain) DomainAuthorizationList(ctx context.Context, req *v1.DomainAuthorizationListReq) (res *v1.DomainAuthorizationListRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DomainAuthorizationList")
	defer span.End()

	res = &v1.DomainAuthorizationListRes{}
	if res.DomainAuthorizationListOutput, err = service.Domain().DomainAuthorizationList(ctx, req.DomainAuthorizationListInput); err != nil {
		g.Log().Error(ctx, "ListDomainAuthorization failed err:", err)
	}
	return
}

// DomainAuthorizationDetail gets the domain.
func (c *cDomain) DomainAuthorizationDetail(ctx context.Context, req *v1.DomainAuthorizationDetailReq) (res *v1.DomainAuthorizationDetailRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DomainAuthorizationDetail")
	defer span.End()

	res = &v1.DomainAuthorizationDetailRes{}
	if res.DomainAuthorizationDetailOutput, err = service.Domain().DomainAuthorizationDetail(ctx, req.DomainAuthorizationDetailInput); err != nil {
		g.Log().Error(ctx, "DomainAuthorizationDetail failed err:", err)
	}
	return
}
