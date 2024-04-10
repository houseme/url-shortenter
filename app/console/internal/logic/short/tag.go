/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package short

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/utility/helper"
)

// CreateTag is the handler for CreateTag.
func (s *sShort) CreateTag(ctx context.Context, in *model.CreateTagInput) (out *model.CreateTagOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-CreateTag")
	defer span.End()

	var logger = g.Log(helper.Helper().Logger(ctx))
	logger.Debug(ctx, "short-CreateTag in:", in)
	return
}

// DelTag is the handler for DelTag.
func (s *sShort) DelTag(ctx context.Context, in *model.DelTagInput) (out *model.DelTagOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-DelTag")
	defer span.End()

	var logger = g.Log(helper.Helper().Logger(ctx))
	logger.Debug(ctx, "short-DelTag in:", in)
	return
}

// AddTag is the handler for AddTag.
func (s *sShort) AddTag(ctx context.Context, in *model.AddTagInput) (out *model.AddTagOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-AddTag")
	defer span.End()

	var logger = g.Log(helper.Helper().Logger(ctx))
	logger.Debug(ctx, "short-AddTag in:", in)
	return
}
