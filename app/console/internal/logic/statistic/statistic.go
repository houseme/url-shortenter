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

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/url-shortenter/app/console/internal/consts"
	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
	"github.com/houseme/url-shortenter/internal/database/dao"
	"github.com/houseme/url-shortenter/internal/database/model/do"
	"github.com/houseme/url-shortenter/internal/database/model/entity"
	"github.com/houseme/url-shortenter/utility/helper"
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

	logger := g.Log(helper.Helper().Logger(ctx))
	logger.Debug(ctx, "statistic-List in:", in)

	params := do.AccessLogsSummary{}
	if in.AuthAccountLevel == consts.AccountLevelBusiness {
		params.AccountNo = in.AuthAccountNo
	}
	if in.AuthAccountLevel == consts.AccountLevelPlatform {
		params.UserNo = in.AuthUserNo
	}
	if in.SummaryNo != 0 {
		params.SummaryNo = in.SummaryNo
	}

	if in.ShortNo != 0 {
		params.ShortNo = in.ShortNo
	}

	if in.AccessDate != nil {
		params.YearTime = in.AccessDate.Year()
		params.MonthTime = in.AccessDate.Month()
		params.DayTime = in.AccessDate.Day()
	}
	m := dao.AccessLogsSummary.Ctx(ctx).Where(params)
	if in.ShortURL != "" {
		m = m.WhereLike(dao.AccessLogsSummary.Columns().ShortUrl, "%"+in.ShortURL+"&")
	}
	out = &model.StatisticListOutput{
		List: []*model.StatisticItem{},
		Paginate: &model.PageInfoOutput{
			Page:     in.Page,
			PageSize: in.PageSize,
			Total:    0,
		},
	}
	if out.Paginate.Total, err = m.Count(); err != nil {
		err = gerror.Wrap(err, "获取统计总数失败")
		return
	}
	list := ([]*entity.AccessLogsSummary)(nil)
	if err = m.Page(in.Page, in.PageSize).Scan(&list); err != nil {
		err = gerror.Wrap(err, "获取统计列表失败")
		return
	}

	return
}

// Detail is used to get a statistic detail.
func (s *sStatistic) Detail(ctx context.Context, in *model.StatisticDetailInput) (out *model.StatisticDetailOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-statistic-Detail")
	defer span.End()

	logger := g.Log(helper.Helper().Logger(ctx))
	logger.Debug(ctx, "statistic-Detail in:", in)
	var ()
	params := do.AccessLogsSummary{}
	if in.AuthAccountLevel == consts.AccountLevelBusiness {
		params.AccountNo = in.AuthAccountNo
	}
	if in.AuthAccountLevel == consts.AccountLevelPlatform {
		params.UserNo = in.AuthUserNo
	}
	if in.SummaryNo != 0 {
		params.SummaryNo = in.SummaryNo
	}

	if in.ShortNo != 0 {
		params.ShortNo = in.ShortNo
	}

	if in.AccessDate != nil {
		params.YearTime = in.AccessDate.Year()
		params.MonthTime = in.AccessDate.Month()
		params.DayTime = in.AccessDate.Day()
	}
	m := dao.AccessLogsSummary.Ctx(ctx).Where(params)
	if in.ShortURL != "" {
		m = m.WhereLike(dao.AccessLogsSummary.Columns().ShortUrl, "%"+in.ShortURL+"&")
	}
	detail := (*entity.AccessLogsSummary)(nil)
	if err = m.OrderDesc(dao.AccessLogsSummary.Columns().Id).Scan(&detail, do.AccessLogsSummary{
		ShortNo: in.ShortNo,
	}); err != nil {
		err = gerror.Wrap(err, "获取统计详情失败")
		return
	}
	if detail == nil {
		err = gerror.New("统计详情不存在")
		return
	}

	return
}
