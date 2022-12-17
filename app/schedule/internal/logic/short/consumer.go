package short

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
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
func (s *sShort) AccessLog(ctx context.Context) error {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-AccessLog")
	defer span.End()

	var (
		logger    = helper.Helper().Logger(ctx)
		conn, err = g.Redis(cache.RedisCache().ShortCacheConn(ctx)).Conn(ctx)
	)
	if err != nil {
		err = gerror.Wrap(err, "access log create redis connection error")
		return err
	}

	defer func() {
		_ = conn.Close(ctx)
		if err != nil {
			g.Log(logger).Error(ctx, "access log error is ", err)
		}
	}()

	var val *gvar.Var
	if val, err = conn.Do(ctx, "LLEN", cache.RedisCache().ShortAccessLogQueue(ctx)); err != nil {
		err = gerror.Wrap(err, "access log get queue length error")
		return err
	}

	if val.IsNil() || val.IsEmpty() {
		err = gerror.New("access log queue is empty")
		return err
	}
	allen := val.Int()
	g.Log(logger).Info(ctx, "access log queue length is ", allen)
	if allen <= 0 {
		err = gerror.New("access log queue is empty length is zero")
		return err
	}
	for i := 0; i < allen; i++ {
		if err = s.dealAccessLog(ctx); err != nil {
			err = gerror.Wrap(err, "access log deal error")
			return err
		}
	}
	g.Log(logger).Info(ctx, "access log queue end")
	return nil
}

// dealAccessLog is the function of deal access log
func (s *sShort) dealAccessLog(ctx context.Context) error {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-dealAccessLog")
	defer span.End()

	var (
		logger    = helper.Helper().Logger(ctx)
		val       *gvar.Var
		conn, err = g.Redis(cache.RedisCache().ShortCacheConn(ctx)).Conn(ctx)
	)

	if err != nil {
		err = gerror.New("access log create redis connection error")
		return err
	}

	if val, err = conn.Do(ctx, "RPOP", cache.RedisCache().ShortAccessLogQueue(ctx)); err != nil {
		err = gerror.Wrap(err, "access log get queue error")
		return err
	}

	if val.IsNil() || val.IsEmpty() {
		g.Log(logger).Info(ctx, "access log queue is empty")
		err = gerror.New("access log queue is empty")
		return err
	}

	var accessLog *entity.AccessLogs
	if err = val.Scan(&accessLog); err != nil {
		err = gerror.Wrap(err, "access log scan error")
		return err
	}
	if accessLog == nil {
		g.Log(logger).Info(ctx, "access log is nil")
		err = gerror.New("access log is nil")
		return err
	}
	var (
		stu    *entity.ShortUrls
		lastID int64
	)
	if stu, err = s.GetShortCache(ctx, accessLog.ShortUrl); err != nil {
		err = gerror.Wrap(err, "access log get short url error")
		return err
	}
	accessLog.ShortNo = stu.ShortNo
	accessLog.AccountNo = stu.AccountNo
	if lastID, err = dao.AccessLogs.Ctx(ctx).OmitEmpty().Unscoped().InsertAndGetId(accessLog); err != nil {
		err = gerror.Wrap(err, "access log insert error")
		return err
	}
	g.Log(logger).Info(ctx, "access log insert success last id is ", lastID)
	if val, err = conn.Do(ctx, "LPUSH", cache.RedisCache().ShortAccessLogSummaryQueue(ctx), lastID); err != nil {
		err = gerror.Wrap(err, "access log left push ShortAccessLogSummaryQueue error")
		return err
	}
	g.Log(logger).Info(ctx, "access log left push ShortAccessLogSummaryQueue value is ", val)
	return nil
}

// ShortAccessLogSummary is the struct of short access log summary
func (s *sShort) ShortAccessLogSummary(ctx context.Context) error {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-ShortAccessLogSummary")
	defer span.End()

	var (
		logger    = helper.Helper().Logger(ctx)
		conn, err = g.Redis(cache.RedisCache().ShortCacheConn(ctx)).Conn(ctx)
	)
	if err != nil {
		g.Log(logger).Error(ctx, "access log Summary create redis connection error", err)
		return err
	}

	defer conn.Close(ctx)

	var val *gvar.Var
	if val, err = conn.Do(ctx, "LLEN", cache.RedisCache().ShortAccessLogSummaryQueue(ctx)); err != nil {
		g.Log(logger).Error(ctx, "access log summary get queue length error", err)
		return err
	}

	if val.IsNil() || val.IsEmpty() {
		g.Log(logger).Info(ctx, "access log summary queue is empty")
		return gerror.New("access log summary queue is empty")
	}
	llen := val.Int()
	g.Log(logger).Info(ctx, "access log summary queue length is ", llen)
	if llen <= 0 {
		g.Log(logger).Error(ctx, "access log Summary queue is empty length is zero")
		return gerror.New("access log Summary queue is empty length is zero")
	}

	for i := 0; i < llen; i++ {
		if err = s.dealLogSummary(ctx); err != nil {
			g.Log(logger).Error(ctx, "access log summary deal error", err)
		}
	}
	g.Log(logger).Info(ctx, "access log queue end")
	return nil
}

