package account

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/account/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// CreateAccount is the handler for Account Controller action.
func (c *ControllerV1) CreateAccount(ctx context.Context, req *v1.CreateAccountReq) (res *v1.CreateAccountRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-account-CreateAccount")
	defer span.End()

	res = &v1.CreateAccountRes{}
	res.CreateAccountOutput, err = service.Account().CreateAccount(ctx, req.CreateAccountInput)
	return
}
