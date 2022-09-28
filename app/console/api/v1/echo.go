// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

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
