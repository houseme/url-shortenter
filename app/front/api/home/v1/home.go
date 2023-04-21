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

	"github.com/houseme/url-shortenter/app/front/internal/model"
)

// Req is the request struct for :short.
type Req struct {
	g.Meta `path:"/:short" tags:"home" method:"get" summary:"You first home api"`
	*model.HomeInput
}

// Res is the response struct.
type Res string
