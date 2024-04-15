// Copyright Url-Shortenter Author(https://houseme.github.io/url-shortenter/). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/houseme/url-shortenter.

package crawler

// // GeziyorMain is the main function for the utility package.
// func GeziyorMain() {
// 	geziyor.NewGeziyor(&geziyor.Options{
// 		StartURLs: []string{"https://quotes.toscrape.com/"},
// 		ParseFunc: quotesParse,
// 		Exporters: []export.Exporter{&export.JSON{}},
// 	}).Start()
// }
//
// func quotesParse(g *geziyor.Geziyor, r *client.Response) {
// 	r.HTMLDoc.Find("div.quote").Each(func(i int, s *goquery.Selection) {
// 		g.Exports <- map[string]interface{}{
// 			"text":   s.Find("span.text").Text(),
// 			"author": s.Find("small.author").Text(),
// 		}
// 	})
// 	if href, ok := r.HTMLDoc.Find("li.next > a").Attr("href"); ok {
// 		g.Get(r.JoinURL(href), quotesParse)
// 	}
// }
