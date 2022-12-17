package lark

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

// Env lark environment.
type Env struct {
	appID               string
	appSecret           string
	encryptKey          string
	verificationToken   string
	customBotWebHookURL string
	customBotSecret     string
}

// New .create a new lark environment
func New(ctx context.Context) (*Env, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-lark-New")
	defer span.End()

	var v, err = g.Cfg().Get(ctx, "lark")
	if err != nil {
		err = gerror.Wrap(err, "config lark get failed")
		return nil, err
	}

	if v == nil || v.IsNil() || v.IsEmpty() {
		err = gerror.New("config lark is empty")
		return nil, err
	}

	var config = v.MapStrStr()
	return &Env{
		appID:               config["appID"],
		appSecret:           config["appSecret"],
		encryptKey:          config["encryptKey"],
		verificationToken:   config["verificationToken"],
		customBotWebHookURL: config["customBotWebHookURL"],
		customBotSecret:     config["customBotSecret"],
	}, nil
}

// APPID .
func (e *Env) APPID(_ context.Context) string {
	return e.appID
}

// APPSecret .
func (e *Env) APPSecret(_ context.Context) string {
	return e.appSecret
}

// EncryptKey .
func (e *Env) EncryptKey(_ context.Context) string {
	return e.encryptKey
}

// VerificationToken .
func (e *Env) VerificationToken(_ context.Context) string {
	return e.verificationToken
}

// CustomBotWebHookURL .
func (e *Env) CustomBotWebHookURL(_ context.Context) string {
	return e.customBotWebHookURL
}

// CustomBotSecret .
func (e *Env) CustomBotSecret(_ context.Context) string {
	return e.customBotSecret
}

// String returns the string representation of the environment.
func (e *Env) String(_ context.Context) string {
	return `{"appID":"` + e.appID + `","appSecret":"` + e.appSecret + `","encryptKey":"` + e.encryptKey +
		`","verificationToken":"` + e.verificationToken + `","customBotWebHookURL":"` + e.customBotWebHookURL +
		`","customBotSecret":"` + e.customBotSecret + `"}`
}
