// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

// Package short logic
package short

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/houseme/url-shortenter/app/console/internal/consts"
	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
	"github.com/houseme/url-shortenter/internal/database/dao"
	"github.com/houseme/url-shortenter/internal/database/model/do"
	"github.com/houseme/url-shortenter/internal/database/model/entity"
	"github.com/houseme/url-shortenter/utility/helper"
	"github.com/houseme/url-shortenter/utility/tld"
)

type sShort struct{}

// init
func init() {
	service.RegisterShort(&sShort{})
}

// CreateShort is the handler for CreateShort
func (s *sShort) CreateShort(ctx context.Context, in *model.CreateShortInput) (out *model.CreateShortOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-CreateShort")
	defer span.End()

	var (
		logger   = g.Log(helper.Helper().Logger(ctx))
		base     = (*entity.ShortUrls)(nil)
		shortURL string
	)
	logger.Debug(ctx, "short-CreateShort in:", in)

	// Generate a short URL using the destination URL and account number
	if shortURL, err = helper.Helper().GenerateShortLink(ctx, in.DestURL+gconv.String(in.AuthAccountNo)); err != nil {
		err = gerror.Wrap(err, "short-CreateShort error")
		return
	}

	// Generate a fixed-length hash for the destination URL and account number
	hash := helper.Helper().GenerateFixedLengthHash(in.DestURL + gconv.String(in.AuthAccountNo))
	// 输出固定长度的哈希值
	logger.Debug(ctx, "short-CreateShort hash:", hash)

	if err = dao.ShortUrls.Ctx(ctx).Scan(&base, do.ShortUrls{
		UserNo:   in.AuthUserNo,
		ShortUrl: shortURL,
		DestHash: hash,
	}); err != nil {
		return
	}

	// If the short url already exists, return the existing short url
	if base != nil {
		out = &model.CreateShortOutput{
			ShortURL: base.ShortUrl,
			ShortNo:  base.ShortNo,
		}
		return
	}

	var (
		shortNo = helper.Helper().InitTrxID(ctx, in.AuthUserNo)
		now     = gtime.Now()
		tldResp *tld.DomainTLDResp
		lastID  int64
	)
	if tldResp, err = tld.GetTLD(ctx, in.DestURL, 0); err != nil {
		err = gerror.Wrap(err, "short-CreateShort error")
		return
	}
	if lastID, err = dao.ShortUrls.Ctx(ctx).Data(do.ShortUrls{
		UserNo:        in.AuthUserNo,
		DestUrl:       in.DestURL,
		ShortDomain:   consts.DefaultShortDomain,
		ShortDomainNo: 0,
		ShortNo:       shortNo,
		ShortUrl:      shortURL,
		DestHash:      hash,
		Domain:        tldResp.Domain,
		IsValid:       consts.ShortUrlStateNormal,
		Memo:          in.Memo,
		RawState:      consts.ShortUrlRawStateDefault,
		Sort:          now.TimestampMilli(),
		CollectState:  consts.CollectStateDefault,
		CreateTime:    now,
		ModifyTime:    now,
	}).InsertAndGetId(); err != nil {
		err = gerror.Wrap(err, "short-CreateShort error")
		return
	}
	logger.Debug(ctx, "short-CreateShort lastID:", lastID)
	if lastID <= 0 {
		err = gerror.New("short-CreateShort error")
		return
	}
	out = &model.CreateShortOutput{
		ShortURL: shortURL,
		ShortNo:  shortNo,
	}
	logger.Debug(ctx, "short-CreateShort out:", out)
	return
}

// ModifyShort is the handler for ModifyShort
func (s *sShort) ModifyShort(ctx context.Context, in *model.ModifyShortInput) (out *model.ModifyShortOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-ModifyShort")
	defer span.End()

	logger := g.Log(helper.Helper().Logger(ctx))
	logger.Debug(ctx, "short-ModifyShort in:", in)

	return
}

// QueryShort is the handler for QueryShort
func (s *sShort) QueryShort(ctx context.Context, in *model.QueryShortInput) (out *model.QueryShortOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-QueryShort")
	defer span.End()

	logger := g.Log(helper.Helper().Logger(ctx))
	logger.Debug(ctx, "short-QueryShort in:", in)
	short := (*entity.ShortUrls)(nil)
	if err = dao.ShortUrls.Ctx(ctx).Scan(&short, do.ShortUrls{
		ShortNo: in.ShortNo,
	}); err != nil {
		return
	}

	if short == nil {
		err = gerror.New("short-QueryShort error")
		return
	}

	return
}

// QueryShortList is the handler for QueryShortList
func (s *sShort) QueryShortList(ctx context.Context, in *model.QueryShortListInput) (out *model.QueryShortListOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-QueryShortList")
	defer span.End()

	logger := g.Log(helper.Helper().Logger(ctx))
	logger.Debug(ctx, "short-QueryShortList in:", in)

	return
}

// QueryStat is the handler for QueryStat
func (s *sShort) QueryStat(ctx context.Context, in *model.QueryStatInput) (out *model.QueryStatOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-QueryStat")
	defer span.End()

	logger := g.Log(helper.Helper().Logger(ctx))
	logger.Debug(ctx, "short-QueryStat in:", in)

	return
}

// ShortDomain is the handler for ShortDomain
func (s *sShort) ShortDomain(ctx context.Context, in *model.ShortDomainInput) (out *model.ShortDomainOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-ShortDomain")
	defer span.End()

	logger := g.Log(helper.Helper().Logger(ctx))
	logger.Debug(ctx, "short-ShortDomain in:", in)
	out = &model.ShortDomainOutput{
		List: []*model.ShortDomainItem{},
	}

	defer func() {
		if err != nil {
			logger.Error(ctx, "short-ShortDomain error:", err)
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
	logger.Debug(ctx, "short-ShortDomain out:", out)
	return
}
