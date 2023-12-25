package account

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/account/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// ModifyPassword is the handler for Account Controller action.
func (c *ControllerV1) ModifyPassword(ctx context.Context, req *v1.ModifyPasswordReq) (res *v1.ModifyPasswordRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-account-ModifyPassword")
	defer span.End()

	res = &v1.ModifyPasswordRes{}
	res.ModifyPasswordOutput, err = service.Account().ModifyPassword(ctx, req.ModifyPasswordInput)
	return
}
