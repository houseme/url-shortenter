/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package filing

import (
	"context"
	"fmt"
	"os"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/gtrace"
	filing "github.com/houseme/icp-filing"
	"github.com/houseme/icp-filing/tld"
)

// QueryICP query icp for domain
func QueryICP(ctx context.Context, domain string) (resp *filing.QueryResponse, err error) {
	ctx, span := gtrace.NewSpan(ctx, "tracing-utility-filing-QueryICP")
	defer span.End()

	f := filing.New(ctx, filing.WithLogPath(os.TempDir()))
	var tldResp *tld.DomainTLDResp
	if tldResp, err = f.DomainTLD(ctx, domain, 0); err != nil {
		err = gerror.Wrap(err, "filing query DomainTLD failed")
		return
	}
	fmt.Println("tldResp:", tldResp)
	domain = tldResp.Domain
	if resp, err = f.DomainFilling(ctx, &filing.QueryRequest{
		UnitName: domain,
	}); err != nil {
		err = gerror.Wrap(err, "filing query DomainFilling failed")
		return
	}

	fmt.Println("resp:", resp)
	return
}
