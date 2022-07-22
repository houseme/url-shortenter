package icp

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"

	"github.com/houseme/url-shortenter/utility"
)

// QueryResp is a struct for icp data
type QueryResp struct {
	IcpNumber string `json:"icp_number"`
	IcpName   string `json:"icp_name"`
	Attr      string `json:"attr"`
	Date      string `json:"date"`
}

// Query is a function for icp query
func Query(url string) (*QueryResp, error) {
	url = "https://icp.chinaz.com/" + url
	client := &http.Client{Timeout: 5 * time.Second}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("user-agent", RandomUserAgent())

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	gp, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	icp := &QueryResp{}
	gp.Find("#first > li").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			icp.IcpName = strings.TrimSpace(s.Find("p").Text())
		}

		if i == 1 {
			icp.Attr = strings.TrimSpace(s.Find("p").Text())
		}

		if i == 2 {
			icp.IcpNumber = strings.TrimSpace(s.Find("p > font").Text())
		}

		if i == 6 {
			icp.Date = strings.TrimSpace(s.Find("p").Text())
		}
	})

	if icp.IcpName == "" {
		return icp, fmt.Errorf("没有查询到备案信息")
	}
	return icp, nil
}

// QueryAiZhan is a function for icp query
func QueryAiZhan(url string) (*QueryResp, error) {
	url = "https://icp.aizhan.com/geticp/?host=" + url
	client := &http.Client{Timeout: 5 * time.Second}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("user-agent", RandomUserAgent())

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("没有查询到备案信息")
	}

	reg, err := regexp.Compile(`document\.write\('(.*?)'\);`)
	if err != nil {
		return nil, fmt.Errorf("没有查询到备案信息")
	}

	result := reg.FindAllStringSubmatch(string(all), -1)
	icp := &QueryResp{}
	// 过滤<></>
	for _, text := range result {
		if len(text) >= 2 {
			if text[1] == "未找到备案信息" {
				return nil, fmt.Errorf("没有查询到备案信息")
			}
			icp.IcpNumber = text[1]
			return icp, nil
		}
	}

	return nil, fmt.Errorf("没有查询到备案信息")
}

var uaGens = []func() string{
	firefoxUserAgent,
	chromeUserAgent,
}

// RandomUserAgent generates a random browser user agent on every request
func RandomUserAgent() string {
	rand.Seed(time.Now().Unix())
	return uaGens[rand.Intn(len(uaGens))]()
}

var ffVersions = []float32{
	58.0,
	57.0,
	56.0,
	52.0,
	48.0,
	40.0,
	35.0,
}

var chromeVersions = []string{
	"65.0.3325.146",
	"64.0.3282.0",
	"41.0.2228.0",
	"40.0.2214.93",
	"37.0.2062.124",
}

var osStrings = []string{
	"Macintosh; Intel Mac OS X 10_10",
	"Windows NT 10.0",
	"Windows NT 5.1",
	"Windows NT 6.1; WOW64",
	"Windows NT 6.1; Win64; x64",
	"X11; Linux x86_64",
}

