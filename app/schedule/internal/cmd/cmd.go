package cmd

import (
	"context"
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gproc"

	"github.com/houseme/url-shortenter/app/schedule/internal/consts"
	"github.com/houseme/url-shortenter/app/schedule/internal/service"
	"github.com/houseme/url-shortenter/utility/helper"
)

var (
	// Main is the main function of the program.
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start crontab job",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Log().Info(ctx, `cron job start`)
			// _, err = gcron.Add(ctx, "* * * * * *", func(ctx context.Context) {
			// 	g.Log().Debug(ctx, `cron job running`)
			// })
			// if err != nil {
			// 	return err
			// }

			go service.Short().Execute(ctx)
			go service.Short().ExecuteAudit(ctx)
			if _, err = gcron.AddSingleton(ctx, "*/10 * * * * *", func(ctx context.Context) {
				ctx = helper.Helper().SetLogger(context.Background(), consts.Logger)
				// 执行任务分配 处理镜像记录的任务
				if err := service.Short().AssignTask(ctx); err != nil {
					g.Log(consts.Logger).Info(ctx, "assign task error", err)
				}
				g.Log(consts.Logger).Info(ctx, "assign task success")
			}); err != nil {
				g.Log(consts.Logger).Error(ctx, "assign task add cron job error", err)
				return err
			}

			if _, err = gcron.AddSingleton(ctx, "0 */5 * * * *", func(ctx context.Context) {
				ctx = helper.Helper().SetLogger(context.Background(), consts.Logger)
				// 执行任务分配 处理网页内容信息跟踪的任务
				if err := service.Short().AuditAssignTask(ctx); err != nil {
					g.Log(consts.Logger).Info(ctx, "Audit assign task error", err)
				}
				g.Log(consts.Logger).Info(ctx, "Audit assign task success")
			}); err != nil {
				g.Log(consts.Logger).Error(ctx, "Audit assign task add cron job error", err)
				return err
			}

			if _, err = gcron.AddSingleton(ctx, "*/30 * * * * *", func(ctx context.Context) {
				ctx = helper.Helper().SetLogger(context.Background(), consts.Logger)
				// 针对访问记录入库的汇总处理
				if err := service.Short().ShortAccessLogSummary(ctx); err != nil {
					g.Log(consts.Logger).Info(ctx, "Short Access Log Summary task error", err)
				}
				g.Log(consts.Logger).Info(ctx, "Short Access Log Summary task success")
			}); err != nil {
				g.Log(consts.Logger).Error(ctx, "Short Access Log Summary task add cron job error", err)
				return err
			}

			if _, err = gcron.AddSingleton(ctx, "*/15 * * * * *", func(ctx context.Context) {
				ctx = helper.Helper().SetLogger(context.Background(), consts.Logger)
				// 访问log信息入库汇总处理
				if err := service.Short().AccessLog(ctx); err != nil {
					g.Log(consts.Logger).Info(ctx, "AccessLog task error", err)
				}
				g.Log(consts.Logger).Info(ctx, "AccessLog task success")
			}); err != nil {
				g.Log(consts.Logger).Error(ctx, "AccessLog task add cron job error", err)
				return err
			}

			// Register shutdown handler.
			gproc.AddSigHandlerShutdown(func(sig os.Signal) {
				g.Log().Info(ctx, `cron job shutdown`)
			})
			// Block listening the shutdown signal.
			g.Listen()
			return
		},
	}
)
