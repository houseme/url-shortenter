# URL-Shortenter 

English｜[中文文档](README-ZH.md)

[![Go Reference](https://pkg.go.dev/badge/github.com/houseme/url-shortenter.svg)](https://pkg.go.dev/github.com/houseme/url-shortenter)
[![Url-Shortenter CI](https://github.com/houseme/url-shortenter/actions/workflows/go.yml/badge.svg)](https://github.com/houseme/url-shortenter/actions/workflows/go.yml)
[![License](https://img.shields.io/github/license/houseme/url-shortenter.svg?style=flat)](https://github.com/houseme/url-shortenter)
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/houseme/url-shortenter/main)

A short link service system suitable for small and medium-sized community websites.

Supports short link generation, query, 302 redirection, click statistics, independent IP statistics, and access log query.

## Quick Start

```shell
go install -u -v github.com/houseme/url-shortenter@latest
```

## Console Default Account

Default account: `shortenter`  
Default password: `B9Mazv5M2J6%1zU2@nxC`

The database stores encrypted passwords, which are annotated in `document\structure.sql`. If you need to customize other passwords, you can modify them here.

Encryption rules are in `utility/helper.go`.

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

## HTTP API Support

### `/api` Interface Permission Description

All `/console/api/*` interfaces require `Bearer Token` authentication. That is, each request header must include:

```shell
 Authorization: Bearer {sha256_of_password}
```

### 1. Add Short Link `POST /api/url`

## Short Link Generation Code

Located in `utility/helper.go`.

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
	// number := new(big.Int).SetBytes(urlHash).Uint64()
	// str := u.Base58Encode(gconv.Bytes(number))
    str := u.Base58Encode(urlHash)
	g.Log(logger).Debug(ctx, "utilHelper GenerateShortLink str:", str, " number:", number)
	return str[:8], nil
}
```

## License

`URL-Shortenter` is licensed under the [MIT License](LICENSE), 100% free and open-source, forever.

## Thanks

<a href="https://www.jetbrains.com/?from=URL-Shortenter"><img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jetbrains-training-partner.png" height="120" alt="JetBrains"/></a>
```