package domain

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/domain/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// DomainAuthorizationDelete is the handler for domain Controller action.
func (c *ControllerV1) DomainAuthorizationDelete(ctx context.Context, req *v1.DomainAuthorizationDeleteReq) (res *v1.DomainAuthorizationDeleteRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DeleteDomainAuthorize")
	defer span.End()

	res = &v1.DomainAuthorizationDeleteRes{}
	res.DomainAuthorizeDeleteOutput, err = service.Domain().DeleteDomainAuthorize(ctx, req.DomainAuthorizeDeleteInput)
	return
}
