// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package icp

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

// AuthParams auth params
type AuthParams struct {
	Bussiness string `json:"bussiness"`
	Expire    int64  `json:"expire"`
	Refresh   string `json:"refresh"`
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
