package account

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/account/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// ModifyAccount is the handler for Account Controller action.
func (c *ControllerV1) ModifyAccount(ctx context.Context, req *v1.ModifyAccountReq) (res *v1.ModifyAccountRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-account-ModifyAccount")
	defer span.End()

	res = &v1.ModifyAccountRes{}
	res.ModifyAccountOutput, err = service.Account().ModifyAccount(ctx, req.ModifyAccountInput)
	return
}
