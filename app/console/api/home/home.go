// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package home

import (
	"context"

	v1 "github.com/houseme/url-shortenter/app/console/api/home/v1"
)

type IHomeV1 interface {
	HomeIndex(ctx context.Context, req *v1.HomeIndexReq) (res *v1.HomeIndexRes, err error)
}
