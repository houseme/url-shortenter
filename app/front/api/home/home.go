// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package home

import (
	"context"

	v1 "github.com/houseme/url-shortenter/app/front/api/home/v1"
)

type IHomeV1 interface {
	Home(ctx context.Context, req *v1.HomeReq) (res *v1.HomeRes, err error)
}
