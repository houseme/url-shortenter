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

// CreateTagReq is the request struct for the CreateTag endpoint.
// 创建标签请求 如果存在直接返回标签，如果不存在则创建标签
type CreateTagReq struct {
	g.Meta `path:"/tag" tags:"Account Service" method:"Post" summary:"create a tag"`
	*model.CreateTagInput
}

// CreateTagRes is the response struct for the CreateTag endpoint.
type CreateTagRes struct {
	*model.CreateTagOutput
}

// DelTagReq is the request struct for the DelTag endpoint.
// 删除短链标签请求
type DelTagReq struct {
	g.Meta `path:"/tag/del" tags:"Account Service" method:"Post" summary:"delete a tag"`
	*model.DelTagInput
}

// DelTagRes is the response struct for the DelTag endpoint.
type DelTagRes struct {
	*model.DelTagOutput
}

// AddTagReq is the request struct for the AddTag endpoint.
// 用于以及创建的短链添加标签
type AddTagReq struct {
	g.Meta `path:"/tag/add" tags:"Account Service" method:"Post" summary:"add a tag"`
	*model.AddTagInput
}

// AddTagRes is the response struct for the AddTag endpoint.
type AddTagRes struct {
	*model.AddTagOutput
}
