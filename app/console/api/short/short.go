// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package short

import (
	"context"

	v1 "github.com/houseme/url-shortenter/app/console/api/short/v1"
)

type IShortV1 interface {
	CreateShort(ctx context.Context, req *v1.CreateShortReq) (res *v1.CreateShortRes, err error)
	QueryShort(ctx context.Context, req *v1.QueryShortReq) (res *v1.QueryShortRes, err error)
	ModifyShort(ctx context.Context, req *v1.ModifyShortReq) (res *v1.ModifyShortRes, err error)
	QueryStat(ctx context.Context, req *v1.QueryStatReq) (res *v1.QueryStatRes, err error)
	ShortDomain(ctx context.Context, req *v1.ShortDomainReq) (res *v1.ShortDomainRes, err error)
}
