package echo

import (
	"context"

	v1 "github.com/houseme/url-shortenter/app/console/api/echo/v1"
)

// EchoSay is the handler for Echo Controller action.
func (c *ControllerV1) EchoSay(ctx context.Context, req *v1.EchoSayReq) (res *v1.EchoSayRes, err error) {
	return &v1.EchoSayRes{Content: req.Content}, nil
}
