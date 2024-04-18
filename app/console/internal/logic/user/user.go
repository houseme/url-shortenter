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
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/houseme/url-shortenter/app/console/internal/consts"
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

// Update is the handler for Update
func (s *sUser) Update(ctx context.Context, in *model.UpdateUserInput) (out *model.UpdateUserOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-user-Update")
	defer span.End()

	logger := g.Log(helper.Helper().Logger(ctx))
	logger.Debug(ctx, "user-Update in:", in)

	base := (*entity.Users)(nil)
	if err = dao.Users.Ctx(ctx).Scan(&base, do.Users{UserNo: in.AuthUserNo}); err != nil {
		err = gerror.Wrap(err, "dao.Users.Scan failed")
		return
	}
	if base == nil {
		err = gerror.New("user not found")
		return
	}
	doUser := do.Users{
		ModifyTime: gtime.Now(),
	}
	if in.Avatar != "" {
		doUser.Avatar = in.Avatar
	}
	var lastID int64
	if lastID, err = dao.Users.Ctx(ctx).UpdateAndGetAffected(doUser, do.Users{
		UserNo: in.AuthUserNo,
	}); err != nil {
		err = gerror.Wrap(err, "dao.Users.UpdateAndGetAffected failed")
		return
	}
	logger.Debug(ctx, "user-UpdatePassword update user db lastID:", lastID)
	out = &model.UpdateUserOutput{
		ResultCode: consts.ResponseCodeSuccess,
		ResultMsg:  consts.ResponseMsgSuccess,
	}
	return
}

// UpdatePassword is the handler for UpdatePassword
func (s *sUser) UpdatePassword(ctx context.Context, in *model.UpdatePasswordInput) (out *model.UpdatePasswordOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-user-UpdatePassword")
	defer span.End()

	logger := g.Log(helper.Helper().Logger(ctx))
	logger.Debug(ctx, "user-UpdatePassword in:", in)

	base := (*entity.Users)(nil)
	if err = dao.Users.Ctx(ctx).Scan(&base, do.Users{UserNo: in.AuthUserNo}); err != nil {
		err = gerror.Wrap(err, "dao.Users.Scan failed")
		return
	}
	if base == nil {
		err = gerror.New("user not found")
		return
	}
	// 新密码与系统中的就密码效验，如果一样则返回错误
	if err = helper.Helper().VerifyPassword(ctx, base.Password, in.Password, base.Salt); err == nil {
		err = gerror.New("new password is the same as the old password")
		return
	} else {
		logger.Errorf(ctx, "helper.VerifyPassword failed, err:%+v", err)
	}

	var newPass string
	if newPass, err = helper.Helper().GeneratePasswordHash(ctx, in.Password, base.Salt); err != nil {
		err = gerror.Wrap(err, "helper.GeneratePasswordHash failed")
		return
	}

	var lastID int64
	if lastID, err = dao.Users.Ctx(ctx).UpdateAndGetAffected(do.Users{
		Password:   newPass,
		ModifyTime: gtime.Now(),
	}, do.Users{
		UserNo: in.AuthUserNo,
	}); err != nil {
		err = gerror.Wrap(err, "dao.Users.UpdateAndGetAffected failed")
		return
	}
	logger.Debug(ctx, "user-UpdatePassword update user db lastID:", lastID)
	out = &model.UpdatePasswordOutput{
		ResultCode: consts.ResponseCodeSuccess,
		ResultMsg:  consts.ResponseMsgSuccess,
	}

	return
}
