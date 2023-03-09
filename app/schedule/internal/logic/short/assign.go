// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package short

import (
	"context"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"

	"github.com/houseme/url-shortenter/app/schedule/internal/consts"
	"github.com/houseme/url-shortenter/internal/alibaba"
	"github.com/houseme/url-shortenter/internal/database/dao"
	"github.com/houseme/url-shortenter/internal/database/model/do"
	"github.com/houseme/url-shortenter/internal/database/model/entity"
	"github.com/houseme/url-shortenter/internal/tencent"
	"github.com/houseme/url-shortenter/utility/cache"
	"github.com/houseme/url-shortenter/utility/env"
	"github.com/houseme/url-shortenter/utility/helper"
)

// AssignTask is the assign task 镜像抓起分发队列
func (s *sShort) AssignTask(ctx context.Context) error {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-AssignTask")
	defer span.End()

	var (
		logger = helper.Helper().Logger(ctx)
		err    error
	)
	g.Log(logger).Info(ctx, "AssignTask start")

	defer func() {
		if err != nil {
			g.Log(logger).Error(ctx, "AssignTask failed:", err)
		}
		g.Log(logger).Info(ctx, "AssignTask end")
	}()

	var list = ([]*entity.ShortUrls)(nil)
	if err = dao.ShortUrls.Ctx(ctx).Where(do.ShortUrls{IsValid: consts.ShortValid, CollectState: consts.ShortCollectStateProcessing}).Scan(&list); err != nil {
		err = gerror.Wrap(err, "AssignTask dao.ShortUrls.Ctx(ctx).Where failed")
		return err
	}

	if list == nil || len(list) == 0 {
		g.Log(logger).Info(ctx, "AssignTask list is nil or len(list) == 0")
		return nil
	}

	var conn gredis.Conn
	if conn, err = g.Redis(cache.RedisCache().ShortConn(ctx)).Conn(ctx); err != nil {
		g.Log(logger).Error(ctx, "AssignTask g.Redis(cache.RedisCache().ShortConn(ctx)).Conn(ctx) failed:", err)
		return err
	}
	g.Log(logger).Info(ctx, "AssignTask list len", len(list))
	defer func() {
		if errs := conn.Close(ctx); errs != nil {
			g.Log(logger).Error(ctx, "AssignTask conn.Close failed:", errs)
		}
		g.Log(logger).Info(ctx, "AssignTask conn.Close")
	}()
	var result *gvar.Var
	for i := 0; i < len(list); i++ {
		if result, err = conn.Do(ctx, "LPUSH", cache.RedisCache().ShortMirrorQueue(ctx), list[i].ShortNo); err != nil {
			g.Log(logger).Error(ctx, "AssignTask conn.Do LPUSH failed:", err)
		}
		g.Log(logger).Info(ctx, "AssignTask LPUSH result:", result)
	}

	g.Log(logger).Info(ctx, "AssignTask end")
	return nil
}

// AuditAssignTask is the assign task 审核跟踪队列
func (s *sShort) AuditAssignTask(ctx context.Context) error {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-AuditAssignTask")
	defer span.End()

	var (
		logger = helper.Helper().Logger(ctx)
		list   = ([]*entity.ShortUrls)(nil)
		err    error
	)
	g.Log(logger).Info(ctx, "AuditAssignTask start")
	defer func() {
		if err != nil {
			g.Log(logger).Error(ctx, "AuditAssignTask failed:", err)
		}
		g.Log(logger).Info(ctx, "AuditAssignTask end")
	}()

	if err = dao.ShortUrls.Ctx(ctx).Where(do.ShortUrls{IsValid: consts.ShortValid,
		CollectState: consts.ShortCollectStateSuccess}).Scan(&list); err != nil {
		err = gerror.Wrap(err, "AuditAssignTask dao.ShortUrls.Ctx(ctx).Where failed")
		return err
	}

	if list == nil || len(list) == 0 {
		g.Log(logger).Info(ctx, "AuditAssignTask list is nil or len(list) == 0")
		return nil
	}

	var (
		conn, _ = g.Redis(cache.RedisCache().ShortConn(ctx)).Conn(ctx)
		llen    = len(list)
		result  *gvar.Var
	)
	g.Log(logger).Info(ctx, "AuditAssignTask list len", llen)
	defer func() {
		if errs := conn.Close(ctx); errs != nil {
			g.Log(logger).Error(ctx, "AuditAssignTask conn.Close failed:", errs)
		}
	}()
	for i := 0; i < llen; i++ {
		if result, err = conn.Do(ctx, "LPUSH", cache.RedisCache().ShortAuditQueue(ctx), list[i].ShortNo); err != nil {
			g.Log(logger).Error(ctx, "AuditAssignTask conn.Do LPUSH failed:", err)
		}
		g.Log(logger).Info(ctx, "AuditAssignTask LPUSH result:", result)
	}

	g.Log(logger).Info(ctx, "AuditAssignTask end")
	return nil
}

