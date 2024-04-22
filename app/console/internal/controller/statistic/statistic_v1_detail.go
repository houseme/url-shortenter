package statistic

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/statistic/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// Detail is used to get a statistic detail.
func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-statistic-Detail")
	defer span.End()

	res = &v1.DetailRes{}
	res.StatisticDetailOutput, err = service.Statistic().Detail(ctx, req.StatisticDetailInput)
	return
}
