// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package home

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

type sHome struct {
}

func init() {
	service.RegisterHome(initHome())
}

func initHome() *sHome {
	return &sHome{}
}

// Index home page.
func (s *sHome) Index(ctx context.Context, in *model.HomeIndexInput) (out *model.HomeIndexOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-home-Index")
	defer span.End()

	return
}
