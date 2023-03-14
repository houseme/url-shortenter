// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package short

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"

	"github.com/houseme/url-shortenter/app/schedule/internal/consts"
	"github.com/houseme/url-shortenter/internal/database/dao"
	"github.com/houseme/url-shortenter/internal/database/model/do"
	"github.com/houseme/url-shortenter/internal/database/model/entity"
	"github.com/houseme/url-shortenter/utility/cache"
	"github.com/houseme/url-shortenter/utility/env"
	"github.com/houseme/url-shortenter/utility/helper"
)

// GrabImage grab image from url.
func (s *sShort) GrabImage(ctx context.Context, shortURL *entity.ShortUrls) error {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-GrabImage")
	defer span.End()

	var (
		log         = g.Log(helper.Helper().Logger(ctx))
		appEnv, err = env.New(ctx)
	)
	log.Info(ctx, "GrabImage shortURL: ", shortURL)
	if err != nil {
		err = gerror.Wrap(err, "GrabImage env.NewLark failed")
		return err
	}

	var (
		statusCode         int
		lastID             int64
		content            []byte
		now                = gtime.Now()
		filePathHTML       = "html/" + gconv.String(shortURL.ShortNo) + "/" + now.Format("Ymd") + "/mirror"
		fileNameHTML       = filePathHTML + "/" + now.TimestampMicroStr() + "-" + grand.S(32) + ".html"
		filePathScreenshot = "screenshot/" + gconv.String(shortURL.ShortNo) + "/" + now.Format("Ymd") + "/mirror"
		fileNameScreenshot = filePathScreenshot + "/" + now.TimestampMicroStr() + "-" + grand.S(32) + ".png"
		shortMirror        = &entity.ShortMirror{
			ShortNo: shortURL.ShortNo,
			DestUrl: shortURL.DestUrl,
		}
	)

	defer func() {
		if statusCode == http.StatusFound || statusCode == http.StatusMovedPermanently {
			log.Info(ctx, "site redirect to:", shortURL.DestUrl, " statusCode:", statusCode)
			if _, errs := dao.ShortUrls.Ctx(ctx).Where(do.ShortUrls{ShortNo: shortURL.ShortNo}).Update(g.Map{
				dao.ShortUrls.Columns().IsValid:     consts.ShortInvalid,
				dao.ShortUrls.Columns().DisableTime: gdb.Raw("current_timestamp(6)"),
				dao.ShortUrls.Columns().ModifyTime:  gdb.Raw("current_timestamp(6)"),
			}); errs != nil {
				log.Error(ctx, "GrabImage shortURL.ShortNo:", shortURL.ShortNo, " update failed:", errs)
			}
		}
	}()

	// 1、抓起网页内容，
	if err = helper.Helper().CheckFileExists(ctx, appEnv.UploadPath(ctx)+filePathHTML); err != nil {
		log.Error(ctx, "GrabImage CheckFileExists html failed:", err)
	}
	if statusCode, err = s.RequestStatusCode(ctx, shortURL.DestUrl); err != nil {
		err = gerror.Wrap(err, "GrabImage RequestStatusCode failed")
		return err
	}

	if content, err = s.RequestContent(ctx, shortURL.DestUrl, appEnv.UploadPath(ctx)+fileNameHTML); err != nil {
		err = gerror.Wrap(err, "GrabImage RequestContent failed")
		return err
	}
	if statusCode != 200 {
		log.Error(ctx, "GrabImage RequestContent statusCode:", statusCode)
		err = gerror.Wrap(err, "GrabImage RequestContent statusCode: "+gconv.String(statusCode))
		return err
	}
	var sct = &entity.ShortContentRecord{
		ShortNo:     shortURL.ShortNo,
		TrxId:       shortURL.ShortNo,
		ContentType: consts.ContentTypeMirror,
		Content:     string(content),
	}
	shortMirror.ContentPath = fileNameHTML
	shortMirror.HashContent = helper.Helper().Sha256Of(content)
	sct.HashContent = shortMirror.HashContent

	// 2、网页图片
	if err = helper.Helper().CheckFileExists(ctx, appEnv.UploadPath(ctx)+filePathScreenshot); err != nil {
		log.Error(ctx, "GrabImage CheckFileExists screenshot failed:", err)
	}
	if err = s.DownloadFullScreenshot(ctx, shortURL.DestUrl, appEnv.UploadPath(ctx)+fileNameScreenshot); err == nil {
		shortMirror.FullScreenshot = fileNameScreenshot
	} else {
		log.Error(ctx, "GrabImage DownloadFullScreenshot failed:", err)
	}
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		err = gerror.Wrap(err, "GrabImage g.DB().Begin failed")
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()
	if lastID, err = dao.ShortMirror.Ctx(ctx).TX(tx).OmitEmpty().Unscoped().InsertAndGetId(shortMirror); err != nil {
		err = gerror.Wrap(err, "GrabImage ShortMirror.InsertAndGetId failed")
		return err
	}
	log.Info(ctx, "GrabImage ShortMirror.InsertAndGetId lastID: ", lastID)

	if lastID, err = dao.ShortContentRecord.Ctx(ctx).TX(tx).OmitEmpty().Unscoped().InsertAndGetId(sct); err != nil {
		err = gerror.Wrap(err, "GrabImage ShortContentRecord.InsertAndGetId failed")
		return err
	}
	log.Info(ctx, "GrabImage ShortContentRecord.InsertAndGetId lastID: ", lastID)
	if _, err = dao.ShortUrls.Ctx(ctx).TX(tx).Where(do.ShortUrls{ShortNo: shortURL.ShortNo}).OmitEmpty().Unscoped().Update(g.Map{
		dao.ShortUrls.Columns().CollectState: consts.ShortCollectStateSuccess,
		dao.ShortUrls.Columns().CollectTime:  gdb.Raw("current_timestamp(6)"),
		dao.ShortUrls.Columns().ModifyTime:   gdb.Raw("current_timestamp(6)"),
	}); err != nil {
		err = gerror.Wrap(err, "GrabImage ShortUrls.Update failed")
		return err
	}
	return nil
}

