/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package model

// StatisticListInput is the input of StatisticList.
type StatisticListInput struct {
	*Base `json:"-"`
	*PageInfoInput
}

// StatisticListOutput is the output of StatisticList.
type StatisticListOutput struct {
	List     []*StatisticItem `json:"list" description:"短链访问统计列表"`
	Paginate *PageInfoOutput  `json:"paginate" description:"分页信息"`
}

// StatisticItem is the item of statistic.
type StatisticItem struct {
	ID       uint64 `json:"id,string" description:"访问 log ID"`
	ShortNo  uint64 `json:"shortNo,string" description:"短链编号"`
	ShortURL string `json:"shortUrl" description:"短链"`
}

// StatisticDetailInput is the input of StatisticDetail.
type StatisticDetailInput struct {
	*Base `json:"-"`
}

// StatisticDetailOutput is the output of StatisticDetail.
type StatisticDetailOutput struct {
	ID       uint64 `json:"id,string" description:"访问 log ID"`
	ShortNo  uint64 `json:"shortNo,string" description:"短链编号"`
	ShortURL string `json:"shortUrl" description:"短链"`
}
