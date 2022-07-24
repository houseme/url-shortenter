package lark

import (
	"context"

	"github.com/houseme/url-shortenter/app/schedule/internal/model"
)

type sLark struct {
}

func new() *sLark {
	return &sLark{}
}

// SendMessage sends a message to a user.
func (s *sLark) SendMessage(ctx context.Context, in *model.SendMessageInput) (out *model.SendMessageOutput, err error) {
	return
}
