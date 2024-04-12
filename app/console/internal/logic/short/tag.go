/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package short

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/houseme/url-shortenter/app/console/internal/consts"
	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/internal/database/dao"
	"github.com/houseme/url-shortenter/internal/database/model/do"
	"github.com/houseme/url-shortenter/internal/database/model/entity"
	"github.com/houseme/url-shortenter/utility/helper"
)

// CreateTag is the handler for CreateTag.
func (s *sShort) CreateTag(ctx context.Context, in *model.CreateTagInput) (out *model.CreateTagOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-CreateTag")
	defer span.End()

	var (
		logger  = g.Log(helper.Helper().Logger(ctx))
		tagInfo = (*entity.UsersTagRelation)(nil)
	)
	logger.Debug(ctx, "short-CreateTag in:", in)
	// 查询用户是否与标签绑定关系
	if err = dao.UsersTagRelation.Ctx(ctx).Scan(&tagInfo, entity.ShortTag{
		TagName: in.Name,
	}); err != nil {
		return
	}
	// 如果没有绑定关系则创建
	if tagInfo == nil {
		var (
			baseTag = (*entity.ShortTag)(nil)
			TagNo   = helper.Helper().InitTrxID(ctx, in.AuthUserNo)
		)
		// 先查询标签是否全局存在
		if err = dao.ShortTag.Ctx(ctx).Scan(&baseTag, entity.ShortTag{
			TagName: in.Name,
		}); err != nil {
			return
		}
		// 如果标签存在就覆盖 tagNo 的值
		if baseTag != nil {
			TagNo = baseTag.TagNo
		}

		if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			var lastID int64
			// 如果标签不存在则创建
			if baseTag == nil {
				if lastID, err = dao.ShortTag.Ctx(ctx).TX(tx).OmitEmpty().Unscoped().InsertAndGetId(do.ShortTag{
					TagNo:   TagNo,
					TagName: in.Name,
					State:   consts.TagStateNormal,
				}); err != nil {
					return
				}
				logger.Debug(ctx, "short-CreateTag insert short tag db lastID:", lastID)
			}
			// 创建用户与标签的关系
			if lastID, err = dao.UsersTagRelation.Ctx(ctx).TX(tx).OmitEmpty().Unscoped().InsertAndGetId(do.UsersTagRelation{
				TagNo:  TagNo,
				UserNo: in.AuthAccountNo,
				State:  consts.TagStateNormal,
			}); err != nil {
				return
			}
			logger.Debug(ctx, "short-CreateTag insert users tag relation db lastID:", lastID)
			return nil
		}); err != nil {
			return
		}
		tagInfo = &entity.UsersTagRelation{
			TagNo: TagNo,
		}
	}
	out = &model.CreateTagOutput{
		TagNo: tagInfo.TagNo,
	}
	return
}

// DelTag is the handler for DelTag.
func (s *sShort) DelTag(ctx context.Context, in *model.DelTagInput) (out *model.DelTagOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-DelTag")
	defer span.End()

	var (
		logger  = g.Log(helper.Helper().Logger(ctx))
		tagInfo = (*entity.ShortTagRelation)(nil)
	)
	logger.Debug(ctx, "short-DelTag in:", in)
	// 查询句子和标签的关系
	if err = dao.ShortTagRelation.Ctx(ctx).Scan(&tagInfo, do.ShortTagRelation{
		UserNo:  in.AuthUserNo,
		TagNo:   in.TagNo,
		ShortNo: in.ShortNo,
	}); err != nil {
		return
	}

	// 如果不存在，直接提示错误
	if tagInfo == nil {
		logger.Debug(ctx, "short-DelTag tag not found")
		out = &model.DelTagOutput{
			ResultCode: consts.ResponseCodeFailed,
			ResultMsg:  "标签关系不存在",
		}
		return
	}

	// 删除句子和标签的关系，设置为失效状态
	var lastID int64
	if lastID, err = dao.ShortTagRelation.Ctx(ctx).OmitEmpty().Unscoped().UpdateAndGetAffected(do.ShortTag{
		State:      consts.TagStateInvalid,
		ModifyTime: gtime.Now(),
	}, do.ShortTagRelation{
		Id: tagInfo.Id,
	}); err != nil {
		return
	}
	logger.Debug(ctx, "short-DelTag update short tag db lastID:", lastID)

	out = &model.DelTagOutput{
		ResultCode: consts.ResponseCodeSuccess,
		ResultMsg:  consts.ResponseMsgSuccess,
	}

	return
}

// AddTag is the handler for AddTag.
func (s *sShort) AddTag(ctx context.Context, in *model.AddTagInput) (out *model.AddTagOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-short-AddTag")
	defer span.End()

	var (
		logger  = g.Log(helper.Helper().Logger(ctx))
		tagInfo = (*entity.ShortTagRelation)(nil)
	)
	logger.Debug(ctx, "short-AddTag in:", in)
	if err = dao.ShortTagRelation.Ctx(ctx).Scan(&tagInfo, do.ShortTagRelation{
		UserNo:  in.AuthUserNo,
		TagNo:   in.TagNo,
		ShortNo: in.ShortNo,
	}); err != nil {
		return
	}

	// 如果不存在，直接提示错误
	if tagInfo != nil {
		logger.Debug(ctx, "short-AddTag tag not found")
		out = &model.AddTagOutput{
			ResultCode: consts.ResponseCodeSuccess,
			ResultMsg:  consts.ResponseMsgSuccess,
		}
		return
	}
	var userTagRelation = (*entity.UsersTagRelation)(nil)
	if err = dao.UsersTagRelation.Ctx(ctx).Scan(&userTagRelation, do.UsersTagRelation{
		UserNo: in.AuthUserNo,
		TagNo:  in.TagNo,
	}); err != nil {
		return
	}

	if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var lastID int64
		if lastID, err = dao.ShortTagRelation.Ctx(ctx).TX(tx).OmitEmpty().Unscoped().InsertAndGetId(do.ShortTagRelation{
			UserNo:     in.AuthUserNo,
			TagNo:      in.TagNo,
			ShortNo:    in.ShortNo,
			State:      consts.TagStateNormal,
			CreateTime: gtime.Now(),
			ModifyTime: gtime.Now(),
		}); err != nil {
			return
		}
		logger.Debug(ctx, "short-AddTag insert short tag relation db lastID:", lastID)
		if userTagRelation == nil {
			if lastID, err = dao.UsersTagRelation.Ctx(ctx).TX(tx).OmitEmpty().Unscoped().InsertAndGetId(do.UsersTagRelation{
				UserNo:     in.AuthUserNo,
				TagNo:      in.TagNo,
				State:      consts.TagStateNormal,
				CreateTime: gtime.Now(),
				ModifyTime: gtime.Now(),
			}); err != nil {
				return
			}
			logger.Debug(ctx, "short-AddTag insert users tag relation db lastID:", lastID)
		}
		return nil
	}); err != nil {
		logger.Errorf(ctx, "short-AddTag insert users tag relation db error:%+v", err)
		out = &model.AddTagOutput{
			ResultCode: consts.ResponseCodeFailed,
			ResultMsg:  err.Error(),
		}
		return
	}
	out = &model.AddTagOutput{
		ResultCode: consts.ResponseCodeSuccess,
		ResultMsg:  consts.ResponseMsgSuccess,
	}

	return
}
