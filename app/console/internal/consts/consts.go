// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package consts

const (
	// AppID 应用 ID
	AppID = `console`

	// DefaultLogger 应用默认日志名称
	DefaultLogger = `console`

	// AuthorizationHeaderKey 授权头
	AuthorizationHeaderKey = "authorization"
	// AuthorizationTypeBearer 授权类型
	AuthorizationTypeBearer = "Bearer"

	// AccountLevelPlatform 超级管理员 权限级别 0 默认平台超级管理员;1000 商家管理员;10000 商家员工
	AccountLevelPlatform = 0
	// AccountLevelBusiness 商家管理员
	AccountLevelBusiness = 1000
	// AccountLevelBusinessEmployee 商家员工
	AccountLevelBusinessEmployee = 10000

	// UserStateDefault 用户状态 0 默认，100 正常，200 失效
	UserStateDefault = 0
	// UserStateNormal 正常
	UserStateNormal = 100
	// UserStateInvalid 失效
	UserStateInvalid = 200

	// AuthSecretStateDefault 授权密钥状态 0 默认 100 正常 200 禁用
	AuthSecretStateDefault = 0
	// AuthSecretStateNormal 正常
	AuthSecretStateNormal = 100
	// AuthSecretStateInvalid 失效
	AuthSecretStateInvalid = 200

	// AuthTypeAPIKey 授权类型 ApiKey 授权类型 账户密码授权类型
	AuthTypeAPIKey = "api_key"
	// AuthTypePassword 授权类型 账户密码授权类型
	AuthTypePassword = "password"

	// APIKeyExpireTime ApiKey 授权有效期时间 单位秒
	APIKeyExpireTime = 7200

	// PasswordExpireTime 账号密码授权有效期时间 单位秒
	PasswordExpireTime = 7200

	// AccessTokenExpireTime 访问令牌有效期时间 单位秒
	AccessTokenExpireTime = 7200

	// RefreshTokenExpireTime 刷新令牌有效期时间 单位秒
	RefreshTokenExpireTime = 864000

	// TokenExpireTime 有效期时间 单位秒
	TokenExpireTime = 864000

	// ShortDomainStateDefault 短链域名状态  0 默认 100 正常 200 失效
	ShortDomainStateDefault = 0
	// ShortDomainStateNormal 正常
	ShortDomainStateNormal = 100
	// ShortDomainStateInvalid 失效
	ShortDomainStateInvalid = 200
)
