// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

// Package home
package home

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gipv4"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/houseme/url-shortenter/app/front/internal/consts"
	"github.com/houseme/url-shortenter/app/front/internal/model"
	"github.com/houseme/url-shortenter/app/front/internal/service"
	"github.com/houseme/url-shortenter/internal/database/dao"
	"github.com/houseme/url-shortenter/internal/database/model/do"
	"github.com/houseme/url-shortenter/internal/database/model/entity"
	"github.com/houseme/url-shortenter/internal/protocol"
	"github.com/houseme/url-shortenter/utility/cache"
	"github.com/houseme/url-shortenter/utility/helper"
)

type sHome struct{}

var sfg singleflight.Group

func init() {
	service.RegisterHome(&sHome{})
}

// ShortDetail short url detail
func (s *sHome) ShortDetail(ctx context.Context, in *model.HomeInput) (out string, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-lite-service-home-ShortDetail")
	defer span.End()

	var (
		isSendLog bool
		logger    = g.Log(helper.Helper().Logger(ctx))
	)
	logger.Debug(ctx, "home-short-detail in:", in)
	in.VisitState = consts.VisitState
	defer func() {
		if isSendLog {
			s.NewAccessLog(ctx, in)
		}
	}()

	var value *gvar.Var
	if value, err = g.Redis(cache.RedisCache().ShortRequestConn(ctx)).Do(ctx, "GET", in.Short); err != nil {
		logger.Errorf(ctx, "home-short-detail get redis value failed err:%+v", err)
	}
	if !value.IsNil() && !value.IsEmpty() && value.String() != "" {
		out = value.String()
		logger.Debug(ctx, "home-short-detail from redis out:", out)
		isSendLog = true
		in.VisitState = consts.VisitStateNormal
		return
	}

	// 开始查询数据库 防止缓存击穿
	v, err, _ := sfg.Do(in.Short, func() (interface{}, error) {
		// query DB
		ent := (*entity.ShortUrls)(nil)
		if err = dao.ShortUrls.Ctx(ctx).Scan(&ent, do.ShortUrls{ShortUrl: in.Short}); err != nil {
			return nil, err
		}

		if ent == nil {
			logger.Debug(ctx, "home-short-detail short url not found")
			return nil, nil
		}
		logger.Debug(ctx, "home-short-detail select from db:", ent)
		if ent.IsValid != consts.ShortValid {
			logger.Debug(ctx, "home-short-detail short url ent.IsValid != consts.ShortValid")
			in.VisitState = consts.VisitStateInvalid
			isSendLog = true
			return nil, nil
		}

		// set cache
		if value, err = g.Redis(cache.RedisCache().ShortRequestConn(ctx)).Do(ctx, "SETEX", in.Short, 86400*2+grand.Intn(2022), ent.DestUrl); err != nil {
			logger.Errorf(ctx, "home-short-detail storage.Redis Set failed err:%+v", err)
		}

		logger.Info(ctx, "home-short-detail set redis cache end shortUrl:", in.Short, "destUrl:", ent.DestUrl)
		return ent.DestUrl, nil
	})

	if err != nil {
		err = gerror.Wrap(err, "query from db failed")
		return
	}

	if v == nil {
		logger.Debug(ctx, "home-short-detail query from db result is nil")
		return "", nil
	}
	isSendLog = true
	in.VisitState = consts.VisitStateNormal
	logger.Debug(ctx, "home-short-detail query from db result:", v)
	out = v.(string)
	return
}

// NewAccessLog 创建访问日志
func (s *sHome) NewAccessLog(ctx context.Context, in *model.HomeInput) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-home-NewAccessLog")
	defer span.End()

	var (
		now        = gtime.Now()
		accessLogs = entity.AccessLogs{
			ShortUrl:   in.Short,
			AccessTime: now,
			AccessDate: now,
			YearTime:   uint(now.Year()),
			MonthTime:  uint(now.Month()),
			DayTime:    uint(now.Day()),
			Ip:         in.ClientIP,
			UserAgent:  in.UserAgent,
			ShortAll:   in.ShortAll,
			TraceId:    span.SpanContext().TraceID().String(),
			VisitState: in.VisitState,
			ServerIp:   helper.Helper().GetOutBoundIP(ctx),
		}
		serverIP, err = gipv4.GetIntranetIp()
		logger        = g.Log(helper.Helper().Logger(ctx))
		value         *gvar.Var
	)

	if err == nil {
		accessLogs.ServerIp = serverIP
	} else {
		logger.Errorf(ctx, "NewAccessLog get server ip failed err:%+v", err)
	}
	logger.Debug(ctx, "NewAccessLog AccessLogs:", accessLogs)
	if value, err = g.Redis(cache.RedisCache().ShortCacheConn(ctx)).Do(ctx, "LPUSH", cache.RedisCache().ShortAccessLogQueue(ctx), accessLogs); err != nil {
		logger.Errorf(ctx, "NewAccessLog redis left push failed err: %+v", err)
		return
	}
	logger.Debug(ctx, "NewAccessLog set redis success value:", value)
}

// ShortAll 短链列表
func (s *sHome) ShortAll(ctx context.Context, in *model.HomeInput) (out []entity.ShortUrls, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-service-home-ShortAll")
	defer span.End()

	var (
		logger          = g.Log(helper.Helper().Logger(ctx))
		now             = gtime.Now()
		intranetIPArray []string
		serverIP        = "NoHostIpFound"
		isSendLog       bool
	)

	if intranetIPArray, err = gipv4.GetIntranetIpArray(); err != nil {
		return
	}

	if len(intranetIPArray) == 0 {
		if intranetIPArray, err = gipv4.GetIpArray(); err != nil {
			return
		}
	}
	if len(intranetIPArray) > 0 {
		serverIP = intranetIPArray[0]
	}

	logger.Debug(ctx, "home-short-all in:", in)
	in.VisitState = consts.VisitState
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
		AccessTime: timestamppb.New(time.UnixMicro(now.UnixMicro())),
		AccessDate: timestamppb.New(time.UnixMicro(now.UnixMicro())),
		YearTime:   &wrapperspb.UInt64Value{Value: uint64(now.Year())},
		MonthTime:  &wrapperspb.UInt64Value{Value: uint64(now.Month())},
		DayTime:    &wrapperspb.UInt64Value{Value: uint64(now.Day())},
	}
	logger.Debug(ctx, "home-short-all acl:", acl)
	return
}
