package short

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/short/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// CreateShort is the handler for Short Controller action.
func (c *ControllerV1) CreateShort(ctx context.Context, req *v1.CreateShortReq) (res *v1.CreateShortRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-CreateShort")
	defer span.End()

	res = &v1.CreateShortRes{}
	res.CreateShortOutput, err = service.Short().CreateShort(ctx, req.CreateShortInput)
	return
}
