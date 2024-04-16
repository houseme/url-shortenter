/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

// Package statistic logic
package statistic

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

type sStatistic struct{}

// init is the initialization of sStatistic.
func init() {
	service.RegisterStatistic(&sStatistic{})
}

// List is used to list statistic.
func (s *sStatistic) List(ctx context.Context, in *model.StatisticListInput) (out *model.StatisticListOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-statistic-List")
	defer span.End()

	return
}

// Detail is used to get a statistic detail.
func (s *sStatistic) Detail(ctx context.Context, in *model.StatisticDetailInput) (out *model.StatisticDetailOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-statistic-Detail")
	defer span.End()

	return
}
