// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package auth

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"

	"github.com/houseme/url-shortenter/app/console/internal/consts"
	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
	"github.com/houseme/url-shortenter/internal/database/dao"
	"github.com/houseme/url-shortenter/internal/database/model/do"
	"github.com/houseme/url-shortenter/internal/database/model/entity"
	"github.com/houseme/url-shortenter/utility/cache"
	"github.com/houseme/url-shortenter/utility/helper"
)

type sAuth struct {
}

func init() {
	service.RegisterAuth(&sAuth{})
}

// CreateAccessToken creates an initAuth access token.
func (s *sAuth) CreateAccessToken(ctx context.Context, in *model.CreateAccessTokenInput) (out *model.CreateAccessTokenOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-auth-CreateAccessToken")
	defer span.End()

	var (
		log          = g.Log(helper.Helper().Logger(ctx))
		account      = (*entity.Users)(nil)
		accessSecret = (*entity.UsersAccessSecret)(nil)
	)
	log.Debug(ctx, "auth create access token AppID:", in.AppID)
	if err = dao.UsersAccessSecret.Ctx(ctx).Scan(&account, do.UsersAccessSecret{
		SecretId:  in.AppID,
		GrantType: gstr.ToLower(in.GrantType),
	}); err != nil {
		err = gerror.Wrap(err, "query UsersAccessSecret failed  err:")
	}
	if accessSecret == nil {
		err = gerror.New("AppID not found")
		return
	}

	if accessSecret.State == consts.AuthSecretStateInvalid {
		err = gerror.New("AppID is invalid")
		return
	}

	var aesHash string
	if aesHash, err = helper.Helper().AESEncrypt(ctx, []byte(accessSecret.SaltKey), []byte(accessSecret.Salt+in.Secret)); err != nil {
		err = gerror.Wrap(err, "AESEncrypt failed")
		return
	}

	if aesHash != accessSecret.SecretKey {
		err = gerror.New("SecretKey is invalid")
		return
	}
	// salt 16 位 salt key 16 位 需要加密的内容位 salt+secret aes 加密之后于数据库对比  检验完成 处理 accessToken 相关的处理
	// 创建 accessToken
	if err = dao.Users.Ctx(ctx).Scan(&account, do.Users{
		AccountNo:  accessSecret.AccountNo,
		GroupLevel: consts.AccountLevelBusiness,
	}); err != nil {
		err = gerror.Wrap(err, "query Users failed  err:")
		return
	}

	if account == nil {
		err = gerror.New("AccountNo not found")
		return
	}

	if account.State != consts.UserStateNormal {
		err = gerror.New("AccountNo is invalid")
		return
	}

	var (
		authToken = &model.AuthorizationToken{
			AuthAccountNo:    account.AccountNo,
			AuthAccountLevel: account.GroupLevel,
			AuthType:         consts.AuthTypeAPIKey,
			AuthTime:         gtime.Now().Unix(),
		}
		v     *gvar.Var
		token string
	)
	if token, err = helper.Helper().CreateAccessToken(ctx, account.AccountNo); err != nil {
		err = gerror.Wrap(err, "CreateAccessToken failed")
		return
	}
	authToken.AuthToken = token
	if v, err = s.setRedisToken(ctx,
		&model.TokenCache{
			Token:     token,
			ExpiresIn: consts.APIKeyExpireTime + grand.N(10, 50),
			AuthToken: authToken,
		}); err != nil {
		err = gerror.Wrap(err, "setRedisToken failed")
		return
	}

	log.Debug(ctx, "auth-CreateAccessToken v1:", v)
	out = &model.CreateAccessTokenOutput{
		AccessToken: token,
		ExpiresIn:   consts.AccessTokenExpireTime,
	}
	return
}

// Authorization authorizes the user.
func (s *sAuth) Authorization(ctx context.Context, in *model.AuthInput) (out *model.AuthOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-auth-authorization")
	defer span.End()

	var (
		log     = g.Log(helper.Helper().Logger(ctx))
		account = (*entity.Users)(nil)
	)
	log.Debug(ctx, "auth-authorization params account:", in.Account)
	if err = dao.Users.Ctx(ctx).Scan(&account, do.Users{AccountNo: in.Account}); err != nil {
		err = gerror.Wrap(err, "query Users failed  err:")
		return
	}

	if account == nil {
		err = gerror.New("AccountNo not found")
		return
	}
	if account.State != consts.UserStateNormal {
		err = gerror.New("Account is invalid")
		return
	}
	// 比对密码
	var shaHash string
	if shaHash, err = helper.Helper().PasswordBase58Hash(in.Password); err != nil {
		err = gerror.Wrap(err, "PasswordBase58Hash failed")
		return
	}
	if shaHash != account.Password {
		err = gerror.New("Password is invalid")
		return
	}
	// 检验完成 处理 accessToken 相关的处理
	// 创建 accessToken
	var (
		authToken = &model.AuthorizationToken{
			AuthAccountNo:    account.AccountNo,
			AuthAccountLevel: account.GroupLevel,
			AuthType:         consts.AuthTypePassword,
			AuthTime:         gtime.Now().Unix(),
		}
		v     *gvar.Var
		token string
	)
	if token, err = helper.Helper().CreateAccessToken(ctx, account.AccountNo); err != nil {
		err = gerror.Wrap(err, "CreateAccessToken failed")
		return
	}
	authToken.AuthToken = token

	if v, err = s.setRedisToken(ctx,
		&model.TokenCache{
			Token:     token,
			ExpiresIn: consts.AccessTokenExpireTime + grand.N(10, 50),
			AuthToken: authToken,
		}); err != nil {
		err = gerror.Wrap(err, "Redis SETEX failed")
		return
	}
	log.Debug(ctx, "auth-authorization v1:", v)
	out = &model.AuthOutput{
		AccessToken: token,
		ExpiresIn:   consts.AccessTokenExpireTime,
		RefreshIn:   consts.RefreshTokenExpireTime,
	}
	return
}

// SetRedisToken set redis token.
func (s *sAuth) setRedisToken(ctx context.Context, in *model.TokenCache) (val *gvar.Var, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-auth-setRedisToken")
	defer span.End()

	var (
		log       = g.Log(helper.Helper().Logger(ctx))
		redisName = cache.RedisCache().ShortAccessTokenConn(ctx)
		redisKey  = cache.RedisCache().ShortAuthorizationKey(ctx, in.Token)
	)
	log.Debug(ctx, "auth-setRedisToken params account:", in)
	if val, err = g.Redis(redisName).Do(ctx, "SETEX", redisKey, in.ExpiresIn, in.AuthToken); err != nil {
		err = gerror.Wrap(err, "Redis SETEX failed")
		return
	}
	return
}
