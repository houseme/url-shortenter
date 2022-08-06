package domain

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

type sDomain struct {
}

func init() {
	service.RegisterDomain(initDomain())
}

func initDomain() *sDomain {
	return &sDomain{}
}

// CreateDomainAuthorization creates a initDomain domain.
func (s *sDomain) CreateDomainAuthorization(ctx context.Context, in *model.DomainAuthorizationInput) (out *model.DomainAuthorizationOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-CreateDomainAuthorization")
	defer span.End()

	return
}

// QueryDomainAuthorization queries the domain.
func (s *sDomain) QueryDomainAuthorization(ctx context.Context, in *model.QueryDomainAuthorizationInput) (out *model.QueryDomainAuthorizationOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-QueryDomainAuthorization")
	defer span.End()

	return
}

// UpdateDomainAuthorization updates the domain.
func (s *sDomain) UpdateDomainAuthorization(ctx context.Context, in *model.DomainAuthorizationUpdateInput) (out *model.DomainAuthorizationUpdateOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-UpdateDomainAuthorization")
	defer span.End()

	return
}

// DeleteDomainAuthorization deletes the domain.
func (s *sDomain) DeleteDomainAuthorization(ctx context.Context, in *model.DomainAuthorizationDeleteInput) (out *model.DomainAuthorizationDeleteOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DeleteDomainAuthorization")
	defer span.End()

	return
}

// DomainAuthorizationDetail queries the domain.
func (s *sDomain) DomainAuthorizationDetail(ctx context.Context, in *model.DomainAuthorizationDetailInput) (out *model.DomainAuthorizationDetailOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DomainAuthorizationDetail")
	defer span.End()

	return
}

// DomainAuthorizationList queries the domain.
func (s *sDomain) DomainAuthorizationList(ctx context.Context, in *model.DomainAuthorizationListInput) (out *model.DomainAuthorizationListOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DomainAuthorizationList")
	defer span.End()

	return
}
