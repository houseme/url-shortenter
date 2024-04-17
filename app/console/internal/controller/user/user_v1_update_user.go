package user

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/user/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// UpdateUser updates the user.
func (c *ControllerV1) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-user-UpdateUser")
	defer span.End()

	res = &v1.UpdateUserRes{}
	res.UpdateUserOutput, err = service.User().Update(ctx, req.UpdateUserInput)
	return
}
