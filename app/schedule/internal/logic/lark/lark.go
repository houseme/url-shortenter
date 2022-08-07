package lark

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/houseme/url-shortenter/app/schedule/internal/model"
	"github.com/houseme/url-shortenter/app/schedule/internal/service"
)

type sLark struct {
}

func init() {
	service.RegisterLark(initLark())
}

func initLark() *sLark {
	return &sLark{}
}

// SendMessage sends a message to a user.
func (s *sLark) SendMessage(ctx context.Context, in *model.SendMessageInput) (out *model.SendMessageOutput, err error) {
	g.Log().Debug(ctx, "SendMessage params:", in)
	return
}
