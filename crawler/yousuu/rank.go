package yousuu

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	. "crawler/utils"

	"github.com/gocolly/colly"
)

func getTodayPopularNovels() (RankingList, error) {
	oneRankList := RankingList{
		Time: time.Now(),
		Name: "今日热门小说",
		Url:  "https://www.yousuu.com/rank/today",
	}

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.yousuu.com"),
	)

	// 在回调函数中使用 CSS 选择器查找特定的 div 元素
	c.OnHTML("div.result-item-layout-body", func(e *colly.HTMLElement) {
		// 在这里处理找到的 div 元素
		fmt.Println("Found div with class result-item-layout-body:")
		// // 获取 div 元素的文本内容
		text := e.Text
		// 清洗文本内容
		text = strings.TrimSpace(text) // 去除首尾空白字符
		// text = strings.ReplaceAll(text, "\n", "") // 去除换行符
		// text = strings.ReplaceAll(text, "\t", "") // 去除换行符

		regex := regexp.MustCompile(`\s+`)
		cleanedStr := regex.ReplaceAllString(text, " ")
		fmt.Println("Cleaned String:", cleanedStr)

		entity := Fmt2Entity(cleanedStr)
		oneRankList.Entities = append(oneRankList.Entities, entity)

		// attr := strings.Split(cleanedStr, " ")
		// fmt.Printf("%#v\n", attr)

		// name := e.DOM.Find(".book-name").Text()
		// fmt.Println("name:", name)

		// wordNum := e.DOM.Find("div.list-card-content > div > p:nth-child(2) > span:nth-child(1)").Text()
		// fmt.Println("wordNum:", wordNum)

		// rank := e.DOM.Find(".card-bookInfo-cover .index").Text()
		// fmt.Println("rank:", rank)

		// updateTime := e.DOM.Find("div.list-card-content > div > p:nth-child(2) > span:nth-child(2)").Text()
		// fmt.Println("updateTime:", updateTime)

		// state := e.DOM.Find("div.list-card-content > div > p:nth-child(2) > span:nth-child(3)").Text()
		// fmt.Println("state:", state)

		// grade := e.DOM.Find("div.book-score")
		// fmt.Println("grade:", grade.IsNodes())

		// author := e.DOM.Find(".author-name").Text()
		// fmt.Println("author:", author)

		// // 查找包含图片的 img 元素
		// img := e.DOM.Find("img.BookCoverImage")

		// // 获取图片的 src 属性值
		// src, ok := img.Attr("src")
		// if !ok {
		// 	fmt.Println("Image Source not found")
		// 	return
		// }

		// // 打印图片的 src 属性值
		// fmt.Println("Image Source:", src)
	})

	// c.OnHTML(".rank-view", func(h *colly.HTMLElement) {
	// 	fmt.Println(h.Text)
	// })

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.yousuu.com/rank/today")

	return oneRankList, nil
}

var GetTodayPopularNovels = getTodayPopularNovels
