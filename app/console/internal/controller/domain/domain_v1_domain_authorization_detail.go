package domain

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/domain/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// DomainAuthorizationDetail is the handler for Domain Controller action.
func (c *ControllerV1) DomainAuthorizationDetail(ctx context.Context, req *v1.DomainAuthorizationDetailReq) (res *v1.DomainAuthorizationDetailRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DomainAuthorizeDetail")
	defer span.End()

	res = &v1.DomainAuthorizationDetailRes{}
	res.DomainAuthorizeDetailOutput, err = service.Domain().DomainAuthorizeDetail(ctx, req.DomainAuthorizeDetailInput)
	return
}
