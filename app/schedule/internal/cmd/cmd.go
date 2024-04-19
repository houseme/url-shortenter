// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package cmd

import (
	"context"
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"

	"github.com/houseme/url-shortenter/app/schedule/internal/consts"
	"github.com/houseme/url-shortenter/app/schedule/internal/service"
	"github.com/houseme/url-shortenter/utility/helper"
)

// Main is the main function of the program.
var Main = gcmd.Command{
	Name:  "main",
	Usage: "main",
	Brief: "start crontab job",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		logger := g.Log(consts.Logger)
		logger.Debug(ctx, `cron job start`)
		if _, err = gcron.Add(gctx.GetInitCtx(), "* * * * * *", func(ctx context.Context) {
			logger.Debug(ctx, `cron job running`)
		}); err != nil {
			return err
		}

		go service.Short().Execute(ctx)
		go service.Short().ExecuteAudit(ctx)
		if _, err = gcron.AddSingleton(ctx, "*/10 * * * * *", func(ctx context.Context) {
			ctx = helper.Helper().SetLogger(gctx.New(), consts.Logger)
			if err := service.Short().AssignTask(ctx); err != nil {
				logger.Errorf(ctx, "assign task error:%+v", err)
			}
			logger.Debug(ctx, "assign task success")
		}); err != nil {
			logger.Errorf(ctx, "assign task add cron job error:%+v", err)
			return err
		}

		if _, err = gcron.AddSingleton(ctx, "0 */5 * * * *", func(ctx context.Context) {
			ctx = helper.Helper().SetLogger(gctx.New(), consts.Logger)
			if err := service.Short().AuditAssignTask(ctx); err != nil {
				logger.Errorf(ctx, "Audit assign task error:%+v", err)
			}
			logger.Info(ctx, "Audit assign task success")
		}); err != nil {
			logger.Errorf(ctx, "Audit assign task add cron job error:%+v", err)
			return err
		}

		if _, err = gcron.AddSingleton(ctx, "*/30 * * * * *", func(ctx context.Context) {
			ctx = helper.Helper().SetLogger(gctx.New(), consts.Logger)
			// 针对访问记录入库的汇总处理
			if err := service.Short().ShortAccessLogSummary(ctx); err != nil {
				logger.Errorf(ctx, "Short Access Log Summary task error:%+v", err)
			}
			logger.Info(ctx, "Short Access Log Summary task success")
		}); err != nil {
			logger.Errorf(ctx, "Short Access Log Summary task add cron job error:%+v", err)
			return err
		}

		if _, err = gcron.AddSingleton(ctx, "*/15 * * * * *", func(ctx context.Context) {
			ctx = helper.Helper().SetLogger(gctx.New(), consts.Logger)
			// Access log information for inbound summary processing
			if err := service.Short().AccessLog(ctx); err != nil {
				logger.Errorf(ctx, "AccessLog task error:%+v", err)
			}
			logger.Info(ctx, "AccessLog task success")
		}); err != nil {
			logger.Errorf(ctx, "AccessLog task add cron job error:%+v", err)
			return err
		}

		// Register shutdown handler.
		gproc.AddSigHandlerShutdown(func(sig os.Signal) {
			g.Log().Info(ctx, `cron job shutdown`)
		})
		// Block listening to the shutdown signal.
		g.Listen()
		return
	},
}
