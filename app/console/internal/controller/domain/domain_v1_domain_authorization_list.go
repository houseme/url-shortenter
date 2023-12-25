package domain

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/domain/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// DomainAuthorizationList is the handler for Domain Controller action.
func (c *ControllerV1) DomainAuthorizationList(ctx context.Context, req *v1.DomainAuthorizationListReq) (res *v1.DomainAuthorizationListRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-domain-DomainAuthorizationList")
	defer span.End()

	res = &v1.DomainAuthorizationListRes{}
	res.DomainAuthorizeListOutput, err = service.Domain().DomainAuthorizeList(ctx, req.DomainAuthorizeListInput)
	return
}
