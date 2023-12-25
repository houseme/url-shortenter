package short

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/short/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// ModifyShort is the handler for Short Controller action.
func (c *ControllerV1) ModifyShort(ctx context.Context, req *v1.ModifyShortReq) (res *v1.ModifyShortRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-short-ModifyShort")
	defer span.End()

	res = &v1.ModifyShortRes{}
	res.ModifyShortOutput, err = service.Short().ModifyShort(ctx, req.ModifyShortInput)
	return
}
