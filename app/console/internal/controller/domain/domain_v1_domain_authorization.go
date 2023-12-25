package domain

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/domain/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// DomainAuthorization is the handler for domain Controller action.
func (c *ControllerV1) DomainAuthorization(ctx context.Context, req *v1.DomainAuthorizationReq) (res *v1.DomainAuthorizationRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-CreateDomainAuthorize")
	defer span.End()

	res = &v1.DomainAuthorizationRes{}
	res.DomainAuthorizeOutput, err = service.Domain().CreateDomainAuthorize(ctx, req.DomainAuthorizeInput)
	return
}
