// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package lark

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/houseme/url-shortenter/app/schedule/internal/model"
	"github.com/houseme/url-shortenter/internal/webhooks/lark"
)

// SendMessage sends a message to a user.
func (s *sLark) SendMessage(ctx context.Context, in *model.SendMessageInput) (out *model.SendMessageOutput, err error) {
	g.Log().Debug(ctx, "SendMessage params:", in)

	lark.NewCustomBot(ctx, "", "").SendRawMessage(ctx)

	return
}
