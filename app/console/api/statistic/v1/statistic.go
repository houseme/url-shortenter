/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/houseme/url-shortenter/app/console/internal/model"
)

// ListReq is the request structure for statistic list.
type ListReq struct {
	g.Meta `path:"/statistic/list" tags:"Statistic Service" method:"Get" summary:"short url statistic list"`
	*model.StatisticListInput
}

// ListRes is the response structure for the statistic list.
type ListRes struct {
	*model.StatisticListOutput
}

// DetailReq is the request structure for statistic detail.
type DetailReq struct {
	g.Meta `path:"/statistic/detail" tags:"Statistic Service" method:"Get" summary:"short url statistic detail"`
	*model.StatisticDetailInput
}

// DetailRes is the response structure for the statistic detail.
type DetailRes struct {
	*model.StatisticDetailOutput
}
