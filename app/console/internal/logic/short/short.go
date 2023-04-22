// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package short

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
	"github.com/houseme/url-shortenter/utility/helper"
)

type sShort struct {
}

// init
func init() {
	service.RegisterShort(&sShort{})
}

// CreateShort is the handler for CreateShort
func (s *sShort) CreateShort(ctx context.Context, in *model.CreateShortInput) (out *model.CreateShortOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-CreateShort")
	defer span.End()

	var log = g.Log(helper.Helper().Logger(ctx))
	log.Debug(ctx, "short-CreateShort in:", in)

	return
}

// ModifyShort is the handler for ModifyShort
func (s *sShort) ModifyShort(ctx context.Context, in *model.ModifyShortInput) (out *model.ModifyShortOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-ModifyShort")
	defer span.End()

	var log = g.Log(helper.Helper().Logger(ctx))
	log.Debug(ctx, "short-ModifyShort in:", in)

	return
}

// QueryShort is the handler for QueryShort
func (s *sShort) QueryShort(ctx context.Context, in *model.QueryShortInput) (out *model.QueryShortOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-QueryShort")
	defer span.End()

	var log = g.Log(helper.Helper().Logger(ctx))
	log.Debug(ctx, "short-QueryShort in:", in)

	return
}

// QueryStat is the handler for QueryStat
func (s *sShort) QueryStat(ctx context.Context, in *model.QueryStatInput) (out *model.QueryStatOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-QueryStat")
	defer span.End()

	var log = g.Log(helper.Helper().Logger(ctx))
	log.Debug(ctx, "short-QueryStat in:", in)

	return
}
