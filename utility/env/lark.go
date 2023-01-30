package env

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
)

// LarkEnv lark environment.
type LarkEnv struct {
	appID               string
	appSecret           string
	encryptKey          string
	verificationToken   string
	customBotWebHookURL string
	customBotSecret     string
}

// NewLark .create a new lark environment
func NewLark(ctx context.Context) (*LarkEnv, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-lark-NewLark")
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
	return &LarkEnv{
		appID:               config["appID"],
		appSecret:           config["appSecret"],
		encryptKey:          config["encryptKey"],
		verificationToken:   config["verificationToken"],
		customBotWebHookURL: config["customBotWebHookURL"],
		customBotSecret:     config["customBotSecret"],
	}, nil
}

// APPID .
func (e *LarkEnv) APPID(_ context.Context) string {
	return e.appID
}

// APPSecret .
func (e *LarkEnv) APPSecret(_ context.Context) string {
	return e.appSecret
}

// EncryptKey .
func (e *LarkEnv) EncryptKey(_ context.Context) string {
	return e.encryptKey
}

// VerificationToken .
func (e *LarkEnv) VerificationToken(_ context.Context) string {
	return e.verificationToken
}

// CustomBotWebHookURL .
func (e *LarkEnv) CustomBotWebHookURL(_ context.Context) string {
	return e.customBotWebHookURL
}

// CustomBotSecret .
func (e *LarkEnv) CustomBotSecret(_ context.Context) string {
	return e.customBotSecret
}

// String returns the string representation of the environment.
func (e *LarkEnv) String(_ context.Context) string {
	return `{"appID":"` + e.appID + `","appSecret":"` + e.appSecret + `","encryptKey":"` + e.encryptKey +
		`","verificationToken":"` + e.verificationToken + `","customBotWebHookURL":"` + e.customBotWebHookURL +
		`","customBotSecret":"` + e.customBotSecret + `"}`
}
