package short

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/short/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// CreateTag is the handler for CreateTag
func (c *ControllerV1) CreateTag(ctx context.Context, req *v1.CreateTagReq) (res *v1.CreateTagRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-CreateTag")
	defer span.End()

	res = &v1.CreateTagRes{}
	res.CreateTagOutput, err = service.Short().CreateTag(ctx, req.CreateTagInput)
	return
}
