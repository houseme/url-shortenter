package domain

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/domain/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// QueryDomainAuthorization is the handler for Domain Controller action.
func (c *ControllerV1) QueryDomainAuthorization(ctx context.Context, req *v1.QueryDomainAuthorizationReq) (res *v1.QueryDomainAuthorizationRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-QueryDomainAuthorize")
	defer span.End()

	res = &v1.QueryDomainAuthorizationRes{}
	res.QueryDomainAuthorizeOutput, err = service.Domain().QueryDomainAuthorize(ctx, req.QueryDomainAuthorizeInput)
	return
}
