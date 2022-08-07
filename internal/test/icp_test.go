package test

import (
	"testing"

	"github.com/houseme/url-shortenter/utility/icp"
)

func TestWebQuery(t *testing.T) {
	queryResp, err := icp.Query("yuerso.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(queryResp)

	queryResp, err = icp.QueryAiZhan("yuerso.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(queryResp)
}

func TestIcp_Query(t *testing.T) {
	// icp := &ICP{}
	// domainInfo, err := icp.Query("wasair.com")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// info, _ := json.Marshal(domainInfo)
	// fmt.Println(string(info))
	// t.Log(domainInfo, "info", string(info))
}

func TestQueryDomainICP(t *testing.T) {
	// resp, err := icp.QueryDomainICP(gctx.New(), "wasair.com")
	// if err != nil {
	// 	t.Fatal("QueryDomainICP err:", err)
	// }
	//
	// t.Log("QueryDomainICP response:", resp)
}
