package short

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/short/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// QueryStat is the handler for Short Controller action.
func (c *ControllerV1) QueryStat(ctx context.Context, req *v1.QueryStatReq) (res *v1.QueryStatRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-QueryStat")
	defer span.End()

	res = &v1.QueryStatRes{}
	res.QueryStatOutput, err = service.Short().QueryStat(ctx, req.QueryStatInput)
	return
}
