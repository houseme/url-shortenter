package icp

const (
	authorizePath = "auth"
	queryPath     = "icpAbbreviateInfo/queryByCondition"

	authorizeContentType = "application/x-www-form-urlencoded;charset=UTF-8"
	queryContentType     = "application/json;charset=UTF-8"
)

// func Md5(str string) string {
// 	data := []byte(str)
// 	has := md5.Sum(data)
// 	return fmt.Sprintf("%x", has) // 将[]byte转成16进制
// }
//
// func GetHTTPResponse(resp *http.Response, url string, err error, result interface{}) error {
// 	body, err := GetHTTPResponseOrg(resp, url, err)
//
// 	if err == nil {
// 		// log.Println(string(body))
// 		err = json.Unmarshal(body, &result)
//
// 		if err != nil {
// 			log.Printf("请求接口%s解析json结果失败! ERROR: %s\n", url, err)
// 		}
// 	}
//
// 	return err
//
// }

// func GetHTTPResponseOrg(resp *http.Response, url string, err error) ([]byte, error) {
// 	if err != nil {
// 		log.Printf("请求接口%s失败! ERROR: %s\n", url, err)
// 		return nil, err
// 	}
//
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
//
// 	if err != nil {
// 		log.Printf("请求接口%s失败! ERROR: %s\n", url, err)
// 	}
//
// 	// 300及以上状态码都算异常
// 	if resp.StatusCode >= 300 {
// 		errMsg := fmt.Sprintf("请求接口 %s 失败! 返回内容: %s ,返回状态码: %d\n", url, string(body), resp.StatusCode)
// 		log.Printf(errMsg)
// 		err = fmt.Errorf(errMsg)
// 	}
//
// 	return body, err
// }
