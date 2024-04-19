// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package short

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gtime"
	"golang.org/x/sync/singleflight"

	"github.com/houseme/url-shortenter/app/schedule/internal/consts"
	"github.com/houseme/url-shortenter/internal/database/dao"
	"github.com/houseme/url-shortenter/internal/database/model/do"
	"github.com/houseme/url-shortenter/internal/database/model/entity"
	"github.com/houseme/url-shortenter/utility/cache"
	"github.com/houseme/url-shortenter/utility/helper"
)

var sfg singleflight.Group

// AccessLog is the struct of access log
func (s *sShort) AccessLog(ctx context.Context) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-sShort-AccessLog")
	defer span.End()

	var (
		logger = g.Log(helper.Helper().Logger(ctx))
		val    *gvar.Var
	)
	if val, err = g.Redis(cache.RedisCache().ShortCacheConn(ctx)).Do(ctx, "LLEN", cache.RedisCache().ShortAccessLogQueue(ctx)); err != nil {
		err = gerror.Wrap(err, "access logger get queue length failed")
		return err
	}

	if val.IsNil() || val.IsEmpty() {
		err = gerror.New("access logger queue is empty")
		return err
	}
	allen := val.Int()
	logger.Info(ctx, "access logger queue length is ", allen)
	if allen <= 0 {
		err = gerror.New("access logger queue is empty length is zero")
		return err
	}
	for i := 0; i < allen; i++ {
		if err = s.dealAccessLog(ctx); err != nil {
			err = gerror.Wrap(err, "access logger deal failed")
			return err
		}
	}
	logger.Info(ctx, "access logger queue end")
	return nil
}

// dealAccessLog is the function of deal access log
func (s *sShort) dealAccessLog(ctx context.Context) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-sShort-dealAccessLog")
	defer span.End()

	var (
		logger = g.Log(helper.Helper().Logger(ctx))
		val    *gvar.Var
	)
	logger.Debug(ctx, "access logger right pop queue start")
	if val, err = g.Redis(cache.RedisCache().ShortCacheConn(ctx)).Do(ctx, "RPOP", cache.RedisCache().ShortAccessLogQueue(ctx)); err != nil {
		err = gerror.Wrap(err, "access logger right pop queue failed")
		return
	}

	if val.IsNil() || val.IsEmpty() {
		logger.Info(ctx, "access logger queue is empty")
		return
	}

	var accessLog *entity.AccessLogs
	if err = val.Scan(&accessLog); err != nil {
		err = gerror.Wrap(err, "access logger scan error")
		return
	}
	if accessLog == nil {
		logger.Info(ctx, "access logger is nil")
		return
	}
	var (
		stu    *entity.ShortUrls
		lastID int64
	)
	if stu, err = s.GetShortCache(ctx, accessLog.ShortUrl); err != nil {
		err = gerror.Wrap(err, "access logger get short url error")
		return
	}
	accessLog.ShortNo = stu.ShortNo
	accessLog.AccountNo = stu.AccountNo
	if lastID, err = dao.AccessLogs.Ctx(ctx).OmitEmpty().Unscoped().InsertAndGetId(accessLog); err != nil {
		err = gerror.Wrap(err, "access logger insert error")
		return
	}
	logger.Info(ctx, "access logger insert success last id is ", lastID)
	if val, err = g.Redis(cache.RedisCache().ShortCacheConn(ctx)).Do(ctx, "LPUSH", cache.RedisCache().ShortAccessLogSummaryQueue(ctx), lastID); err != nil {
		err = gerror.Wrap(err, "access logger left push ShortAccessLogSummaryQueue failed")
		return
	}
	logger.Info(ctx, "access logger left push ShortAccessLogSummaryQueue value is ", val)
	return
}

// ShortAccessLogSummary is the struct of short access log summary
func (s *sShort) ShortAccessLogSummary(ctx context.Context) error {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-sShort-ShortAccessLogSummary")
	defer span.End()

	var (
		logger   = g.Log(helper.Helper().Logger(ctx))
		val, err = g.Redis(cache.RedisCache().ShortCacheConn(ctx)).Do(ctx, "LLEN", cache.RedisCache().ShortAccessLogSummaryQueue(ctx))
	)
	if err != nil {
		return gerror.Wrap(err, "access logger summary get queue length error")
	}

	if val.IsNil() || val.IsEmpty() {
		logger.Debug(ctx, "access logger summary queue is empty")
		return nil
	}
	listLen := val.Int()
	logger.Debug(ctx, "access logger summary queue length is ", listLen)
	if listLen <= 0 {
		logger.Debug(ctx, "access logger Summary queue is empty length is zero")
		return nil
	}

	for i := 0; i < listLen; i++ {
		if err = s.dealLogSummary(ctx); err != nil {
			logger.Errorf(ctx, "access logger summary deal error:%+v", err)
		}
	}
	logger.Info(ctx, "access logger queue end")
	return nil
}

