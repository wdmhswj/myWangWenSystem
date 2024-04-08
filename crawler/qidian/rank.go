package qidian

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func getTodayPopularNovels() (RankingList, error) {

	oneRankList := RankingList{
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

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.qidian.com/rank/readindex/")

	return oneRankList, nil
}

var GetTodayPopularNovels = getTodayPopularNovels
