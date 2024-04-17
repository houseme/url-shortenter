package user

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/user/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// UserDetail query the user.
func (c *ControllerV1) UserDetail(ctx context.Context, req *v1.UserDetailReq) (res *v1.UserDetailRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-user-UpdateUser")
	defer span.End()

	res = &v1.UserDetailRes{}
	res.UserDetailOutput, err = service.User().Detail(ctx, req.UserDetailInput)
	return
}
