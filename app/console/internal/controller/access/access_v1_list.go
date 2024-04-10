package access

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/access/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// List is the handler for List
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-access-List")
	defer span.End()

	res = &v1.ListRes{}
	res.AccessListOutput, err = service.Access().List(ctx, req.AccessListInput)
	return
}
