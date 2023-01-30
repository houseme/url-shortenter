package crawler

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

// CollyMain is the main function for the utility package.
func CollyMain() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://go-colly.org/")
}
