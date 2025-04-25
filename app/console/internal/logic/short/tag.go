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
		tagInfo = (*entity.UserTagRelation)(nil)
	)
	logger.Debug(ctx, "short-CreateTag in:", in)
	// query whether the user is bound to the tag
	if err = dao.UserTagRelation.Ctx(ctx).Scan(&tagInfo, entity.ShortTag{
		TagName: in.Name,
	}); err != nil {
		return
	}
	// create if there is no binding relationship
	if tagInfo == nil {
		var (
			baseTag = (*entity.ShortTag)(nil)
			TagNo   = helper.Helper().InitTrxID(ctx, in.AuthUserNo)
		)
		// first check whether the tag exists globally
		if err = dao.ShortTag.Ctx(ctx).Scan(&baseTag, entity.ShortTag{
			TagName: in.Name,
		}); err != nil {
			return
		}
		// if the tag exists, overwrite the value of tag no
		if baseTag != nil {
			TagNo = baseTag.TagNo
		}

		if err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			var lastID int64
			// create if the tag does not exist
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
			// create a user tagged relationship
			if lastID, err = dao.UserTagRelation.Ctx(ctx).TX(tx).OmitEmpty().Unscoped().InsertAndGetId(do.UserTagRelation{
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
		tagInfo = &entity.UserTagRelation{
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
	// Query the relationship between sentences and labels
	if err = dao.ShortTagRelation.Ctx(ctx).Scan(&tagInfo, do.ShortTagRelation{
		UserNo:  in.AuthUserNo,
		TagNo:   in.TagNo,
		ShortNo: in.ShortNo,
	}); err != nil {
		return
	}

	// If it does not exist, it will directly prompt an error
	if tagInfo == nil {
		logger.Debug(ctx, "short-DelTag tag not found")
		out = &model.DelTagOutput{
			ResultCode: consts.ResponseCodeFailed,
			ResultMsg:  "标签关系不存在",
		}
		return
	}

	// Delete the relationship between sentences and tags and set them to invalid state
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

	// If it does not exist, it will directly prompt an error
	if tagInfo != nil {
		logger.Debug(ctx, "short-AddTag tag not found")
		out = &model.AddTagOutput{
			ResultCode: consts.ResponseCodeSuccess,
			ResultMsg:  consts.ResponseMsgSuccess,
		}
		return
	}
	var userTagRelation = (*entity.UserTagRelation)(nil)
	if err = dao.UserTagRelation.Ctx(ctx).Scan(&userTagRelation, do.UserTagRelation{
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
			if lastID, err = dao.UserTagRelation.Ctx(ctx).TX(tx).OmitEmpty().Unscoped().InsertAndGetId(do.UserTagRelation{
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
