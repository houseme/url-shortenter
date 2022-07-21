package controller

import (
	"context"

	v1 "github.com/houseme/url-shortenter/app/console/api/v1"
)

var (
	Echo = cEcho{}
)

type cEcho struct{}

func (c *cEcho) Say(ctx context.Context, req *v1.EchoSayReq) (res *v1.EchoSayRes, err error) {
	return &v1.EchoSayRes{Content: req.Content}, nil
}
