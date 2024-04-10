// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package echo

import (
	"context"

	"github.com/houseme/url-shortenter/app/console/api/echo/v1"
)

type IEchoV1 interface {
	EchoSay(ctx context.Context, req *v1.EchoSayReq) (res *v1.EchoSayRes, err error)
}
