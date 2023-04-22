/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package access

import (
	"context"

	"github.com/houseme/url-shortenter/app/console/internal/model"
	"github.com/houseme/url-shortenter/app/console/internal/service"
)

type sAccess struct {
}

// init
func init() {
	service.RegisterAccess(&sAccess{})
}

// List is the handler for List.
func (s *sAccess) List(ctx context.Context, in *model.AccessListInput) (out *model.AccessListOutput, err error) {
	return
}

// Detail is the handler for Detail.
func (s *sAccess) Detail(ctx context.Context, in *model.AccessDetailInput) (out *model.AccessDetailOutput, err error) {
	return
}