func firefoxUserAgent() string {
	version := ffVersions[rand.Intn(len(ffVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s; rv:%.1f) Gecko/20100101 Firefox/%.1f", os, version, version)
}

func chromeUserAgent() string {
	version := chromeVersions[rand.Intn(len(chromeVersions))]
	os := osStrings[rand.Intn(len(osStrings))]
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", os, version)
}

// ErrNotForRecord .
// var ErrNotForRecord = errors.New("域名未备案")

// ICP number
type ICP struct {
	token string
	ip    string
}

// // Query is a function for icp query
// func (i *ICP) Query(domain string) (*DomainInfo, error) {
// 	i.getIP()
// 	if err := i.auth(); err != nil {
// 		return nil, err
// 	}
//
// 	return i.query(domain)
// }
//
// func (i *ICP) query(domain string) (*DomainInfo, error) {
// 	queryRequest, _ := json.Marshal(&QueryRequest{
// 		UnitName: domain,
// 	})
//
// 	result := &QueryResponse{Params: &QueryParams{}}
// 	err := i.post("icpAbbreviateInfo/queryByCondition", bytes.NewReader(queryRequest), "application/json;charset=UTF-8", i.token, result)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if !result.Success {
// 		return nil, fmt.Errorf("查询：%s", result.Msg)
// 	}
//
// 	queryParams := result.Params
// 	if len(queryParams.List) == 0 {
// 		return nil, ErrNotForRecord
// 	}
//
// 	return queryParams.List[0], nil
// }
//
// // auth .
// func (i *ICP) auth() error {
// 	timestamp := time.Now().Unix()
// 	authKey := Md5(fmt.Sprintf("testtest%d", timestamp))
// 	authBody := fmt.Sprintf("authKey=%s&timeStamp=%d", authKey, timestamp)
//
// 	result := &AuthorizeResponse{Params: &AuthParams{}}
// 	err := i.post("auth", bytes.NewReader([]byte(authBody)), "application/x-www-form-urlencoded;charset=UTF-8", "0", result)
// 	if err != nil {
// 		return err
// 	}
//
// 	if !result.Success {
// 		return fmt.Errorf("获取token失败：%s", result.Msg)
// 	}
//
// 	authParams := result.Params
// 	i.token = authParams.Bussiness
//
// 	return nil
// }
//
// // post .
// func (i *ICP) post(url string, data io.Reader, contentType string, token string, result interface{}) error {
// 	postURL := fmt.Sprintf("https://hlwicpfwc.miit.gov.cn/icpproject_query/api/%s", url)
// 	queryReq, err := http.NewRequest(http.MethodPost, postURL, data)
// 	if err != nil {
// 		return err
// 	}
//
// 	queryReq.Header.Set("Content-Type", contentType)
// 	queryReq.Header.Set("Origin", "https://beian.miit.gov.cn/")
// 	queryReq.Header.Set("Referer", "https://beian.miit.gov.cn/")
// 	queryReq.Header.Set("token", token)
// 	queryReq.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36")
// 	queryReq.Header.Set("CLIENT_IP", i.ip)
// 	queryReq.Header.Set("X-FORWARDED-FOR", i.ip)
//
// 	resp, err := http.DefaultClient.Do(queryReq)
// 	fmt.Sprintf("http.DefaultClient.Do %+v", resp)
// 	return GetHTTPResponse(resp, postURL, err, result)
// }

// Request .
func (i *ICP) Request(ctx context.Context, contentType, path string, params interface{}) ([]byte, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-icp-Request")
	defer span.End()
	var (
		header = g.MapStrStr{
			"Content-Type":    contentType,
			"Origin":          "https://beian.miit.gov.cn/",
			"Referer":         "https://beian.miit.gov.cn/",
			"token":           i.token,
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36",
			"CLIENT_IP":       i.ip,
			"X-FORWARDED-FOR": i.ip,
		}
		logger = utility.Helper().Logger(ctx)
	)
	g.Log(logger).Debug(ctx, "icp-Request header", header, " params:", params)
	client := g.Client().SetHeaderMap(header)
	if path != authorizePath {
		client = client.ContentJson()
	}
	resp, err := client.Post(ctx, "https://hlwicpfwc.miit.gov.cn/icpproject_query/api/"+path, params)
	if err != nil {
		g.Log(logger).Error(ctx, "icp-Request error", err)
		return nil, err
	}
	defer func() {
		_ = resp.Close()
	}()
	content := resp.ReadAll()
	g.Log(logger).Debug(ctx, "icp-Request content", string(content))
	g.Log(logger).Debug(ctx, "icp-Request all: \n")
	resp.RawDump()
	// 300及以上状态码都算异常
	if resp.StatusCode >= http.StatusMultipleChoices {
		errMsg := "请求接口 " + path + " 失败! ,返回状态码: " + gconv.String(resp.StatusCode) + " 返回内容: " + string(content)
		g.Log(logger).Debug(ctx, "icp-Request StatusCode error", errMsg)
		err = gerror.New(errMsg)
		return nil, err
	}

	return content, nil
}

func (i *ICP) getIP() {
	if i.ip != "" {
		return
	}
	i.ip = "101." + gconv.String(grand.N(1, 255)) + "." + gconv.String(grand.N(1, 255)) + "." + gconv.String(grand.N(1, 255))
}

// authorize .
func (i *ICP) authorize(ctx context.Context) error {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-icp-authorize")
	defer span.End()

	var (
		logger    = utility.Helper().Logger(ctx)
		timestamp = gtime.Now().TimestampStr()
		authKey   = i.Md5("testtest" + timestamp)
		req       = &AuthorizeRequest{
			AuthKey:   authKey,
			TimeStamp: timestamp,
		}
	)
	g.Log(logger).Debug(ctx, "icp-authorize req", req)
	resp, err := i.Request(ctx, authorizeContentType, authorizePath, req)
	if err != nil {
		err = gerror.Wrap(err, "icp-authorize request error")
		return err
	}
	var authResp *AuthorizeResponse
	if err = gjson.New(resp).Scan(&authResp); err != nil {
		err = gerror.Wrap(err, "icp-authorize json scan  error")
		return err
	}
	g.Log(logger).Debug(ctx, "icp-authorize resp", authResp)
	if !authResp.Success {
		err = gerror.New(authResp.Msg)
		return err
	}
	i.token = authResp.Params.Bussiness
	g.Log(logger).Debug(ctx, "icp-authorize end ICP:", i.String())
	return nil
}

// QueryICP .
func (i *ICP) QueryICP(ctx context.Context, domain string) (*QueryResponse, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-icp-QueryICP")
	defer span.End()
	var (
		logger = utility.Helper().Logger(ctx)
		req    = &QueryRequest{
			UnitName: domain,
		}
	)
	g.Log(logger).Debug(ctx, "icp-QueryICP req", req)
	resp, err := i.Request(ctx, queryContentType, queryPath, req)
	if err != nil {
		err = gerror.Wrap(err, "icp-QueryICP request error")
		return nil, err
	}
	var queryResp *QueryResponse
	if err = gjson.New(resp).Scan(&queryResp); err != nil {
		err = gerror.Wrap(err, "icp-QueryICP json scan  error")
		return nil, err
	}
	g.Log(logger).Debug(ctx, "icp-QueryICP resp", queryResp)
	return queryResp, nil
}

// Md5 .
func (i *ICP) Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) // 将[]byte转成16进制
}

// String .
func (i *ICP) String() string {
	return `{"ip":"` + i.ip + `","token":"` + i.token + `"}`
}

// QueryDomainICP .
func QueryDomainICP(ctx context.Context, domain string) (*QueryResponse, error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-icp-QueryDomainICP")
	defer span.End()
	i := &ICP{
		token: "0",
	}
	i.getIP()
	err := i.authorize(ctx)
	if err != nil {
		return nil, err
	}
	return i.QueryICP(ctx, domain)
}
