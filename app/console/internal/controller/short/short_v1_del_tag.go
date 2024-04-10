package short

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/short/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

func (c *ControllerV1) DelTag(ctx context.Context, req *v1.DelTagReq) (res *v1.DelTagRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-DelTag")
	defer span.End()

	res = &v1.DelTagRes{}
	res.DelTagOutput, err = service.Short().DelTag(ctx, req.DelTagInput)
	return
}
