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
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-CreateDomainAuthorize")
	defer span.End()

	res = &v1.DomainAuthorizationRes{}
	if res.DomainAuthorizeOutput, err = service.Domain().CreateDomainAuthorize(ctx, req.DomainAuthorizeInput); err != nil {
		g.Log().Error(ctx, "CreateDomainAuthorize failed err:", err)
	}
	return
}

// QueryDomainAuthorization queries the domain.
func (c *cDomain) QueryDomainAuthorization(ctx context.Context, req *v1.QueryDomainAuthorizationReq) (res *v1.QueryDomainAuthorizationRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-QueryDomainAuthorize")
	defer span.End()

	res = &v1.QueryDomainAuthorizationRes{}
	if res.QueryDomainAuthorizeOutput, err = service.Domain().QueryDomainAuthorize(ctx, req.QueryDomainAuthorizeInput); err != nil {
		g.Log().Error(ctx, "QueryDomainAuthorize failed err:", err)
	}
	return
}

// UpdateDomainAuthorization updates the domain.
func (c *cDomain) UpdateDomainAuthorization(ctx context.Context, req *v1.DomainAuthorizationUpdateReq) (res *v1.DomainAuthorizationUpdateRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-UpdateDomainAuthorize")
	defer span.End()

	res = &v1.DomainAuthorizationUpdateRes{}
	if res.DomainAuthorizeUpdateOutput, err = service.Domain().UpdateDomainAuthorize(ctx, req.DomainAuthorizeUpdateInput); err != nil {
		g.Log().Error(ctx, "UpdateDomainAuthorize failed err:", err)
	}
	return
}

// DeleteDomainAuthorization deletes the domain.
func (c *cDomain) DeleteDomainAuthorization(ctx context.Context, req *v1.DomainAuthorizationDeleteReq) (res *v1.DomainAuthorizationDeleteRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DeleteDomainAuthorize")
	defer span.End()

	res = &v1.DomainAuthorizationDeleteRes{}
	if res.DomainAuthorizeDeleteOutput, err = service.Domain().DeleteDomainAuthorize(ctx, req.DomainAuthorizeDeleteInput); err != nil {
		g.Log().Error(ctx, "DeleteDomainAuthorize failed err:", err)
	}
	return
}

// DomainAuthorizationList lists the domain.
func (c *cDomain) DomainAuthorizationList(ctx context.Context, req *v1.DomainAuthorizationListReq) (res *v1.DomainAuthorizationListRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DomainAuthorizationList")
	defer span.End()

	res = &v1.DomainAuthorizationListRes{}
	if res.DomainAuthorizeListOutput, err = service.Domain().DomainAuthorizeList(ctx, req.DomainAuthorizeListInput); err != nil {
		g.Log().Error(ctx, "ListDomainAuthorization failed err:", err)
	}
	return
}

// DomainAuthorizationDetail gets the domain.
func (c *cDomain) DomainAuthorizationDetail(ctx context.Context, req *v1.DomainAuthorizationDetailReq) (res *v1.DomainAuthorizationDetailRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DomainAuthorizeDetail")
	defer span.End()

	res = &v1.DomainAuthorizationDetailRes{}
	if res.DomainAuthorizeDetailOutput, err = service.Domain().DomainAuthorizeDetail(ctx, req.DomainAuthorizeDetailInput); err != nil {
		g.Log().Error(ctx, "DomainAuthorizeDetail failed err:", err)
	}
	return
}
