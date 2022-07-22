package lark

import (
	"context"
	"fmt"

	"github.com/chyroc/go-ptr"
	"github.com/chyroc/lark"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	appID             = "cli_a1d80430c1fXXXXXX"
	appSecret         = "tIVGn0avbrYN0nQCZnHb4g21XXXXXXXXX"
	encryptKey        = ""
	verificationToken = "oZwu6fN7TeNtlJKRxnousbZkXXXXXXXX"
)

type iLark struct {
	lark *lark.Lark
}

// Lark .
func Lark() *iLark {
	return &iLark{}
}

// NewLark create iLark
func NewLark(ctx context.Context, appID, appSecret, encryptKey, verificationToken string) *iLark {
	return &iLark{
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
	fmt.Println(resp, res, err)
}

// SendCustomMessage  sends Custom message
func (s *iLark) SendCustomMessage(ctx context.Context) {
	resp, res, err := s.lark.Message.Send().SendText(ctx, `测试自定义 网页异常 https://www.baidu.com`)
	fmt.Println(resp, res, err)
	fmt.Println(ptr.String("success"))
	resp, res, err = s.lark.Message.Send().SendText(ctx, `测试自定义 网页异常 https://www.wasair.com`)
	fmt.Println(resp, res, err)
}

// Lark