func (s *sShort) dealLogSummary(ctx context.Context) error {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-sShort-dealLogSummary")
	defer span.End()

	var (
		logger   = g.Log(helper.Helper().Logger(ctx))
		val, err = g.Redis(cache.RedisCache().ShortCacheConn(ctx)).Do(ctx, "RPOP", cache.RedisCache().ShortAccessLogSummaryQueue(ctx))
	)

	if err != nil {
		return gerror.Wrap(err, "access logger summary get queue length failed")
	}

	if val.IsNil() || val.IsEmpty() {
		logger.Debug(ctx, "access logger summary queue is empty")
		return nil
	}

	var (
		ID  = val.Uint64()
		alg *entity.AccessLogs
	)

	defer func() {
		if err != nil {
			if _, errs := g.Redis(cache.RedisCache().ShortCacheConn(ctx)).Do(ctx, "LPUSH", cache.RedisCache().ShortAccessLogSummaryQueue(ctx), ID); errs != nil {
				logger.Error(ctx, "access logger summary left push error", errs)
			}
		}
	}()

	if err = dao.AccessLogs.Ctx(ctx).Scan(&alg, do.AccessLogs{Id: ID}); err != nil {
		return gerror.Wrap(err, "access logger summary select db  failed")
	}

	if alg == nil {
		logger.Debug(ctx, "access logger is nil")
		return nil
	}

	var (
		slts   *entity.AccessLogsSummary
		lastID int64
	)
	if err = dao.AccessLogsSummary.Ctx(ctx).Scan(&slts, do.AccessLogsSummary{
		AccountNo:  alg.AccountNo,
		ShortNo:    alg.ShortNo,
		Ip:         alg.Ip,
		YearTime:   alg.AccessTime.Year(),
		MonthTime:  alg.AccessTime.Month(),
		DayTime:    alg.AccessTime.Day(),
		AccessDate: alg.AccessDate,
	}); err != nil {
		return gerror.Wrap(err, "access logger select summary from db failed")
	}
	if slts == nil {
		slts = &entity.AccessLogsSummary{
			ShortNo:    alg.ShortNo,
			AccountNo:  alg.AccountNo,
			ShortUrl:   alg.ShortUrl,
			ShortAll:   alg.ShortAll,
			YearTime:   alg.YearTime,
			MonthTime:  alg.MonthTime,
			DayTime:    alg.DayTime,
			AccessDate: alg.AccessDate,
			UserAgent:  alg.UserAgent,
			Ip:         alg.Ip,
			Summary:    1,
			SuccessSum: 0,
			FailSum:    0,
		}
		if alg.VisitState == consts.VisitStateNormal {
			slts.SuccessSum = 1
		}
		if alg.VisitState == consts.VisitStateInvalid {
			slts.FailSum = 1
		}

		if lastID, err = dao.AccessLogsSummary.Ctx(ctx).OmitEmpty().Unscoped().InsertAndGetId(slts); err != nil {
			return gerror.Wrap(err, "access logger summary insert failed")
		}
		logger.Info(ctx, "access logger summary insert success last id is ", lastID)
	} else {
		if alg.VisitState == consts.VisitStateNormal {
			slts.SuccessSum = slts.SuccessSum + 1
		}
		if alg.VisitState == consts.VisitStateInvalid {
			slts.FailSum = slts.FailSum + 1
		}

		if lastID, err = dao.AccessLogsSummary.Ctx(ctx).OmitEmpty().Unscoped().UpdateAndGetAffected(do.AccessLogsSummary{
			Summary:    slts.Summary + 1,
			SuccessSum: slts.SuccessSum,
			FailSum:    slts.FailSum,
			ModifyTime: gtime.Now(),
		}, do.AccessLogsSummary{
			Id: slts.Id,
		}); err != nil {
			return gerror.Wrap(err, "access logger summary update error")
		}
		logger.Info(ctx, "access logger summary update success lastID:", lastID)
	}

	logger.Info(ctx, "access logger summary modify end success")

	return nil
}

// GetShortCache is the function to get short cache
func (s *sShort) GetShortCache(ctx context.Context, short string) (*entity.ShortUrls, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-GetShortCache")
	defer span.End()

	var (
		log      = g.Log(helper.Helper().Logger(ctx))
		su       *entity.ShortUrls
		val, err = g.Redis(cache.RedisCache().ShortCacheConn(ctx)).Do(ctx, "GET", cache.RedisCache().ShortCacheObject(ctx, short))
	)

	if err == nil && !val.IsNil() && !val.IsEmpty() {
		if err = val.Scan(&su); err == nil && su != nil {
			log.Debug(ctx, "get short cache success", su)
			return su, nil
		}
		if err != nil {
			log.Errorf(ctx, "get short cache scan error:%+v", err)
		} else {
			log.Info(ctx, "get short cache scan after su is nil")
		}
	} else {
		log.Errorf(ctx, "get short cache get from redis error:%+v", err)
	}

	// query DB
	if err = dao.ShortUrls.Ctx(ctx).Scan(&su, do.ShortUrls{ShortUrl: short}); err != nil {
		log.Errorf(ctx, "get short cache query db error:%+v", err)
		return nil, err
	}
	if su == nil {
		log.Info(ctx, "get short cache query db su is nil")
		return nil, gerror.New("short url not found")
	}

	// set cache
	if val, err = g.Redis(cache.RedisCache().ShortCacheConn(ctx)).Do(ctx, "SETEX", cache.RedisCache().ShortCacheObject(ctx, short), 1800, su); err != nil {
		log.Error(ctx, "get short cache set cache error", err)
	}
	log.Info(ctx, "get short cache set cache success", su)
	return su, nil
}
