package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// EchoSayReq is the request struct for the EchoSay endpoint.
type EchoSayReq struct {
	g.Meta  `path:"/echo/say" tags:"Echo Service" method:"get" summary:"You say, I echo"`
	Content string `v:"required" dc:"Say something?"`
}

// EchoSayRes is the response struct for the EchoSay endpoint.
type EchoSayRes struct {
	Content string `dc:"Reply content"`
}
