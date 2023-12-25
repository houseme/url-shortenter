package auth

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/auth/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// Auth is the handler for Auth Controller action.
func (c *ControllerV1) Auth(ctx context.Context, req *v1.AuthReq) (res *v1.AuthRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-auth-authorization")
	defer span.End()

	res = &v1.AuthRes{}
	res.AuthOutput, err = service.Auth().Authorization(ctx, req.AuthInput)
	return
}
