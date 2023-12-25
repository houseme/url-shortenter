package short

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/short/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// QueryShort is the handler for Short Controller action.
func (c *ControllerV1) QueryShort(ctx context.Context, req *v1.QueryShortReq) (res *v1.QueryShortRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-QueryShort")
	defer span.End()

	res = &v1.QueryShortRes{}
	res.QueryShortOutput, err = service.Short().QueryShort(ctx, req.QueryShortInput)
	return
}
