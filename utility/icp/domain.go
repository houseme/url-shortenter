package icp

// QueryRequest 查询请求
type QueryRequest struct {
	PageNum  string `json:"pageNum"`
	PageSize string `json:"pageSize"`
	UnitName string `json:"unitName"`
}

// QueryResponse 查询响应
type QueryResponse struct {
	Code    int          `json:"code"`
	Msg     string       `json:"msg"`
	Success bool         `json:"success"`
	Params  *QueryParams `json:"params"`
}

// AuthParams 认证参数
type AuthParams struct {
	Bussiness string `json:"bussiness"`
	Expire    int64  `json:"expire"`
	Refresh   string `json:"refresh"`
}

// QueryParams 查询参数
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

// DomainInfo 域名信息
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

// AuthorizeRequest .授权请求
type AuthorizeRequest struct {
	AuthKey   string `json:"authKey"`
	TimeStamp string `json:"timeStamp"`
}

// AuthorizeResponse .授权响应
type AuthorizeResponse struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Success bool        `json:"success"`
	Params  *AuthParams `json:"params"`
}
