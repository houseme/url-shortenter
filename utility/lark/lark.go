package lark

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/util/gconv"
)

// Env lark environment.
type Env struct {
	ctx                 context.Context
	appID               string
	appSecret           string
	encryptKey          string
	verificationToken   string
	customBotWebHookURL string
	customBotSecret     string
}

// New .create a new lark environment
func New(ctx context.Context) *Env {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-lark-New")
	defer span.End()

	var (
		e      *Env
		logger = gconv.String(ctx.Value("logger"))
	)
	if err := g.Cfg().MustGet(ctx, "lark").Scan(&e); err != nil {
		g.Log(logger).Error(ctx, " config lark fail err:", err)
		span.RecordError(err)
		return nil
	}
	if e == nil {
		g.Log(logger).Error(ctx, " config lark is empty")
		return nil
	}
	e.ctx = ctx

	return e
}

// APPID .
func (l *Env) APPID(ctx context.Context) string {
	l.ctx = ctx
	return l.appID
}

// APPSecret .
func (l *Env) APPSecret(ctx context.Context) string {
	l.ctx = ctx
	return l.appSecret
}

// EncryptKey .
func (l *Env) EncryptKey(ctx context.Context) string {
	l.ctx = ctx
	return l.encryptKey
}

// VerificationToken .
func (l *Env) VerificationToken(ctx context.Context) string {
	l.ctx = ctx
	return l.verificationToken
}

// CustomBotWebHookURL .
func (l *Env) CustomBotWebHookURL(ctx context.Context) string {
	l.ctx = ctx
	return l.customBotWebHookURL
}

// CustomBotSecret .
func (l *Env) CustomBotSecret(ctx context.Context) string {
	l.ctx = ctx
	return l.customBotSecret
}

// String returns the string representation of the environment.
func (l *Env) String(ctx context.Context) string {
	l.ctx = ctx
	return `{"appID":"` + l.appID + `","appSecret":"` + l.appSecret + `","encryptKey":"` + l.encryptKey +
		`","verificationToken":"` + l.verificationToken + `","customBotWebHookURL":"` + l.customBotWebHookURL +
		`","customBotSecret":"` + l.customBotSecret + `"}`
}