// Execute the given command. 获取镜像信息
func (s *sShort) Execute(ctx context.Context) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-Execute")
	defer span.End()
	var (
		pool = grpool.New(10)
		log  = g.Log(helper.Helper().Logger(ctx))
	)

	log.Info(ctx, "Execute start")
	gtimer.SetInterval(ctx, time.Second, func(ctx context.Context) {
		ctx = helper.Helper().SetLogger(ctx, "schedule")
		log.Info(ctx, "Execute loop")
		for i := 0; i < 3; i++ {
			if err := pool.Add(ctx, s.queryShortAndGrab); err != nil {
				log.Error(ctx, "Execute pool.Add failed:", err)
			}
		}
		log.Info(ctx, "Execute loop end")
	})
	select {}
}

// queryShortAndGrab .
func (s *sShort) queryShortAndGrab(ctx context.Context) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-queryShortAndGrab")
	defer span.End()
	var (
		log       = g.Log(helper.Helper().Logger(ctx))
		conn, err = g.Redis(cache.RedisCache().ShortConn(ctx)).Conn(ctx)
		redisKey  = cache.RedisCache().ShortMirrorQueue(ctx) // 待抓取的镜像队列
	)

	if err != nil {
		log.Error(ctx, "queryShortAndGrab Redis failed:", err)
		return
	}

	defer func() {
		if err != nil {
			log.Error(ctx, "queryShortAndGrab defer error failed:", err)
		}
		// 关闭redis连接
		if err = conn.Close(ctx); err != nil {
			log.Error(ctx, "queryShortAndGrab Redis Close failed:", err)
		}
	}()

	log.Info(ctx, "queryShortAndGrab start")
	// 取出队列中的镜像
	var val *gvar.Var
	if val, err = conn.Do(ctx, "RPOP", redisKey); err != nil {
		err = gerror.Wrap(err, "queryShortAndGrab right pop failed")
		return
	}

	if val.IsNil() || val.IsEmpty() {
		log.Info(ctx, "queryShortAndGrab right pop is empty")
		return
	}
	log.Info(ctx, "queryShortAndGrab right pop success shortNo:", val.String())
	var (
		shortNo  = val.Uint64()
		shortKey = cache.RedisCache().ShortMirrorKey(ctx, val.String())
	)
	if val, err = conn.Do(ctx, "SETNX", shortKey, 1); err != nil {
		err = gerror.Wrap(err, "queryShortAndGrab setnx failed")
		return
	}
	if !val.IsNil() && !val.IsEmpty() && val.Int() < 1 {
		log.Info(ctx, "queryShortAndGrab setnx success shortNo:", val.String())
		return
	}
	defer func() {
		// 删除锁
		if _, err = conn.Do(ctx, "DEL", shortKey); err != nil {
			log.Error(ctx, "queryShortAndGrab del failed:", err)
		}
	}()
	var shortURL = (*entity.ShortUrls)(nil)
	if err = dao.ShortUrls.Ctx(ctx).Scan(&shortURL, "short_no = ?", shortNo); err != nil {
		err = gerror.Wrap(err, "queryShortAndGrab ShortUrls.Scan failed")
		return
	}

	if shortURL == nil {
		log.Info(ctx, "queryShortAndGrab ShortUrls.Scan no data")
		return
	}

	if shortURL.CollectState != consts.ShortCollectStateProcessing {
		log.Info(ctx, "queryShortAndGrab ShortUrls.Scan no data")
		return
	}

	if err = s.GrabImage(ctx, shortURL); err != nil {
		err = gerror.Wrap(err, "queryShortAndGrab GrabImage failed")
		return
	}
	log.Info(ctx, "queryShortAndGrab end")
}

