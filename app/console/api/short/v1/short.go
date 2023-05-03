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

// CreateShortReq is the request struct for the CreateShort endpoint.
type CreateShortReq struct {
	g.Meta `path:"/url" tags:"Account Service" method:"Post" summary:"create a short url"`
	*model.CreateShortInput
}

// CreateShortRes is the response struct for the CreateShort endpoint.
type CreateShortRes struct {
	*model.CreateShortOutput
}

// QueryShortReq is the request struct for the QueryShort endpoint.
type QueryShortReq struct {
	g.Meta `path:"/url/:shortUrl" tags:"Account Service" method:"Get" summary:"query a short url"`
	*model.QueryShortInput
}

// QueryShortRes is the response struct for the QueryShort endpoint.
type QueryShortRes struct {
	*model.QueryShortOutput
}

// ModifyShortReq is the request struct for the ModifyShort endpoint.
type ModifyShortReq struct {
	g.Meta `path:"/url/:shortUrl/change_state" tags:"Account Service" method:"Post" summary:"modify a short url"`
	*model.ModifyShortInput
}

// ModifyShortRes is the response struct for the ModifyShort endpoint.
type ModifyShortRes struct {
	*model.ModifyShortOutput
}

// QueryStatReq is the request struct for the QueryStat endpoint.
type QueryStatReq struct {
	g.Meta `path:"/url/:shortUrl/stat" tags:"Account Service" method:"Get" summary:"query a short url stat"`
	*model.QueryStatInput
}

// QueryStatRes is the response struct for the QueryStat endpoint.
type QueryStatRes struct {
	*model.QueryStatOutput
}

// ShortDomainReq is the request struct for the ShortDomain endpoint.
type ShortDomainReq struct {
	g.Meta `path:"/url/short_domain" tags:"Account Service" method:"Get" summary:"query short domain"`
	*model.ShortDomainInput
}

// ShortDomainRes is the response struct for the ShortDomain endpoint.
type ShortDomainRes struct {
	*model.ShortDomainOutput
}
