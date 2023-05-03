/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package short

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/short/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// Controller is the controller for the short module.
type Controller struct {
}

// New is the constructor of Controller.
func New() *Controller {
	return &Controller{}
}

// CreateShort is the handler for CreateShort
func (c *Controller) CreateShort(ctx context.Context, req *v1.CreateShortReq) (res *v1.CreateShortRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-CreateShort")
	defer span.End()

	res = &v1.CreateShortRes{}
	if res.CreateShortOutput, err = service.Short().CreateShort(ctx, req.CreateShortInput); err != nil {
		err = gerror.Wrap(err, "account CreateShort failed")
	}
	return
}

// ModifyShort is the handler for ModifyShort
func (c *Controller) ModifyShort(ctx context.Context, req *v1.ModifyShortReq) (res *v1.ModifyShortRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-ModifyShort")
	defer span.End()

	res = &v1.ModifyShortRes{}
	if res.ModifyShortOutput, err = service.Short().ModifyShort(ctx, req.ModifyShortInput); err != nil {
		err = gerror.Wrap(err, "short ModifyShort failed")
	}
	return
}

// QueryShort is the handler for QueryShort
func (c *Controller) QueryShort(ctx context.Context, req *v1.QueryShortReq) (res *v1.QueryShortRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-QueryShort")
	defer span.End()

	res = &v1.QueryShortRes{}
	if res.QueryShortOutput, err = service.Short().QueryShort(ctx, req.QueryShortInput); err != nil {
		err = gerror.Wrap(err, "short QueryShort failed")
	}
	return
}

// QueryStat is the handler for QueryStat
func (c *Controller) QueryStat(ctx context.Context, req *v1.QueryStatReq) (res *v1.QueryStatRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-QueryStat")
	defer span.End()

	res = &v1.QueryStatRes{}
	if res.QueryStatOutput, err = service.Short().QueryStat(ctx, req.QueryStatInput); err != nil {
		err = gerror.Wrap(err, "short QueryStat failed")
	}
	return
}

// ShortDomain is the handler for ShortDomain
func (c *Controller) ShortDomain(ctx context.Context, req *v1.ShortDomainReq) (res *v1.ShortDomainRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-ShortDomain")
	defer span.End()

	res = &v1.ShortDomainRes{}
	if res.ShortDomainOutput, err = service.Short().ShortDomain(ctx, req.ShortDomainInput); err != nil {
		err = gerror.Wrap(err, "short ShortDomain failed")
	}
	return
}
