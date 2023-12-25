// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package account

import (
	"context"

	v1 "github.com/houseme/url-shortenter/app/console/api/account/v1"
)

type IAccountV1 interface {
	CreateAccount(ctx context.Context, req *v1.CreateAccountReq) (res *v1.CreateAccountRes, err error)
	ModifyAccount(ctx context.Context, req *v1.ModifyAccountReq) (res *v1.ModifyAccountRes, err error)
	ModifyPassword(ctx context.Context, req *v1.ModifyPasswordReq) (res *v1.ModifyPasswordRes, err error)
}
