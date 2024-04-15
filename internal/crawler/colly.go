// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

// Package crawler is a utility package for crawler.
package crawler

// import (
// 	"fmt"
//
// 	"github.com/gocolly/colly/v2"
// )
//
// // CollyMain is the main function for the utility package.
// func CollyMain() {
// 	c := colly.NewCollector()
//
// 	// Find and visit all links
// 	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 		_ = e.Request.Visit(e.Attr("href"))
// 	})
//
// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL)
// 	})
//
// 	_ = c.Visit("https://go-colly.org/")
// }