// ExecuteAudit the given command. 执行跟踪审核
func (s *sShort) ExecuteAudit(ctx context.Context) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-ExecuteAudit")
	defer span.End()
	var (
		pool   = grpool.New(10)
		logger = helper.Helper().Logger(ctx)
	)

	g.Log(logger).Info(ctx, "Execute start")
	gtimer.SetInterval(ctx, 5*time.Second, func(ctx context.Context) {
		ctx = helper.Helper().SetLogger(ctx, "schedule")
		g.Log(logger).Info(ctx, "Execute loop")
		for i := 0; i < 5; i++ {
			if err := pool.Add(ctx, s.QueryShortAndGrabAudit); err != nil {
				g.Log(logger).Error(ctx, "Execute pool.Add failed:", err)
			}
		}
		g.Log(logger).Info(ctx, "Execute loop end")
	})
	select {}
}

// QueryShortAndGrabAudit the given command.
func (s *sShort) QueryShortAndGrabAudit(ctx context.Context) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-QueryShortAndGrabAudit")
	defer span.End()
	var (
		logger    = helper.Helper().Logger(ctx)
		conn, err = g.Redis(cache.RedisCache().ShortConn(ctx)).Conn(ctx)
		shortURL  = (*entity.ShortUrls)(nil)
		result    *gvar.Var
	)

	defer func() {
		if errs := conn.Close(ctx); errs != nil {
			g.Log(logger).Error(ctx, "QueryShortAndGrabAudit conn.Close failed:", errs)
		}
		if err != nil {
			g.Log(logger).Error(ctx, "QueryShortAndGrabAudit failed err:", err)
		}
		g.Log(logger).Info(ctx, "QueryShortAndGrabAudit conn.Close")
	}()

	if err != nil {
		g.Log(logger).Error(ctx, "QueryShortAndGrabAudit g.Redis(cache.RedisCache().ShortConn(ctx)).Conn(ctx) failed:", err)
		return
	}

	if result, err = conn.Do(ctx, "RPOP", cache.RedisCache().ShortAuditQueue(ctx)); err != nil {
		err = gerror.Wrap(err, "QueryShortAndGrabAudit conn.Do RPOP failed")
		return
	}
	if result == nil || result.IsNil() || result.IsEmpty() {
		g.Log(logger).Info(ctx, "QueryShortAndGrabAudit result is nil")
		return
	}
	g.Log(logger).Info(ctx, "QueryShortAndGrabAudit shortNo:", result.String())
	if err = dao.ShortUrls.Ctx(ctx).Scan(&shortURL, "short_no = ?", result.Uint64()); err != nil {
		err = gerror.Wrap(err, "QueryShortAndGrabAudit ShortUrls.Scan failed")
		return
	}

	if shortURL == nil {
		g.Log(logger).Info(ctx, "QueryShortAndGrabAudit ShortUrls.Scan no data")
		return
	}

	if shortURL.CollectState != consts.ShortCollectStateSuccess {
		g.Log(logger).Info(ctx, "QueryShortAndGrabAudit shortURL.CollectState != consts.ShortCollectStateSuccess")
		return
	}

	if err = s.GrabImageAudit(ctx, shortURL); err != nil {
		err = gerror.Wrap(err, "QueryShortAndGrabAudit GrabImageAudit failed")
		return
	}
	g.Log(logger).Info(ctx, "QueryShortAndGrabAudit end")
}

