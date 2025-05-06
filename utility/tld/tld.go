/*
 *  Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
 *
 *  This Source Code Form is subject to the terms of the MIT License.
 *  If a copy of the MIT was not distributed with this file,
 *  You can obtain one at https://github.com/houseme/url-shortenter.
 */

package tld

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"sync"
)

var (
	tldMap = make(map[string]DomainTLD, 9568)
	pool   = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}
)

// Initialization Top Level Domain Table
func init() {
	initTld()
}

// GetTLD get the domain name and TLD
func GetTLD(ctx context.Context, url string, level int) (resp *DomainTLDResp, err error) {
	return parseDomainTLD(ctx, url, level)
}

// GetSubdomain get a subdomain from URL
func GetSubdomain(ctx context.Context, url string, level int) (subdomain, domain, tld string) {
	resp, err := parseDomainTLD(ctx, url, level)
	if err != nil {
		return "", "", ""
	}
	return resp.SubDomain, resp.Domain, resp.Tld
}

// parseDomainTLD parse domain TLD
func parseDomainTLD(_ context.Context, url string, level int) (resp *DomainTLDResp, err error) {
	var (
		buffer = pool.Get().(*bytes.Buffer)
		dm     = strings.Split(url, ".")
		size   = len(dm)
		isTLD  bool
	)
	resp = &DomainTLDResp{
		Link: url,
	}
	if size > 1 {
		idx := 0
		for i := size - 1; i >= 0; i-- {
			// combined domain names
			for j := i; j < size; j++ {
				buffer.WriteString(dm[j])
				if j != size-1 {
					buffer.WriteString(".")
				}
			}
			resp.SubDomain = buffer.String()
			// reset buffer
			buffer.Reset()
			// determine whether it is a tld
			if value, ok := tldMap[resp.SubDomain]; ok {
				resp.Tld = value.Tld
				isTLD = true
				continue
			}
			// after finding out the tld, the domain name is the last one
			if isTLD {
				if resp.Domain == "" {
					resp.Domain = resp.SubDomain
				}
				if idx >= level {
					break
				}
				idx++
			}
		}
	} else {
		if tld, ok := tldMap[url]; !ok {
			err = errors.New("Can't get tld from " + url)
		} else {
			resp.Tld = tld.Tld
		}
	}
	pool.Put(buffer)
	if resp.Tld == "" {
		err = errors.New("Can't get tld from " + url)
	} else {
		resp.Label = size
	}

	return
}
