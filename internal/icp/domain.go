// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package icp

import (
	"github.com/gogf/gf/v2/util/gconv"
)

const (
	authorizePath = "auth"
	queryPath     = "icpAbbreviateInfo/queryByCondition"

	authorizeContentType = "application/x-www-form-urlencoded;charset=UTF-8"
	queryContentType     = "application/json;charset=UTF-8"
)

// QueryRequest query request
type QueryRequest struct {
	PageNum  string `json:"pageNum"`
	PageSize string `json:"pageSize"`
	UnitName string `json:"unitName"`
}

// QueryResponse query response
type QueryResponse struct {
	Code    int          `json:"code"`
	Msg     string       `json:"msg"`
	Success bool         `json:"success"`
	Params  *QueryParams `json:"params"`
}

// String string
func (r *QueryResponse) String() string {
	return `{"code": ` + gconv.String(r.Code) + `, "msg": "` + r.Msg + `", "success": ` + gconv.String(r.Success) + `, "params": ` + r.Params.String() + `}`
}

// AuthParams auth params
type AuthParams struct {
	Business string `json:"bussiness"`
	Expire   int64  `json:"expire"`
	Refresh  string `json:"refresh"`
}

// QueryParams query params
type QueryParams struct {
	EndRow           int           `json:"endRow"`
	FirstPage        int           `json:"firstPage"`
	HasNextPage      bool          `json:"hasNextPage"`
	HasPreviousPage  bool          `json:"hasPreviousPage"`
	IsFirstPage      bool          `json:"isFirstPage"`
	IsLastPage       bool          `json:"isLastPage"`
	LastPage         int           `json:"lastPage"`
	List             []*DomainInfo `json:"list"`
	NavigatePages    int           `json:"navigatePages"`
	NavigatepageNums []int         `json:"navigatepageNums"`
	NextPage         int           `json:"nextPage"`
	PageNum          int           `json:"pageNum"`
	PageSize         int           `json:"pageSize"`
	Pages            int           `json:"pages"`
	PrePage          int           `json:"prePage"`
	Size             int           `json:"size"`
	StartRow         int           `json:"startRow"`
	Total            int           `json:"total"`
}

// String string
func (r *QueryParams) String() string {
	return `{"endRow": ` + gconv.String(r.EndRow) + `, "firstPage": ` + gconv.String(r.FirstPage) + `, "hasNextPage": ` + gconv.String(r.HasNextPage) + `, "hasPreviousPage": ` + gconv.String(r.HasPreviousPage) + `, "isFirstPage": ` + gconv.String(r.IsFirstPage) + `, "isLastPage": ` + gconv.String(r.IsLastPage) + `, "lastPage": ` + gconv.String(r.LastPage) + `, "list": ` + gconv.String(r.List) + `, "navigatePages": ` + gconv.String(r.NavigatePages) + `, "navigatepageNums": ` + gconv.String(r.NavigatepageNums) + `, "nextPage": ` + gconv.String(r.NextPage) + `, "pageNum": ` + gconv.String(r.PageNum) + `, "pageSize": ` + gconv.String(r.PageSize) + `, "pages": ` + gconv.String(r.Pages) + `, "prePage": ` + gconv.String(r.PrePage) + `, "size": ` + gconv.String(r.Size) + `, "startRow": ` + gconv.String(r.StartRow) + `, "total": ` + gconv.String(r.Total) + `}`
}

// DomainInfo domain info
type DomainInfo struct {
	ContentTypeName  string `json:"contentTypeName"`
	Domain           string `json:"domain"`
	DomainID         int64  `json:"domainId"`
	HomeURL          string `json:"homeUrl"`
	LeaderName       string `json:"leaderName"`
	LimitAccess      string `json:"limitAccess"`
	MainID           int64  `json:"mainId"`
	MainLicence      string `json:"mainLicence"`
	NatureName       string `json:"natureName"`
	ServiceID        int64  `json:"serviceId"`
	ServiceLicence   string `json:"serviceLicence"`
	ServiceName      string `json:"serviceName"`
	UnitName         string `json:"unitName"`
	UpdateRecordTime string `json:"updateRecordTime"`
}

// String string
func (r *DomainInfo) String() string {
	return `{"contentTypeName": "` + r.ContentTypeName + `", "domain": "` + r.Domain + `", "domainId": ` + gconv.String(r.DomainID) + `, "homeUrl": "` + r.HomeURL + `", "leaderName": "` + r.LeaderName + `", "limitAccess": "` + r.LimitAccess + `", "mainId": ` + gconv.String(r.MainID) + `, "mainLicence": "` + r.MainLicence + `", "natureName": "` + r.NatureName + `", "serviceId": ` + gconv.String(r.ServiceID) + `, "serviceLicence": "` + r.ServiceLicence + `", "serviceName": "` + r.ServiceName + `", "unitName": "` + r.UnitName + `", "updateRecordTime": "` + r.UpdateRecordTime + `}`
}

// AuthorizeRequest authorize request
type AuthorizeRequest struct {
	AuthKey   string `json:"authKey"`
	TimeStamp string `json:"timeStamp"`
}

// AuthorizeResponse authorize response
type AuthorizeResponse struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
	Params  *AuthParams `json:"params"`
}

// QueryResp is a struct for icp data
type QueryResp struct {
	IcpNumber string `json:"icp_number"`
	IcpName   string `json:"icp_name"`
	Attr      string `json:"attr"`
	Date      string `json:"date"`
}
