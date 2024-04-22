package statistic

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/statistic/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// List is used to list statistic.
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-statistic-List")
	defer span.End()

	res = &v1.ListRes{}
	res.StatisticListOutput, err = service.Statistic().List(ctx, req.StatisticListInput)
	return
}
