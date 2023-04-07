// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

// Package lark is a utility package for Lark.
package lark

import (
	"context"

	"github.com/chyroc/go-ptr"
	"github.com/chyroc/lark"
	"github.com/gogf/gf/v2/frame/g"
)

type iLark struct {
	ctx  context.Context
	lark *lark.Lark
}

// Lark .
func Lark() *iLark {
	return &iLark{}
}

// NewLark create iLark
func NewLark(ctx context.Context, appID, appSecret, encryptKey, verificationToken string) *iLark {
	return &iLark{
		ctx: ctx,
		lark: lark.New(
			lark.WithAppCredential(appID, appSecret),
			lark.WithEventCallbackVerify(encryptKey, verificationToken),
		),
	}
}

// NewCustomBot new customer bot
func NewCustomBot(ctx context.Context, customBotWebHookURL, customBotSecret string) *iLark {
	g.Log().Debug(ctx, "customBotWebHookURL: ", customBotWebHookURL)
	return &iLark{
		ctx:  ctx,
		lark: lark.New(lark.WithCustomBot(customBotWebHookURL, customBotSecret)),
	}
}

// SendRawMessage  sends raw message
func (s *iLark) SendRawMessage(ctx context.Context) {
	// oc_eaffbb5e6622ea1beccc9751ad7060f4
	//
	resp, res, err := s.lark.Message.SendRawMessage(ctx, &lark.SendRawMessageReq{
		ReceiveIDType: lark.IDTypeChatID,
		ReceiveID:     "oc_eaffbb5e6622ea1beccc9751ad7060f4",
		Content:       `{"text":"测试"}`,
		MsgType:       lark.MsgTypeText,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, resp, res)
}

// SendCustomMessage  sends Custom message
func (s *iLark) SendCustomMessage(ctx context.Context) {
	resp, res, err := s.lark.Message.Send().SendText(ctx, `测试自定义 网页异常 https://www.baidu.com`)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, resp, res)
	g.Log().Info(ctx, ptr.String("success"))
	resp, res, err = s.lark.Message.Send().SendText(ctx, `测试自定义 网页异常 https://www.wasair.com`)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Info(ctx, resp, res)
}

// Lark
