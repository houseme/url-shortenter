// Package lark is a lark service.
package lark

import (
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
