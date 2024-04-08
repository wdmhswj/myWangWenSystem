package qidian

import (
	. "crawler/structs"
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func getTodayPopularNovels() (RankingList_qidian, error) {

	oneRankList := RankingList_qidian{
		Time: time.Now(),
		Name: "阅读指数榜",
		Url:  "https://www.qidian.com/rank/readindex/",
	}

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.qidian.com"),
	)

	// 在回调函数中使用 CSS 选择器查找特定的 div 元素
	c.OnHTML("div.main-content-wrap", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnHTML("*", func(e *colly.HTMLElement) {
		// 打印整个页面的 HTML 内容
		fmt.Println(e.Text)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Content-Encoding", "gzip")
		r.Headers.Set("Connection", "keep-alive")

	})

	// 在收到每个响应时调用此函数
	c.OnResponse(func(r *colly.Response) {
		// 打印响应的状态码
		fmt.Println("响应状态码:", r.StatusCode)

		// 打印响应头
		fmt.Println("响应头:", r.Headers)
		// for key, value := range r.Headers {
		// 	fmt.Printf("%s: %s\n", key, value)
		// }

		// 打印响应体（HTML 或者其他数据）
		fmt.Println("响应体:")
		fmt.Println(string(r.Body))
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.qidian.com/rank/readindex/")

	return oneRankList, nil
}

var GetTodayPopularNovels = getTodayPopularNovels
