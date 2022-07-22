package icp

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
)

func TestWebQuery(t *testing.T) {
	icp, err := Query("yuerso.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(icp)

	icp2, err := QueryAiZhan("yuerso.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(icp2)
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
	resp, err := QueryDomainICP(gctx.New(), "wasair.com")
	if err != nil {
		t.Fatal("QueryDomainICP err:", err)
	}

	t.Log("QueryDomainICP response:", resp)
}
