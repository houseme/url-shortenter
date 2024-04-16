/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

// Package access logic
package access

import (
	"context"

	"github.com/gogf/gf/v2/net/gtrace"

	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

type sAccess struct{}

// init
func init() {
	service.RegisterAccess(&sAccess{})
}

// List is the handler for List.
func (s *sAccess) List(ctx context.Context, in *model.AccessListInput) (out *model.AccessListOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-access-List")
	defer span.End()

	return
}

// Detail is the handler for Detail.
func (s *sAccess) Detail(ctx context.Context, in *model.AccessDetailInput) (out *model.AccessDetailOutput, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-logic-access-Detail")
	defer span.End()

	return
}
