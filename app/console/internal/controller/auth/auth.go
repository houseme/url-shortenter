/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/auth/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

// Controller is the controller for the auth.
type Controller struct {
}

// New is the constructor of Controller.
func New() *Controller {
	return &Controller{}
}

// CreateAccessToken is the handler for CreateAccessToken
func (c *Controller) CreateAccessToken(ctx context.Context, req *v1.CreateAccessTokenReq) (res *v1.CreateAccessTokenRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-auth-CreateAccessToken")
	defer span.End()

	res = &v1.CreateAccessTokenRes{}
	if res.CreateAccessTokenOutput, err = service.Auth().CreateAccessToken(ctx, req.CreateAccessTokenInput); err != nil {
		err = gerror.Wrap(err, "auth-CreateAccessToken failed")
	}
	return
}

// Authorization is the handler for Authorization
func (c *Controller) Authorization(ctx context.Context, req *v1.AuthReq) (res *v1.AuthRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-auth-authorization")
	defer span.End()

	res = &v1.AuthRes{}
	if res.AuthOutput, err = service.Auth().Authorization(ctx, req.AuthInput); err != nil {
		err = gerror.Wrap(err, "auth-authorization failed")
	}
	return
}
