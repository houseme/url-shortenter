// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package short

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

// ShortDomain is the handler for ShortDomain
func (s *sShort) ShortDomain(ctx context.Context, in *model.ShortDomainInput) (out *model.ShortDomainOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-ShortDomain")
	defer span.End()

	var log = g.Log(helper.Helper().Logger(ctx))
	log.Debug(ctx, "short-ShortDomain in:", in)
	out = &model.ShortDomainOutput{
		List: []*model.ShortDomainItem{},
	}

	defer func() {
		if err != nil {
			log.Error(ctx, "short-ShortDomain error:", err)
			return
		}
	}()

	if err = dao.ShortDomain.Ctx(ctx).Scan(&out.List, do.ShortDomain{State: consts.ShortDomainStateNormal}); err != nil {
		err = gerror.Wrap(err, "short-ShortDomain error")
		return
	}

	if len(out.List) == 0 {
		err = gerror.New("short-ShortDomain error")
		return
	}
	log.Debug(ctx, "short-ShortDomain out:", out)
	return
}
