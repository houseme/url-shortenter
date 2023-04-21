/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Req is a hello request.
type Req struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
}

// Res is a hello response.
type Res struct {
	g.Meta `mime:"text/html" example:"string"`
}
