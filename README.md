# URL-Shortenter

[![Go Reference](https://pkg.go.dev/badge/github.com/houseme/url-shortenter.svg)](https://pkg.go.dev/github.com/houseme/url-shortenter)
[![Go-Url-Shortenter CI](https://github.com/houseme/url-shortenter/actions/workflows/go.yml/badge.svg)](https://github.com/houseme/url-shortenter/actions/workflows/go.yml)
[![License](https://img.shields.io/github/license/houseme/url-shortenter.svg?style=flat)](https://github.com/houseme/url-shortenter)
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/houseme/url-shortenter/main)

A short link service system suitable for small and medium-sized community websites.

Support short link production, query and 302 redirection, and have its own click statistics, independent IP statistics,
access log query.

## 快速使用

```shell
go install -u -v github.com/houseme/url-shortenter@latest
```

## Console 后台默认帐号

默认帐号: `urlShortener`  
默认密码: `B9Mazv5M2J6%1zU2@nxC`

数据库中存储的是加密后的密码，在 `document\structure.sql` 中标有注释，如果需要自定义其他密码，可以修改这里

加密规则 `utility/helper.go` 中

```go 
func (u *utilHelper) PasswordBase58Hash(password string) (string, error) {
	data, err := u.Sha256OfShort(password)
	if err != nil {
		err = gerror.Wrap(err, "utilHelper PasswordBase58Hash Sha256OfShort error")
		return "", err
	}
	return u.Base58Encode(data), nil
}
```

## HTTP API 支持

### `/api` 接口权限说明

所有 `/console/api/*` 接口需要通过 `Bearer Token` 方式验证权限，亦即：每个请求 Header 须携带

```shell
 Authorization: Bearer {sha256_of_password}
```

### 1. 新增短链接 `POST /api/url`

## 短链接生产过程相关代码

所在文件 `utility/helper.go`

```go
func (u *utilHelper) GenerateShortLink(ctx context.Context, url string) (string, error) {
	var (
		err     error
		urlHash []byte
		logger  = u.Logger(ctx)
	)
	g.Log(logger).Debug(ctx, "utilHelper GenerateShortLink url:", url)
	if urlHash, err = u.Sha256OfShort(url); err != nil {
		err = gerror.Wrap(err, "utilHelper GenerateShortLink Sha256OfShort err")
		return "", err
	}
	number := new(big.Int).SetBytes(urlHash).Uint64()
	str := u.Base58Encode(gconv.Bytes(number))
	g.Log(logger).Debug(ctx, "utilHelper GenerateShortLink str:", str, " number:", number)
	return str[:8], nil
}
```
