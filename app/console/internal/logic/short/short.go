// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package short

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

type sShort struct {
}

// init
func init() {
	service.RegisterShort(initShort())
}

func initShort() *sShort {
	return &sShort{}
}

// CreateShort is the handler for CreateShort
func (s *sShort) CreateShort(ctx context.Context, in *model.CreateShortInput) (out *model.CreateShortOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-CreateShort")
	defer span.End()

	return
}

// ModifyShort is the handler for ModifyShort
func (s *sShort) ModifyShort(ctx context.Context, in *model.ModifyShortInput) (out *model.ModifyShortOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-ModifyShort")
	defer span.End()

	return
}

// QueryShort is the handler for QueryShort
func (s *sShort) QueryShort(ctx context.Context, in *model.QueryShortInput) (out *model.QueryShortOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-QueryShort")
	defer span.End()

	return
}

// QueryStat is the handler for QueryStat
func (s *sShort) QueryStat(ctx context.Context, in *model.QueryStatInput) (out *model.QueryStatOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-QueryStat")
	defer span.End()

	return
}
