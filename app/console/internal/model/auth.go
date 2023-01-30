// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package model

// CreateAccessTokenInput is the input for CreateAccessToken
type CreateAccessTokenInput struct {
	AppID     string `json:"appId" dc:"应用ID" v:"required|passport#应用ID唯一标识|应用ID包含字母、数字和下划线，长度在6~18之间"`
	Secret    string `json:"secret" dc:"应用密钥" v:"required|password2#请填写密钥|密钥需要6-18位,必须包含大小写字母和数字"`
	GrantType string `json:"grantType" dc:"授权类型" v:"required|in:client_credentials|授权类型"`
}

// CreateAccessTokenOutput is the output for CreateAccessToken
type CreateAccessTokenOutput struct {
	AccessToken string `json:"accessToken" dc:"授权码 API操作授权码"`
	ExpiresIn   int    `json:"expiresIn" dc:"过期时间"`
}

// AuthorizationToken is the golang structure for AuthorizationToken.
type AuthorizationToken struct {
	AuthToken        string `json:"authToken"`
	AuthTime         int64  `json:"authTime"`
	AuthAccountNo    uint64 `json:"authAccountNo"`
	AuthAccountLevel uint   `json:"authAccountLevel"`
	AuthType         string `json:"authType"`
}

// TokenCache is the golang structure for TokenCache.
type TokenCache struct {
	Token     string              `json:"token"`
	ExpiresIn int                 `json:"expiresIn"`
	AuthToken *AuthorizationToken `json:"authToken"`
}

// AuthInput is the input for Auth
type AuthInput struct {
	Account  string `json:"account" dc:"账号" v:"required|passport#账号唯一标识|账号包含字母、数字和下划线，长度在6~18之间"`
	Password string `json:"password" dc:"密码" v:"required|password2#请填写密码|密码需要6-18位,必须包含大小写字母和数字"`
}

// AuthOutput is the output for Auth
type AuthOutput struct {
	AccessToken string `json:"accessToken" dc:"授权码 登陆之后的授权码"`
	ExpiresIn   int    `json:"expiresIn" dc:"过期时间"`
	RefreshIn   int    `json:"refreshIn" dc:"刷新时间"`
}
