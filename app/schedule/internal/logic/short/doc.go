// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

// Package short is a short service.
package short

import (
	"github.com/houseme/url-shortenter/app/schedule/internal/service"
)

type sShort struct{}

func init() {
	service.RegisterShort(&sShort{})
}
