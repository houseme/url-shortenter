// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

// Package account logic
package account

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

type sAccount struct{}

func init() {
	service.RegisterAccount(&sAccount{})
}

// CreateAccount is the handler for CreateAccount
func (s *sAccount) CreateAccount(ctx context.Context, in *model.CreateAccountInput) (out *model.CreateAccountOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-account-CreateAccount")
	defer span.End()

	var (
		logger  = g.Log(helper.Helper().Logger(ctx))
		account = (*entity.User)(nil)
		output  = false
	)

	logger.Debug(ctx, "account-CreateAccount in:", in)
	out = (*model.CreateAccountOutput)(&output)
	if in.AuthAccountLevel > consts.AccountLevelBusiness {
		err = gerror.New("Do not have permission to create a new account")
		return
	}

	if in.AuthAccountLevel == consts.AccountLevelPlatform && in.GroupLevel == consts.AccountLevelBusinessEmployee {
		// Platform super administrators prohibit creating ordinary accounts
		err = gerror.New("Platform Super Admin forbids the creation of ordinary accounts")
		return
	}

	// verify data
	if err = dao.User.Ctx(ctx).Scan(&account, do.User{Account: in.Account}); err != nil {
		err = gerror.Wrap(err, "query users failed  err:")
		return
	}

	if account != nil {
		err = gerror.New("account exist")
		return
	}
	// create hash password
	var hashPwd string
	if hashPwd, err = helper.Helper().PasswordBase58Hash(in.Password); err != nil {
		err = gerror.Wrap(err, "hash password failed")
		return
	}

	// create a user
	account = &entity.User{
		Account:    in.Account,
		Password:   hashPwd,
		GroupLevel: in.AuthAccountLevel,
		State:      consts.UserStateNormal,
		UserNo:     helper.Helper().InitTrxID(ctx, in.AuthAccountNo),
	}

	if in.AuthAccountLevel == consts.AccountLevelBusiness {
		account.AccountNo = in.AuthAccountNo
	}
	if in.AuthAccountLevel == consts.AccountLevelPlatform {
		account.AccountNo = helper.Helper().InitTrxID(ctx, in.AuthAccountNo)
	}

	if _, err = dao.User.Ctx(ctx).OmitEmpty().Unscoped().Insert(account); err != nil {
		err = gerror.Wrap(err, "insert users failed")
		return
	}

	output = true
	out = (*model.CreateAccountOutput)(&output)
	logger.Debug(ctx, "account-CreateAccount end out:", out)
	return
}

// ModifyAccount is the handler for ModifyAccount
func (s *sAccount) ModifyAccount(ctx context.Context, in *model.ModifyAccountInput) (out *model.ModifyAccountOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-account-ModifyAccount")
	defer span.End()

	var (
		logger  = g.Log(helper.Helper().Logger(ctx))
		account = (*entity.User)(nil)
	)
	logger.Debug(ctx, "account modify account in:", in)
	if err = dao.User.Ctx(ctx).Scan(&account, do.User{AccountNo: in.AuthAccountNo}); err != nil {
		err = gerror.Wrap(err, "account modify query failed")
		return
	}

	if account == nil {
		err = gerror.New("account is not exists")
		return
	}

	logger.Debug(ctx, "account modify account end out:", out)
	return
}

// ModifyPassword is the handler for ModifyPassword
func (s *sAccount) ModifyPassword(ctx context.Context, in *model.ModifyPasswordInput) (out *model.ModifyPasswordOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-account-ModifyPassword")
	defer span.End()

	var (
		logger  = g.Log(helper.Helper().Logger(ctx))
		account = (*entity.User)(nil)
		output  = false
	)
	out = (*model.ModifyPasswordOutput)(&output)

	logger.Debug(ctx, "account modify password in:", in)
	if err = dao.User.Ctx(ctx).Scan(&account, do.User{AccountNo: in.AuthAccountNo}); err != nil {
		err = gerror.Wrap(err, "query users failed  err:")
		return
	}

	if account == nil {
		err = gerror.New("account not exist")
		return
	}

	var hashPwd string
	if hashPwd, err = helper.Helper().PasswordBase58Hash(in.Password); err != nil {
		err = gerror.Wrap(err, "hash password failed")
		return
	}

	if hashPwd == account.Password {
		err = gerror.New("consistent with the original password")
		return
	}

	logger.Debug(ctx, "account modify password hash password:", hashPwd)
	if _, err = dao.User.Ctx(ctx).Update(do.User{
		Password:   hashPwd,
		ModifyTime: gtime.Now(),
	}, do.User{Id: account.Id}); err != nil {
		err = gerror.Wrap(err, "update users failed")
		return
	}
	output = true
	out = (*model.ModifyPasswordOutput)(&output)
	logger.Debug(ctx, "account modify password end out:", out)
	return
}
