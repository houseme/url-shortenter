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

// UserDetailReq is the request struct for user "/api.v1/console/user/detail".
type UserDetailReq struct {
	g.Meta `path:"/user/detail" tags:"User" method:"Get" summary:"获取用户详情"`
	*model.UserDetailInput
}

// UserDetailRes is the response struct for user "/api.v1/console/user/detail".
type UserDetailRes struct {
	*model.UserDetailOutput
}

// UpdateUserReq is the request struct for user "/api.v1/console/user/update".
type UpdateUserReq struct {
	g.Meta `path:"/user/update" tags:"User" method:"Post" summary:"更新用户信息"`
	*model.UpdateUserInput
}

// UpdateUserRes is the response struct for user "/api.v1/console/user/update".
type UpdateUserRes struct {
	*model.UpdateUserOutput
}

// UpdatePasswordReq is the request struct for user "/api.v1/console/user/password".
type UpdatePasswordReq struct {
	g.Meta `path:"/user/password" tags:"User" method:"Post" summary:"修改用户密码"`
	*model.UpdatePasswordInput
}

// UpdatePasswordRes is the response struct for user "/api.v1/console/user/password".
type UpdatePasswordRes struct {
	*model.UpdatePasswordOutput
}
