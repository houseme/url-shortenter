package domain

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
	"github.com/houseme/url-shortenter/utility"
)

type sDomain struct {
}

func init() {
	service.RegisterDomain(initDomain())
}

func initDomain() *sDomain {
	return &sDomain{}
}

// CreateDomainAuthorize creates a initDomain domain.
func (s *sDomain) CreateDomainAuthorize(ctx context.Context, in *model.DomainAuthorizeInput) (out *model.DomainAuthorizeOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-CreateDomainAuthorize")
	defer span.End()

	var logger = utility.Helper().Logger(ctx)
	g.Log(logger).Debug(ctx, "CreateDomainAuthorize params:", in)

	return
}

// QueryDomainAuthorize queries the domain.
func (s *sDomain) QueryDomainAuthorize(ctx context.Context, in *model.QueryDomainAuthorizeInput) (out *model.QueryDomainAuthorizeOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-QueryDomainAuthorize")
	defer span.End()

	var logger = utility.Helper().Logger(ctx)
	g.Log(logger).Debug(ctx, "QueryDomainAuthorize params:", in)

	return
}

// UpdateDomainAuthorize updates the domain.
func (s *sDomain) UpdateDomainAuthorize(ctx context.Context, in *model.DomainAuthorizeUpdateInput) (out *model.DomainAuthorizeUpdateOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-UpdateDomainAuthorize")
	defer span.End()

	var logger = utility.Helper().Logger(ctx)
	g.Log(logger).Debug(ctx, "UpdateDomainAuthorize params:", in)

	return
}

// DeleteDomainAuthorize deletes the domain.
func (s *sDomain) DeleteDomainAuthorize(ctx context.Context, in *model.DomainAuthorizeDeleteInput) (out *model.DomainAuthorizeDeleteOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DeleteDomainAuthorize")
	defer span.End()

	var logger = utility.Helper().Logger(ctx)
	g.Log(logger).Debug(ctx, "DeleteDomainAuthorize params:", in)

	return
}

// DomainAuthorizeDetail queries the domain.
func (s *sDomain) DomainAuthorizeDetail(ctx context.Context, in *model.DomainAuthorizeDetailInput) (out *model.DomainAuthorizeDetailOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DomainAuthorizeDetail")
	defer span.End()

	var logger = utility.Helper().Logger(ctx)
	g.Log(logger).Debug(ctx, "DomainAuthorizeDetail params:", in)

	return
}

// DomainAuthorizeList queries the domain.
func (s *sDomain) DomainAuthorizeList(ctx context.Context, in *model.DomainAuthorizeListInput) (out *model.DomainAuthorizeListOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DomainAuthorizeList")
	defer span.End()

	var logger = utility.Helper().Logger(ctx)
	g.Log(logger).Debug(ctx, "DomainAuthorizeList params:", in)

	return
}
