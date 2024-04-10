package short

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/short/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// AddTag is the handler for AddTag
func (c *ControllerV1) AddTag(ctx context.Context, req *v1.AddTagReq) (res *v1.AddTagRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-AddTag")
	defer span.End()

	res = &v1.AddTagRes{}
	res.AddTagOutput, err = service.Short().AddTag(ctx, req.AddTagInput)
	return
}
