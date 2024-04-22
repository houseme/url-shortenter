package short

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/short/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// QueryShortList is the handler for Short Controller action.
func (c *ControllerV1) QueryShortList(ctx context.Context, req *v1.QueryShortListReq) (res *v1.QueryShortListRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-QueryShortList")
	defer span.End()

	res = &v1.QueryShortListRes{}
	res.QueryShortListOutput, err = service.Short().QueryShortList(ctx, req.QueryShortListInput)
	return
}
