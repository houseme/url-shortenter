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

	"github.com/houseme/url-shortenter/app/console/internal/model"
)

// HomeIndexReq is the request struct for the HomeIndex endpoint.
type HomeIndexReq struct {
	g.Meta `path:"/index" tags:"Home Service" method:"get" summary:"You say, I echo"`
	*model.HomeIndexInput
}

// HomeIndexRes is the response struct for the HomeIndex endpoint.
type HomeIndexRes struct {
	*model.HomeIndexOutput
}
