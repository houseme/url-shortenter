// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

// Package user logic
package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
	"github.com/houseme/url-shortenter/internal/database/dao"
	"github.com/houseme/url-shortenter/internal/database/model/do"
	"github.com/houseme/url-shortenter/internal/database/model/entity"
	"github.com/houseme/url-shortenter/utility/helper"
)

type sUser struct{}

func init() {
	service.RegisterUser(&sUser{})
}

// CreateMerchant creates a new merchant.
func (s *sUser) CreateMerchant(ctx context.Context, in *model.CreateMerchantInput) (out *model.CreateMerchantOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-user-CreateMerchant")
	defer span.End()

	logger := g.Log(helper.Helper().Logger(ctx))
	logger.Debug(ctx, "user-CreateMerchant in:", in)

	return
}

// QueryMerchant queries merchant by id.
func (s *sUser) QueryMerchant(ctx context.Context, in *model.QueryMerchantInput) (out *model.QueryMerchantOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-user-QueryMerchant")
	defer span.End()

	var (
		logger   = g.Log(helper.Helper().Logger(ctx))
		merchant = (*entity.UsersMerchant)(nil)
	)
	logger.Debug(ctx, "user-QueryMerchant in:", in)
	if err = dao.UsersMerchant.Ctx(ctx).Scan(&merchant, do.UsersMerchant{AccountNo: in.AuthAccountNo}); err != nil {
		err = gerror.Wrap(err, "dao.UsersMerchant.Scan failed")
		return
	}
	if merchant == nil {
		err = gerror.New("merchant not found")
		return
	}

	return
}

// Detail is the handler for Detail
func (s *sUser) Detail(ctx context.Context, in *model.UserDetailInput) (out *model.UserDetailOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-user-Detail")
	defer span.End()

	var (
		logger = g.Log(helper.Helper().Logger(ctx))
		base   = (*entity.Users)(nil)
	)
	logger.Debug(ctx, "user-Detail in:", in)
	if err = dao.Users.Ctx(ctx).Scan(&base, do.Users{UserNo: in.AuthUserNo}); err != nil {
		err = gerror.Wrap(err, "dao.Users.Scan failed")
		return
	}

	if base == nil {
		err = gerror.New("user not found")
		return
	}
	out = &model.UserDetailOutput{
		Username: "",
		Avatar:   "",
	}

	return
}
