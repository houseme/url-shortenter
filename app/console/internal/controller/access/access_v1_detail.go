package access

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/access/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// Detail is the handler for Detail
func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-access-Detail")
	defer span.End()

	res = &v1.DetailRes{}
	res.AccessDetailOutput, err = service.Access().Detail(ctx, req.AccessDetailInput)
	return
}
