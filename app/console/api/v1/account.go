// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/houseme/url-shortenter/app/console/internal/model"
)

// CreateAccountReq is the request struct for the CreateAccount endpoint.
type CreateAccountReq struct {
	g.Meta `path:"/account" tags:"Account Service" method:"Post" summary:"create an account"`
	*model.CreateAccountInput
}

// CreateAccountRes is the response struct for the CreateAccount endpoint.
type CreateAccountRes struct {
	*model.CreateAccountOutput
}

// ModifyAccountReq is the request struct for the ModifyAccount endpoint.
type ModifyAccountReq struct {
	g.Meta `path:"/account/:account/update" tags:"Account Service" method:"Post" summary:"modify an account"`
	*model.ModifyAccountInput
}

// ModifyAccountRes is the response struct for the ModifyAccount endpoint.
type ModifyAccountRes struct {
	*model.ModifyAccountOutput
}

// ModifyPasswordReq is the request struct for the ModifyPassword endpoint.
type ModifyPasswordReq struct {
	g.Meta `path:"/account/:account/password" tags:"Account Service" method:"Post" summary:"modify password"`
	*model.ModifyPasswordInput
}

// ModifyPasswordRes is the response struct for the ModifyPassword endpoint.
type ModifyPasswordRes struct {
	*model.ModifyPasswordOutput
}
