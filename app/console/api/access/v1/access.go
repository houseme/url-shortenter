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

// ListReq is the input structure for the List API.
type ListReq struct {
	g.Meta `path:"/access/list" tags:"Access log" method:"Get" summary:"query access log list"`
	*model.AccessListInput
}

// ListRes is the output structure for the List API.
type ListRes struct {
	*model.AccessListOutput
}

// DetailReq is the input structure for the Detail API.
type DetailReq struct {
	g.Meta `path:"/access/detail" tags:"Access log" method:"Get" summary:"query access log detail"`
	*model.AccessDetailInput
}

// DetailRes is the output structure for the Detail API.
type DetailRes struct {
	*model.AccessDetailOutput
}