// GrabImageAudit the given command.
func (s *sShort) GrabImageAudit(ctx context.Context, shortURL *entity.ShortUrls) error {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-GrabImageAudit")
	defer span.End()

	var (
		logger      = helper.Helper().Logger(ctx)
		appEnv, err = env.New(ctx)
	)
	g.Log(logger).Info(ctx, "GrabImageAudit shortURL: ", shortURL)
	if err != nil {
		err = gerror.Wrap(err, "GrabImageAudit env.NewLark failed")
		return err
	}

	var (
		statusCode         int
		lastID             int64
		content            []byte
		now                = gtime.Now()
		trxID              = helper.Helper().InitTrxID(ctx, shortURL.AccountNo)
		filePathHTML       = "html/" + gconv.String(shortURL.ShortNo) + "/" + now.Format("Ymd") + "/audit/" + gconv.String(trxID)
		fileNameHTML       = filePathHTML + "/" + now.TimestampMicroStr() + "-" + grand.S(32) + ".html"
		filePathScreenshot = "screenshot/" + gconv.String(shortURL.ShortNo) + "/" + now.Format("Ymd") + "/audit/" + gconv.String(trxID)
		fileNameScreenshot = filePathScreenshot + "/" + now.TimestampMicroStr() + "-" + grand.S(32) + ".png"
		shortAudit         = &entity.ShortAuditLog{
			ShortNo: shortURL.ShortNo,
			TrxId:   trxID,
		}
		tx          gdb.TX
		shortMirror = (*entity.ShortMirror)(nil)
		cr          = &entity.ShortContentRecord{
			ShortNo:     shortURL.ShortNo,
			TrxId:       trxID,
			ContentType: consts.ContentTypeAudit,
		}
	)

	defer func() {
		g.Log(logger).Info(ctx, "GrabImageAudit defer statusCode:", statusCode)
		if statusCode == http.StatusFound || statusCode == http.StatusMovedPermanently {
			g.Log(logger).Info(ctx, "site redirect to:", shortURL.DestUrl, " statusCode:", statusCode)
			if _, er := dao.ShortUrls.Ctx(ctx).Where(do.ShortUrls{ShortNo: shortURL.ShortNo}).Update(g.Map{
				dao.ShortUrls.Columns().IsValid:     consts.ShortInvalid,
				dao.ShortUrls.Columns().DisableTime: gdb.Raw("current_timestamp(6)"),
				dao.ShortUrls.Columns().ModifyTime:  gdb.Raw("current_timestamp(6)"),
			}); er != nil {
				g.Log(logger).Error(ctx, "GrabImageAudit ShortUrls.Update failed:", er)
			}
		}
		if err != nil {
			val, errs := g.Redis(cache.RedisCache().ShortConn(ctx)).Do(ctx, "RPUSH",
				cache.RedisCache().ShortMirrorQueue(ctx), shortURL.ShortNo)
			if errs != nil {
				g.Log(logger).Error(ctx, "GrabImageAudit g.Redis(cache.RedisCache().ShortConn(ctx)).Conn(ctx) failed:", errs)
			}
			g.Log(logger).Info(ctx, "GrabImageAudit g.Redis(cache.RedisCache().ShortConn(ctx)).Conn(ctx) right lpush success:", val)
		}
	}()

	if err = dao.ShortMirror.Ctx(ctx).Scan(&shortMirror, "short_no = ?", shortURL.ShortNo); err != nil {
		err = gerror.Wrap(err, "GrabImageAudit ShortMirror.Scan failed")
		return err
	}

	if shortMirror == nil {
		g.Log(logger).Info(ctx, "GrabImageAudit ShortMirror.Scan no data")
		return nil
	}
	// 0、判断网页是否跳转
	if statusCode, err = s.RequestStatusCode(ctx, shortURL.DestUrl); err != nil {
		err = gerror.Wrap(err, "GrabImageAudit RequestStatusCode failed")
		return err
	}
	g.Log(logger).Info(ctx, "GrabImageAudit RequestStatusCode statusCode: ", statusCode)
	shortAudit.RedirectState = 100
	if statusCode != 200 {
		shortAudit.RedirectState = 200
	}

	// 1、抓起网页内容，
	if err = helper.Helper().CheckFileExists(ctx, appEnv.UploadPath(ctx)+filePathHTML); err != nil {
		g.Log(logger).Error(ctx, "GrabImageAudit CheckFileExists html failed:", err)
	}
	if content, err = s.RequestContent(ctx, shortURL.DestUrl, appEnv.UploadPath(ctx)+fileNameHTML); err != nil {
		err = gerror.Wrap(err, "GrabImageAudit RequestContent failed")
		return err
	}

	cr.Content = string(content)
	shortAudit.ContentPath = fileNameHTML
	shortAudit.HashContent = helper.Helper().Sha256Of(content)
	shortAudit.HashState = 100
	if shortAudit.HashContent != shortMirror.HashContent {
		shortAudit.HashState = 200
	}
	cr.HashContent = shortAudit.HashContent

	// 2、网页图片
	if err = helper.Helper().CheckFileExists(ctx, appEnv.UploadPath(ctx)+filePathScreenshot); err != nil {
		g.Log(logger).Error(ctx, "GrabImageAudit CheckFileExists Screenshot failed:", err)
	}
	if err = s.DownloadFullScreenshot(ctx, shortURL.DestUrl, appEnv.UploadPath(ctx)+fileNameScreenshot); err == nil {
		shortAudit.FullScreenshot = fileNameScreenshot
	} else {
		g.Log(logger).Error(ctx, "GrabImageAudit DownloadFullScreenshot failed:", err)
	}

	// 根据hash变化上报阿里云和腾讯云校验结果
	if shortAudit.HashState == 200 && shortAudit.FullScreenshot != "" {
		if err = s.ReportHashChange(ctx, shortAudit, appEnv.UploadPath(ctx)); err != nil {
			g.Log(logger).Error(ctx, "GrabImageAudit ReportHashChange failed:", err)
			err = gerror.Wrap(err, "GrabImageAudit ReportHashChange failed")
			// return err
		}
	}

	if tx, err = g.DB().Begin(ctx); err != nil {
		err = gerror.Wrap(err, "GrabImageAudit g.DB().Begin failed")
		return err
	}

	defer func() {
		if err != nil {
			if errs := tx.Rollback(); errs != nil {
				g.Log(logger).Error(ctx, "GrabImageAudit tx.Rollback failed:", errs)
			}
		} else {
			if errs := tx.Commit(); errs != nil {
				g.Log(logger).Info(ctx, "GrabImageAudit tx.Commit")
			}
		}
	}()

	if lastID, err = dao.ShortAuditLog.Ctx(ctx).TX(tx).OmitEmpty().Unscoped().InsertAndGetId(shortAudit); err != nil {
		err = gerror.Wrap(err, "GrabImageAudit AuditLog.InsertAndGetId failed")
		return err
	}
	g.Log(logger).Info(ctx, "GrabImageAudit AuditLog.InsertAndGetId lastID: ", lastID)

	if lastID, err = dao.ShortContentRecord.Ctx(ctx).TX(tx).OmitEmpty().Unscoped().InsertAndGetId(cr); err != nil {
		err = gerror.Wrap(err, "GrabImageAudit ShortContentRecord.Insert failed")
		return err
	}
	g.Log(logger).Info(ctx, "GrabImageAudit ShortContentRecord.Insert lastID: ", lastID)
	return nil
}

// ReportHashChange 上报阿里云和腾讯云校验结果
func (s *sShort) ReportHashChange(ctx context.Context, shortAudit *entity.ShortAuditLog, basePath string) error {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-GrabImageAudit")
	defer span.End()

	var (
		logger = helper.Helper().Logger(ctx)
		err    error
	)
	g.Log(logger).Info(ctx, "ReportHashChange shortAudit: ", shortAudit)

	if shortAudit.SafetyAuditTencent, err = tencent.Main(ctx, shortAudit.TrxId, basePath+shortAudit.FullScreenshot); err != nil {
		g.Log(logger).Error(ctx, "ReportHashChange tencent.Main failed:", err)
		err = gerror.Wrap(err, "ReportHashChange tencent.Main failed")
		return err
	}
	if shortAudit.SafetyAuditAlibaba, err = alibaba.Main(ctx, shortAudit.TrxId,
		basePath+shortAudit.FullScreenshot); err != nil {
		g.Log(logger).Error(ctx, "ReportHashChange alibaba.Main failed:", err)
		err = gerror.Wrap(err, "ReportHashChange alibaba.Main failed")
		return err
	}
	return nil
}
