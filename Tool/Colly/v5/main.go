package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func main() {
	c := colly.NewCollector(

		//colly.AllowedDomains("leetcode.com"),
		colly.Debugger(&debug.LogDebugger{}),
		colly.MaxDepth(999),
	)
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"

	c.OnHTML("#wrapper  div.cell", func(e *colly.HTMLElement) {
		e.ForEach("table > tbody > tr > td:nth-child(3) > span:nth-child(1) > a ", func(i int, element *colly.HTMLElement) {
			fmt.Println(element.Text)
		})
	})

	c.OnRequest(func(r *colly.Request) {

		fmt.Println("Visiting", r.URL.String())

	})

	c.Visit("https://studygolang.com/")
}
