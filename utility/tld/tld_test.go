/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package tld

import (
	"context"
	"testing"
)

var testUrls = []string{"www.google.com.hk", "www.discuz.net", "com",
	"www.discuz.vip", "www.ritto.shiga.jp", "ritto.shiga.jp", "mp.weixin.qq.com", "jonsen.yang.cn"}

func TestGetTld(t *testing.T) {
	ctx := context.Background()
	for _, url := range testUrls {
		ss, dd, tld := GetSubdomain(ctx, url, 2)
		t.Logf("resp：%s: %v, %s, %s\n", url, ss, dd, tld)
		resp, err := GetTLD(ctx, url, 0)
		if nil != err {
			t.Error("Failed get TLD:" + err.Error())
			return
		}
		t.Logf("resp：%s: %v, %s\n", url, resp.Tld, resp.Domain)
	}

	// t.Fail()
}

func BenchmarkGetTld(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		_, _ = GetTLD(ctx, "www.aaa.bbb.ccc.ddd.forease.com.cn", 0)
	}
}

func BenchmarkGetSubdomain(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		GetSubdomain(ctx, "www.aaa.bbb.ccc.ddd.forease.com.cn", 0)
	}
}