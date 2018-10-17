package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	//c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))
	c := colly.NewCollector()



	c.OnHTML("body", func(e *colly.HTMLElement) {

		e.ForEach("tr", func(i int, element *colly.HTMLElement) {
			fmt.Println(element.Text)
		})

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	//c.Visit("https://leetcode.com/problemset/all/")
	c.Visit("https://leetcode.com/problemset/all/")
}
