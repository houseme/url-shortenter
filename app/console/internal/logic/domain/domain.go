// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package domain

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
	"github.com/houseme/url-shortenter/utility/helper"
)

type sDomain struct {
}

func init() {
	service.RegisterDomain(initDomain())
}

func initDomain() *sDomain {
	return &sDomain{}
}

// CreateDomainAuthorize creates an initDomain domain.
func (s *sDomain) CreateDomainAuthorize(ctx context.Context, in *model.DomainAuthorizeInput) (out *model.DomainAuthorizeOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-domain-CreateDomainAuthorize")
	defer span.End()

	var log = g.Log(helper.Helper().Logger(ctx))
	log.Debug(ctx, "CreateDomainAuthorize params:", in)

	return
}

// QueryDomainAuthorize queries the domain.
func (s *sDomain) QueryDomainAuthorize(ctx context.Context, in *model.QueryDomainAuthorizeInput) (out *model.QueryDomainAuthorizeOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-domain-QueryDomainAuthorize")
	defer span.End()

	var log = g.Log(helper.Helper().Logger(ctx))
	log.Debug(ctx, "QueryDomainAuthorize params:", in)

	return
}

// UpdateDomainAuthorize updates the domain.
func (s *sDomain) UpdateDomainAuthorize(ctx context.Context, in *model.DomainAuthorizeUpdateInput) (out *model.DomainAuthorizeUpdateOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-domain-UpdateDomainAuthorize")
	defer span.End()

	var log = g.Log(helper.Helper().Logger(ctx))
	log.Debug(ctx, "UpdateDomainAuthorize params:", in)

	return
}

// DeleteDomainAuthorize deletes the domain.
func (s *sDomain) DeleteDomainAuthorize(ctx context.Context, in *model.DomainAuthorizeDeleteInput) (out *model.DomainAuthorizeDeleteOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-domain-DeleteDomainAuthorize")
	defer span.End()

	var log = g.Log(helper.Helper().Logger(ctx))
	log.Debug(ctx, "DeleteDomainAuthorize params:", in)

	return
}

// DomainAuthorizeDetail queries the domain.
func (s *sDomain) DomainAuthorizeDetail(ctx context.Context, in *model.DomainAuthorizeDetailInput) (out *model.DomainAuthorizeDetailOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-domain-DomainAuthorizeDetail")
	defer span.End()

	var log = g.Log(helper.Helper().Logger(ctx))
	log.Debug(ctx, "DomainAuthorizeDetail params:", in)

	return
}

// DomainAuthorizeList queries the domain.
func (s *sDomain) DomainAuthorizeList(ctx context.Context, in *model.DomainAuthorizeListInput) (out *model.DomainAuthorizeListOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-domain-DomainAuthorizeList")
	defer span.End()

	var log = g.Log(helper.Helper().Logger(ctx))
	log.Debug(ctx, "DomainAuthorizeList params:", in)

	return
}
