/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package account

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/url-shortenter/app/schedule/internal/model"
	"github.com/houseme/url-shortenter/app/schedule/internal/service"
	"github.com/houseme/url-shortenter/utility/helper"
)

type sAccount struct {
}

func init() {
	service.RegisterAccount(&sAccount{})
}

// Stat account stat.
func (s *sAccount) Stat(ctx context.Context, in *model.AccountStatInput) (out *model.AccountStatOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-account-stat")
	defer span.End()

	var (
		log = g.Log(helper.Helper().Logger(ctx))
	)
	log.Debug(ctx, "account-Stat in:", in)
	return
}
