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

// CreateAccessTokenReq is the request struct for the CreateAccessToken endpoint.
type CreateAccessTokenReq struct {
	g.Meta `path:"/auth/access_token" tags:"Account Service" method:"Get" summary:"create an access token"`
	*model.CreateAccessTokenInput
}

// CreateAccessTokenRes is the response struct for the CreateAccessToken endpoint.
type CreateAccessTokenRes struct {
	*model.CreateAccessTokenOutput
}

// AuthReq is the request struct for the Auth endpoint.
type AuthReq struct {
	g.Meta `path:"/auth/authorization" tags:"Account Service" method:"Post" summary:"authorization"`
	*model.AuthInput
}

// AuthRes is the response struct for the Auth endpoint.
type AuthRes struct {
	*model.AuthOutput
}
