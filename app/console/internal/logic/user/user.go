// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
	"github.com/houseme/url-shortenter/utility/helper"
)

type sUser struct {
}

func init() {
	service.RegisterUser(initUser())
}

// initUser
func initUser() *sUser {
	return &sUser{}
}

// CreateMerchant creates a new merchant.
func (s *sUser) CreateMerchant(ctx context.Context, in *model.CreateMerchantInput) (out *model.CreateMerchantOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-user-CreateMerchant")
	defer span.End()

	var log = g.Log(helper.Helper().Logger(ctx))
	log.Debug(ctx, "user-CreateMerchant in:", in)

	return
}
