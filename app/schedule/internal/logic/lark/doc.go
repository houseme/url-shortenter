// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

// Package lark is a lark service.
package lark

import (
	"github.com/houseme/url-shortenter/app/schedule/internal/service"
)

type sLark struct {
}

func init() {
	service.RegisterLark(&sLark{})
}
