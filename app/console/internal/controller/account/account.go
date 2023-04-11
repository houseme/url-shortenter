/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package account

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtrace"

	v1 "github.com/houseme/url-shortenter/app/console/api/account/v1"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

type cAccount struct {
}

// New is the constructor for cAccount
func New() *cAccount {
	return &cAccount{}
}

// CreateAccount is the handler for CreateAccount
func (c *cAccount) CreateAccount(ctx context.Context, req *v1.CreateAccountReq) (res *v1.CreateAccountRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-account-CreateAccount")
	defer span.End()

	res = &v1.CreateAccountRes{}
	if res.CreateAccountOutput, err = service.Account().CreateAccount(ctx, req.CreateAccountInput); err != nil {
		err = gerror.Wrap(err, "controller create account failed")
	}
	return
}

// ModifyAccount is the handler for ModifyAccount
func (c *cAccount) ModifyAccount(ctx context.Context, req *v1.ModifyAccountReq) (res *v1.ModifyAccountRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-account-ModifyAccount")
	defer span.End()

	res = &v1.ModifyAccountRes{}
	if res.ModifyAccountOutput, err = service.Account().ModifyAccount(ctx, req.ModifyAccountInput); err != nil {
		err = gerror.Wrap(err, "controller modify account failed")
	}
	return
}

// ModifyPassword is the handler for ModifyPassword
func (c *cAccount) ModifyPassword(ctx context.Context, req *v1.ModifyPasswordReq) (res *v1.ModifyPasswordRes, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-controller-account-ModifyPassword")
	defer span.End()

	res = &v1.ModifyPasswordRes{}
	if res.ModifyPasswordOutput, err = service.Account().ModifyPassword(ctx, req.ModifyPasswordInput); err != nil {
		err = gerror.Wrap(err, "controller modify password failed")
	}
	return
}
