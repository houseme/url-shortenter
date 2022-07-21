package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type EchoSayReq struct {
	g.Meta  `path:"/echo/say" tags:"Echo Service" method:"get" summary:"You say, I echo"`
	Content string `v:"required" dc:"Say something?"`
}

type EchoSayRes struct {
	Content string `dc:"Reply content"`
}
