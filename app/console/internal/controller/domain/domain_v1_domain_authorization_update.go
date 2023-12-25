package domain

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/domain/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// DomainAuthorizationUpdate is the handler for Domain Controller action.
func (c *ControllerV1) DomainAuthorizationUpdate(ctx context.Context, req *v1.DomainAuthorizationUpdateReq) (res *v1.DomainAuthorizationUpdateRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-UpdateDomainAuthorize")
	defer span.End()

	res = &v1.DomainAuthorizationUpdateRes{}
	res.DomainAuthorizeUpdateOutput, err = service.Domain().UpdateDomainAuthorize(ctx, req.DomainAuthorizeUpdateInput)
	return
}