func (s *sShort) dealLogSummary(ctx context.Context) error {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-dealLogSummary")
	defer span.End()

	var (
		logger    = helper.Helper().Logger(ctx)
		val       *gvar.Var
		conn, err = g.Redis(cache.RedisCache().ShortCacheConn(ctx)).Conn(ctx)
	)
	if err != nil {
		err = gerror.Wrap(err, "access log Summary create redis connection error")
		return err
	}

	defer conn.Close(ctx)

	if val, err = conn.Do(ctx, "RPOP", cache.RedisCache().ShortAccessLogSummaryQueue(ctx)); err != nil {
		err = gerror.Wrap(err, "access log summary get queue length error")
		return err
	}

	if val.IsNil() || val.IsEmpty() {
		g.Log(logger).Info(ctx, "access log summary queue is empty")
		return gerror.New("access log summary queue is empty")
	}

	var (
		ID  = val.Uint64()
		alg *entity.AccessLogs
	)

	defer func() {
		if err != nil {
			if _, errs := conn.Do(ctx, "LPUSH", cache.RedisCache().ShortAccessLogSummaryQueue(ctx), ID); errs != nil {
				g.Log(logger).Error(ctx, "access log summary left push error", errs)
			}
		}
	}()

	if err = dao.AccessLogs.Ctx(ctx).OmitEmpty().Scan(&alg, "id = ?", ID); err != nil {
		err = gerror.Wrap(err, "access log summary select db  error")
		return err
	}

	if alg == nil {
		g.Log(logger).Info(ctx, "access log is nil")
		err = gerror.New("access log is nil")
		return err
	}

	var slts *entity.AccessLogsSummary
	if err = dao.AccessLogsSummary.Ctx(ctx).Scan(&slts, do.AccessLogsSummary{
		AccountNo:  alg.AccountNo,
		ShortNo:    alg.ShortNo,
		Ip:         alg.Ip,
		YearTime:   alg.AccessTime.Year(),
		MonthTime:  alg.AccessTime.Month(),
		DayTime:    alg.AccessTime.Day(),
		AccessDate: alg.AccessDate,
	}); err != nil {
		err = gerror.Wrap(err, "access log select summary from db error")
		return err
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
		var lastID int64
		if lastID, err = dao.AccessLogsSummary.Ctx(ctx).OmitEmpty().Unscoped().InsertAndGetId(slts); err != nil {
			err = gerror.Wrap(err, "access log summary insert error")
			return err
		}
		g.Log(logger).Info(ctx, "access log summary insert success last id is ", lastID)
	} else {
		if alg.VisitState == consts.VisitStateNormal {
			slts.SuccessSum = slts.SuccessSum + 1
		}
		if alg.VisitState == consts.VisitStateInvalid {
			slts.FailSum = slts.FailSum + 1
		}

		if _, err = dao.AccessLogsSummary.Ctx(ctx).OmitEmpty().Unscoped().Where("id = ?", slts.Id).Update(g.Map{
			dao.AccessLogsSummary.Columns().Summary:    slts.Summary + 1,
			dao.AccessLogsSummary.Columns().SuccessSum: slts.SuccessSum,
			dao.AccessLogsSummary.Columns().FailSum:    slts.FailSum,
			dao.AccessLogsSummary.Columns().ModifyTime: gdb.Raw("current_timestamp(6)"),
		}); err != nil {
			err = gerror.Wrap(err, "access log summary update error")
			return err
		}
		g.Log(logger).Info(ctx, "access log summary update success")
	}
	g.Log(logger).Info(ctx, "access log summary modify end success")

	return nil
}

// GetShortCache is the function to get short cache
func (s *sShort) GetShortCache(ctx context.Context, short string) (*entity.ShortUrls, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-GetShortCache")
	defer span.End()

	var (
		logger    = helper.Helper().Logger(ctx)
		conn, err = g.Redis(cache.RedisCache().ShortCacheConn(ctx)).Conn(ctx)
	)
	if err != nil {
		g.Log(logger).Error(ctx, "get short cache create redis connection error", err)
		return nil, err
	}

	defer conn.Close(ctx)
	var (
		val *gvar.Var
		su  *entity.ShortUrls
	)
	if val, err = conn.Do(ctx, "GET", cache.RedisCache().ShortCacheObject(ctx, short)); err == nil {
		if !val.IsNil() && !val.IsEmpty() {
			if err = val.Scan(&su); err == nil {
				if su != nil {
					g.Log(logger).Info(ctx, "get short cache success", su)
					return su, nil
				}
				g.Log(logger).Info(ctx, "get short cache scan after su is nil")
			} else {
				g.Log(logger).Error(ctx, "get short cache scan error", err)
			}
		} else {
			g.Log(logger).Info(ctx, "get short cache is nil")
		}
	} else {
		g.Log(logger).Error(ctx, "get short cache get from redis error", err)
	}
	// 防止 缓存击穿
	v, err, _ := sfg.Do(short, func() (interface{}, error) {
		// query DB
		if err = dao.ShortUrls.Ctx(ctx).Scan(&su, "short_url=?", short); err != nil {
			g.Log(logger).Error(ctx, "get short cache query db error", err)
			return nil, err
		}
		if su == nil {
			g.Log(logger).Info(ctx, "get short cache query db su is nil")
			return nil, nil
		}

		// set cache
		if val, err = conn.Do(ctx, "SETEX", cache.RedisCache().ShortCacheObject(ctx, short), 1800, su); err != nil {
			g.Log(logger).Error(ctx, "get short cache set cache error", err)
		}
		g.Log(logger).Info(ctx, "get short cache set cache success", su)
		return su, err
	})

	if err != nil {
		return nil, err
	}

	if v == nil {
		g.Log(logger).Info(ctx, "get short cache from db v is nil")
		return nil, gerror.New("get short cache from db v is nil")
	}
	return v.(*entity.ShortUrls), nil
}
