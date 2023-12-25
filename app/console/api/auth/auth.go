// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	v1 "github.com/houseme/url-shortenter/app/console/api/auth/v1"
)

type IAuthV1 interface {
	CreateAccessToken(ctx context.Context, req *v1.CreateAccessTokenReq) (res *v1.CreateAccessTokenRes, err error)
	Auth(ctx context.Context, req *v1.AuthReq) (res *v1.AuthRes, err error)
}
