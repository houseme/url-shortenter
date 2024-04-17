package home

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/home/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// HomeIndex is the handler for Home Controller action.
func (c *ControllerV1) HomeIndex(ctx context.Context, req *v1.HomeIndexReq) (res *v1.HomeIndexRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-home-index")
	defer span.End()

	res = &v1.HomeIndexRes{}
	res.HomeIndexOutput, err = service.Home().Index(ctx, req.HomeIndexInput)
	return
}
