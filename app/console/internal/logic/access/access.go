/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

// Package access logic
package access

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
	"github.com/houseme/url-shortenter/internal/database/dao"
	"github.com/houseme/url-shortenter/internal/database/model/do"
	"github.com/houseme/url-shortenter/internal/database/model/entity"
	"github.com/houseme/url-shortenter/utility/helper"
)

type sAccess struct{}

// init
func init() {
	service.RegisterAccess(&sAccess{})
}

// List is the handler for List.
func (s *sAccess) List(ctx context.Context, in *model.AccessListInput) (out *model.AccessListOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-access-List")
	defer span.End()

	var (
		logger = g.Log(helper.Helper().Logger(ctx))
		list   = ([]*entity.AccessLogs)(nil)
	)
	logger.Debug(ctx, "access-List in:", in)

	if err = dao.AccessLogs.Ctx(ctx).Page(in.Page, in.PageSize).Scan(&list, do.AccessLogs{
		AccountNo: in.AuthAccountNo,
	}); err != nil {
		err = gerror.Wrap(err, "dao.AccessLogs.Scan failed")
		return
	}

	return
}

// Detail is the handler for Detail.
func (s *sAccess) Detail(ctx context.Context, in *model.AccessDetailInput) (out *model.AccessDetailOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-access-Detail")
	defer span.End()

	var (
		logger = g.Log(helper.Helper().Logger(ctx))
		detail = (*entity.AccessLogs)(nil)
	)
	logger.Debug(ctx, "access-Detail in:", in)
	if err = dao.AccessLogs.Ctx(ctx).Scan(&detail, do.AccessLogs{
		AccountNo: in.AuthAccountNo,
		Id:        in.ID,
	}); err != nil {
		err = gerror.Wrap(err, "dao.AccessLogs.Scan failed")
		return
	}

	if detail == nil {
		err = gerror.New("access detail not found")
		return
	}

	return
}
