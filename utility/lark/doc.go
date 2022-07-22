package lark

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"
)

type larkEnv struct {
	appID               string `json:"appID"`
	appSecret           string `json:"appSecret"`
	encryptKey          string `json:"encryptKey"`
	verificationToken   string `json:"verificationToken"`
	customBotWebHookURL string `json:"customBotWebHookURL"`
	customBotSecret     string `json:"customBotSecret"`
}

// New .create a new lark environment
func New(ctx context.Context) *larkEnv {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-lark-New")
	defer span.End()

	var (
		le     *larkEnv
		logger = gconv.String(ctx.Value("logger"))
	)
	if err := g.Cfg().MustGet(ctx, "lark").Scan(&le); err != nil {
		g.Log(logger).Error(ctx, " config lark fail err:", err)
		span.RecordError(err)
		return nil
	}
	if le == nil {
		g.Log(logger).Error(ctx, " config lark is empty")
		return nil
	}

	return le
}

// APPID .
func (l *larkEnv) APPID(ctx context.Context) string {
	return l.appID
}

// APPSecret .
func (l *larkEnv) APPSecret(ctx context.Context) string {
	return l.appSecret
}

// EncryptKey .
func (l *larkEnv) EncryptKey(ctx context.Context) string {
	return l.encryptKey
}

// VerificationToken .
func (l *larkEnv) VerificationToken(ctx context.Context) string {
	return l.verificationToken
}

// CustomBotWebHookURL .
func (l *larkEnv) CustomBotWebHookURL(ctx context.Context) string {
	return l.customBotWebHookURL
}

// CustomBotSecret .
func (l *larkEnv) CustomBotSecret(ctx context.Context) string {
	return l.customBotSecret
}

// String returns the string representation of the environment.
func (l *larkEnv) String() string {
	return `{"appID":"` + l.appID + `","appSecret":"` + l.appSecret + `","encryptKey":"` + l.encryptKey + `","verificationToken":"` + l.verificationToken + `","customBotWebHookURL":"` + l.customBotWebHookURL + `","customBotSecret":"` + l.customBotSecret + `"}`
}
