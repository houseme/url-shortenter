package user

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/user/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// UpdatePassword updates the password of the user.
func (c *ControllerV1) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordReq) (res *v1.UpdatePasswordRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-user-index")
	defer span.End()

	res = &v1.UpdatePasswordRes{}
	res.UpdatePasswordOutput, err = service.User().UpdatePassword(ctx, req.UpdatePasswordInput)
	return
}
