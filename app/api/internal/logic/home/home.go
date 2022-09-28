// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

// Package home
package home

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gipv4"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/houseme/url-shortenter/app/api/internal/consts"
	"github.com/houseme/url-shortenter/app/api/internal/model"
	"github.com/houseme/url-shortenter/app/api/internal/service"
	"github.com/houseme/url-shortenter/internal/database/dao"
	"github.com/houseme/url-shortenter/internal/database/model/entity"
	"github.com/houseme/url-shortenter/internal/protocol"
	"github.com/houseme/url-shortenter/utility"
	"github.com/houseme/url-shortenter/utility/cache"
)

type sHome struct {
}

var (
	sfg singleflight.Group
)

func init() {
	service.RegisterHome(initHome())
}

func initHome() *sHome {
	return &sHome{}
}

// ShortDetail short url detail
func (s *sHome) ShortDetail(ctx context.Context, in *model.HomeInput) (out string, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-lite-service-home-ShortDetail")
	defer span.End()

	var (
		logger    = utility.Helper().Logger(ctx)
		conn      *gredis.RedisConn
		isSendLog bool
	)
	g.Log(logger).Debug(ctx, "home-short-detail in:", in)
	in.VisitState = consts.VisitState
	defer func() {
		if isSendLog {
			s.NewAccessLog(ctx, in)
		}
	}()

	if conn, err = g.Redis(cache.RedisCache().ShortRequestConn(ctx)).Conn(ctx); err != nil {
		err = gerror.Wrap(err, "failed to get redis connection")
		return "", err
	}
	defer func() {
		_ = conn.Close(ctx)
	}()

	var val *gvar.Var
	if val, err = conn.Do(ctx, "GET", in.Short); err != nil {
		g.Log(logger).Error(ctx, "home-short-detail get redis value failed err:", err)
	}
	if !val.IsNil() && !val.IsEmpty() && val.String() != "" {
		out = val.String()
		g.Log(logger).Debug(ctx, "home-short-detail from redis out:", out)
		isSendLog = true
		in.VisitState = consts.VisitStateNormal
		return
	}

	// 开始查询数据库 防止缓存击穿
	v, err, _ := sfg.Do(in.Short, func() (interface{}, error) {
		// query DB
		var ent = (*entity.ShortUrls)(nil)
		if err = dao.ShortUrls.Ctx(ctx).Scan(&ent, "short_url = ?", in.Short); err != nil {
			return nil, err
		}

		if ent == nil {
			g.Log(logger).Debug(ctx, "home-short-detail short url not found")
			return nil, nil
		}
		g.Log(logger).Debug(ctx, "home-short-detail select from db:", ent)
		if ent.IsValid != consts.ShortValid {
			g.Log(logger).Debug(ctx, "home-short-detail short url ent.IsValid != consts.ShortValid")
			in.VisitState = consts.VisitStateInvalid
			isSendLog = true
			return nil, nil
		}

		// set cache
		if val, err = conn.Do(ctx, "SETEX", in.Short, 86400*2+grand.Intn(2022), ent.DestUrl); err != nil {
			g.Log(logger).Error(ctx, "home-short-detail storage.Redis Set failed err:", err)
		}

		g.Log(logger).Info(ctx, "home-short-detail set redis cache end shortUrl:", in.Short, "destUrl:", ent.DestUrl)
		return ent.DestUrl, nil
	})

	if err != nil {
		g.Log(logger).Error(ctx, "home-short-detail query db failed err:", err)
		err = gerror.Wrap(err, "failed to query db")
		return "", err
	}

	if v == nil {
		return "", nil
	}
	isSendLog = true
	in.VisitState = consts.VisitStateNormal
	g.Log(logger).Debug(ctx, "home-short-detail from db out:", v)
	out = v.(string)
	return
}

// NewAccessLog 创建访问日志
func (s *sHome) NewAccessLog(ctx context.Context, in *model.HomeInput) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-home-NewAccessLog")
	defer span.End()

	var (
		t             = gtime.Now()
		serverIP, err = gipv4.GetIntranetIp()
		l             = entity.AccessLogs{
			ShortUrl:   in.Short,
			AccessTime: t,
			AccessDate: t,
			YearTime:   uint(t.Year()),
			MonthTime:  uint(t.Month()),
			DayTime:    uint(t.Day()),
			Ip:         in.ClientIP,
			UserAgent:  in.UserAgent,
			ShortAll:   in.ShortAll,
			TraceId:    span.SpanContext().TraceID().String(),
			VisitState: in.VisitState,
			ServerIp:   serverIP,
		}
		val    *gvar.Var
		logger = utility.Helper().Logger(ctx)
	)
	if err != nil {
		g.Log(logger).Error(ctx, "home-new-access-log get intranet ip failed err:", err)
		l.ServerIp = utility.Helper().GetOutBoundIP(ctx)
	}

	if val, err = g.Redis(cache.RedisCache().ShortCacheConn(ctx)).Do(ctx, "LPUSH",
		cache.RedisCache().ShortAccessLogQueue(ctx), l); err != nil {
		g.Log(logger).Error(ctx, "NewAccessLog err:", err)
	}
	g.Log(logger).Debug(ctx, "NewAccessLog set redis :", val, " access log:", l)
}

// ShortAll 短链列表
func (s *sHome) ShortAll(ctx context.Context, in *model.HomeInput) (out []entity.ShortUrls, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-home-ShortAll")
	defer span.End()

	var (
		logger    = utility.Helper().Logger(ctx)
		isSendLog bool
		serverIP  string
		t         = gtime.Now()
	)

	g.Log(logger).Debug(ctx, "home-short-all in:", in)
	in.VisitState = consts.VisitState
	if serverIP, err = gipv4.GetIntranetIp(); err != nil {
		g.Log(logger).Error(ctx, "home-short-all get server ip failed err:", err)
		return
	}
	defer func() {
		if isSendLog {
			s.NewAccessLog(ctx, in)
		}
	}()

	acl := &protocol.AccessLogRequest{
		ShortUrl:   &wrapperspb.StringValue{Value: in.Short},
		ShortAll:   &wrapperspb.StringValue{Value: in.ShortAll},
		VisitState: &wrapperspb.UInt64Value{Value: uint64(in.VisitState)},
		Ip:         &wrapperspb.StringValue{Value: in.ClientIP},
		ServerIP:   &wrapperspb.StringValue{Value: serverIP},
		TraceID:    &wrapperspb.StringValue{Value: span.SpanContext().TraceID().String()},
		UserAgent:  &wrapperspb.StringValue{Value: in.UserAgent},
		AccessTime: &timestamppb.Timestamp{Seconds: t.Unix(), Nanos: int32(t.Nanosecond())},
		AccessDate: &timestamppb.Timestamp{Seconds: t.Unix(), Nanos: int32(t.Nanosecond())},
		YearTime:   &wrapperspb.UInt64Value{Value: uint64(t.Year())},
		MonthTime:  &wrapperspb.UInt64Value{Value: uint64(t.Month())},
		DayTime:    &wrapperspb.UInt64Value{Value: uint64(t.Day())},
	}
	g.Log().Debug(ctx, "home-short-all acl:", acl)
	return
}
