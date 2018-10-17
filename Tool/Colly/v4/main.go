package main

import (
	"encoding/json"
	"github.com/gocolly/colly/debug"

	"fmt"

	"log"

	"strconv"

	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/gocolly/colly"
)

//证监会行业市盈率

type ZhjhHyShyl struct {
	Hydm string `json:"行业代码"`

	Hymc string `json:"行业名称"`

	Zxsj *float64 `json:"最新数据"`

	Gpjs int `json:"股票家数"`

	Ksjs int `json:"亏损家数"`

	Jygy *float64 `json:"近一个月"`

	Jsgy *float64 `json:"近三个月"`

	Jlgy *float64 `json:"近六个月"`

	Jyn *float64 `json:"近一年"`

	Zhy []*ZhjhHyShyl `json:"细分行业"`
}

func main() {

	var err error

	c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299"

	zjhHyShyl := make([]*ZhjhHyShyl, 0)

	c.OnRequest(func(r *colly.Request) {

		fmt.Printf("%+v\r\n%+v\r\n", *r, *(r.Headers))

	})

	c.OnHTML("td>table.list-div-table>tbody>tr", func(e *colly.HTMLElement) {

		hyShy := ZhjhHyShyl{

			Hydm: e.ChildText("td:first-child"),

			Hymc: e.ChildText("td:nth-child(2)"),
		}

		zxsj, err := strconv.ParseFloat(e.ChildText("td:nth-child(3)"), 64)

		if err == nil {

			hyShy.Zxsj = &zxsj

		}

		gpjs, err := strconv.ParseInt(e.ChildText("td:nth-child(4)"), 10, 32)

		if err == nil {

			hyShy.Gpjs = int(gpjs)

		}

		ksjs, err := strconv.ParseInt(e.ChildText("td:nth-child(5)"), 10, 32)

		if err == nil {

			hyShy.Ksjs = int(ksjs)

		}

		jygy, err := strconv.ParseFloat(e.ChildText("td:nth-child(6)"), 64)

		if err == nil {

			hyShy.Jygy = &jygy

		}

		jsgy, err := strconv.ParseFloat(e.ChildText("td:nth-child(7)"), 64)

		if err == nil {

			hyShy.Jsgy = &jsgy

		}

		jlgy, err := strconv.ParseFloat(e.ChildText("td:nth-child(8)"), 64)

		if err == nil {

			hyShy.Jlgy = &jlgy

		}

		jyn, err := strconv.ParseFloat(e.ChildText("td:nth-child(9)"), 64)

		if err == nil {

			hyShy.Jyn = &jyn

		}

		zjhHyShyl = append(zjhHyShyl, &hyShy)

		hyShy.Zhy = make([]*ZhjhHyShyl, 0)

		e.DOM.Parent().Parent().Next().Find("table.list-div-table>tbody>tr").Each(func(_ int, s *goquery.Selection) {

			zhy := ZhjhHyShyl{

				Hydm: strings.Trim(s.Find("td:nth-child(1)").Text(), "\r\n\t "),

				Hymc: strings.Trim(s.Find("td:nth-child(2)").Text(), "\r\n\t "),
			}

			zxsj, err := strconv.ParseFloat(strings.Trim(s.Find("td:nth-child(3)").Text(), "\r\n\t "), 64)

			if err == nil {

				zhy.Zxsj = &zxsj

			}

			gpjs, err := strconv.ParseInt(strings.Trim(s.Find("td:nth-child(4)").Text(), "\r\n\t "), 10, 32)

			if err == nil {

				zhy.Gpjs = int(gpjs)

			}

			ksjs, err := strconv.ParseInt(strings.Trim(s.Find("td:nth-child(5)").Text(), "\r\n\t "), 10, 32)

			if err == nil {

				zhy.Ksjs = int(ksjs)

			}

			jygy, err := strconv.ParseFloat(strings.Trim(s.Find("td:nth-child(6)").Text(), "\r\n\t "), 64)

			if err == nil {

				zhy.Jygy = &jygy

			}

			jsgy, err := strconv.ParseFloat(strings.Trim(s.Find("td:nth-child(7)").Text(), "\r\n\t "), 64)

			if err == nil {

				zhy.Jsgy = &jsgy

			}

			jlgy, err := strconv.ParseFloat(strings.Trim(s.Find("td:nth-child(8)").Text(), "\r\n\t "), 64)

			if err == nil {

				zhy.Jlgy = &jlgy

			}

			jyn, err := strconv.ParseFloat(strings.Trim(s.Find("td:nth-child(9)").Text(), "\r\n\t "), 64)

			if err == nil {

				zhy.Jyn = &jyn

			}

			hyShy.Zhy = append(hyShy.Zhy, &zhy)

		})

	})

	c.OnScraped(func(_ *colly.Response) {

		bData, _ := json.MarshalIndent(zjhHyShyl, "", "\t")

		fmt.Println(string(bData))

	})

	err = c.Visit("http://www.csindex.com.cn/zh-CN/downloads/industry-price-earnings-ratio?date=2017-12-27&type=zjh1")

	if err != nil {

		log.Fatal(err)

	}

}
