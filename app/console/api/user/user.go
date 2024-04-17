// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"github.com/houseme/url-shortenter/app/console/api/user/v1"
)

type IUserV1 interface {
	UserDetail(ctx context.Context, req *v1.UserDetailReq) (res *v1.UserDetailRes, err error)
	UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error)
	UpdatePassword(ctx context.Context, req *v1.UpdatePasswordReq) (res *v1.UpdatePasswordRes, err error)
}
