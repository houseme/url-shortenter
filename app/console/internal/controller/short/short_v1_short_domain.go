package short

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/short/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// ShortDomain is the handler for Short Controller action.
func (c *ControllerV1) ShortDomain(ctx context.Context, req *v1.ShortDomainReq) (res *v1.ShortDomainRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-ShortDomain")
	defer span.End()

	res = &v1.ShortDomainRes{}
	res.ShortDomainOutput, err = service.Short().ShortDomain(ctx, req.ShortDomainInput)
	return
}
