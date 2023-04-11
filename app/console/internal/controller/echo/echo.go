/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package echo

import (
	"context"

	v1 "github.com/houseme/url-shortenter/app/console/api/echo/v1"
)

// New is the constructor of cEcho.
func New() *cEcho {
	return &cEcho{}
}

type cEcho struct{}

// Say is the handler for Say
func (c *cEcho) Say(ctx context.Context, req *v1.EchoSayReq) (res *v1.EchoSayRes, err error) {
	return &v1.EchoSayRes{Content: req.Content}, nil
}
