package auth

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/auth/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// CreateAccessToken is the handler for Auth Controller action.
func (c *ControllerV1) CreateAccessToken(ctx context.Context, req *v1.CreateAccessTokenReq) (res *v1.CreateAccessTokenRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-auth-CreateAccessToken")
	defer span.End()

	res = &v1.CreateAccessTokenRes{}
	res.CreateAccessTokenOutput, err = service.Auth().CreateAccessToken(ctx, req.CreateAccessTokenInput)
	return
}