// RequestContent request content from url.
func (s *sShort) RequestContent(ctx context.Context, url, fileName string) ([]byte, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-RequestContent")
	defer span.End()

	var (
		log    = g.Log(helper.Helper().Logger(ctx))
		r, err = g.Client().SetAgent(helper.Helper().UserAgent(ctx)).Get(ctx, url)
	)
	log.Info(ctx, "RequestContent start , url:", url, " fileName:", fileName)
	if err != nil {
		err = gerror.Wrap(err, "RequestContent failed")
		return nil, err
	}
	defer r.Body.Close()
	content := r.ReadAll()
	if err = os.WriteFile(fileName, content, 0o644); err != nil {
		err = gerror.Wrap(err, "RequestContent write file failed")
		return nil, err
	}
	return content, nil
}

// RequestStatusCode request content from url.
func (s *sShort) RequestStatusCode(ctx context.Context, url string) (int, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-RequestStatusCode")
	defer span.End()

	var (
		log    = g.Log(helper.Helper().Logger(ctx))
		client = g.Client().SetAgent(helper.Helper().UserAgent(ctx))
	)
	log.Info(ctx, "RequestStatusCode start url:", url)

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	r, err := client.Get(ctx, url)
	if err == http.ErrUseLastResponse {
		return r.StatusCode, nil
	}
	if err != nil {
		err = gerror.Wrap(err, "RequestStatusCode failed")
		return 0, err
	}
	r.Body.Close()
	log.Info(ctx, "RequestStatusCode end statusCode:", r.StatusCode)

	return r.StatusCode, nil
}

// DownloadFullScreenshot full screenshot.
func (s *sShort) DownloadFullScreenshot(ctx context.Context, url, fileName string) (err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-sShort-fullScreenshot")
	defer span.End()

	var (
		buf []byte
		log = g.Log(helper.Helper().Logger(ctx))
	)

	chtCtx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()
	log.Debug(ctx, "DownloadFullScreenshot start url:", url, " fileName:", fileName)
	// capture entire browser viewport, returning png with quality=90 // https://brank.as/
	if err = chromedp.Run(chtCtx, s.fullScreenshot(url, 90, &buf)); err != nil {
		err = gerror.Wrap(err, "fullScreenshot failed")
		return
	}
	if err = os.WriteFile(fileName, buf, 0o644); err != nil {
		err = gerror.Wrap(err, "fullScreenshot write file failed")
		return
	}
	log.Info(ctx, "wrote fullScreenshot successfully to: ", fileName)

	return nil
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings.
// Use device.Reset to reset the emulation and viewport settings.
func (s *sShort) fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
