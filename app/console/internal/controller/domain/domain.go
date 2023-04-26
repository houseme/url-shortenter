/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package domain

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/domain/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// Controller is the controller for domain.
type Controller struct {
}

// New is the constructor of Controller.
func New() *Controller {
	return &Controller{}
}

// CreateDomainAuthorization creates a new domain.
func (c *Controller) CreateDomainAuthorization(ctx context.Context, req *v1.DomainAuthorizationReq) (res *v1.DomainAuthorizationRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-CreateDomainAuthorize")
	defer span.End()

	res = &v1.DomainAuthorizationRes{}
	if res.DomainAuthorizeOutput, err = service.Domain().CreateDomainAuthorize(ctx, req.DomainAuthorizeInput); err != nil {
		err = gerror.Wrap(err, "CreateDomainAuthorize failed")
	}
	return
}

// QueryDomainAuthorization queries the domain.
func (c *Controller) QueryDomainAuthorization(ctx context.Context, req *v1.QueryDomainAuthorizationReq) (res *v1.QueryDomainAuthorizationRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-QueryDomainAuthorize")
	defer span.End()

	res = &v1.QueryDomainAuthorizationRes{}
	if res.QueryDomainAuthorizeOutput, err = service.Domain().QueryDomainAuthorize(ctx, req.QueryDomainAuthorizeInput); err != nil {
		err = gerror.Wrap(err, "QueryDomainAuthorize failed")
	}
	return
}

// UpdateDomainAuthorization updates the domain.
func (c *Controller) UpdateDomainAuthorization(ctx context.Context, req *v1.DomainAuthorizationUpdateReq) (res *v1.DomainAuthorizationUpdateRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-UpdateDomainAuthorize")
	defer span.End()

	res = &v1.DomainAuthorizationUpdateRes{}
	if res.DomainAuthorizeUpdateOutput, err = service.Domain().UpdateDomainAuthorize(ctx, req.DomainAuthorizeUpdateInput); err != nil {
		err = gerror.Wrap(err, "UpdateDomainAuthorize failed")
	}
	return
}

// DeleteDomainAuthorization deletes the domain.
func (c *Controller) DeleteDomainAuthorization(ctx context.Context, req *v1.DomainAuthorizationDeleteReq) (res *v1.DomainAuthorizationDeleteRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DeleteDomainAuthorize")
	defer span.End()

	res = &v1.DomainAuthorizationDeleteRes{}
	if res.DomainAuthorizeDeleteOutput, err = service.Domain().DeleteDomainAuthorize(ctx, req.DomainAuthorizeDeleteInput); err != nil {
		err = gerror.Wrap(err, "DeleteDomainAuthorize failed")
	}
	return
}

// DomainAuthorizationList lists the domain.
func (c *Controller) DomainAuthorizationList(ctx context.Context, req *v1.DomainAuthorizationListReq) (res *v1.DomainAuthorizationListRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DomainAuthorizationList")
	defer span.End()

	res = &v1.DomainAuthorizationListRes{}
	if res.DomainAuthorizeListOutput, err = service.Domain().DomainAuthorizeList(ctx, req.DomainAuthorizeListInput); err != nil {
		err = gerror.Wrap(err, "ListDomainAuthorization failed")
	}
	return
}

// DomainAuthorizationDetail gets the domain.
func (c *Controller) DomainAuthorizationDetail(ctx context.Context, req *v1.DomainAuthorizationDetailReq) (res *v1.DomainAuthorizationDetailRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DomainAuthorizeDetail")
	defer span.End()

	res = &v1.DomainAuthorizationDetailRes{}
	if res.DomainAuthorizeDetailOutput, err = service.Domain().DomainAuthorizeDetail(ctx, req.DomainAuthorizeDetailInput); err != nil {
		err = gerror.Wrap(err, "DomainAuthorizeDetail failed ")
	}
	return
}
