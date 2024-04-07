package yousuu

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

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

		entity := fmt2Entity(cleanedStr)
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

func fmt2Entity(s string) ListEntity {
	attrs := strings.Split(s, " ")

	rank, err := strconv.Atoi(attrs[0])
	if err != nil {
		fmt.Println(err.Error())
		rank = -1
	}
	wordNum, err := wordNumTrans(attrs[2])
	if err != nil {
		fmt.Println(err.Error())
	}

	updateTime, err := timeTrans(attrs[3], attrs[4])
	if err != nil {
		fmt.Println(err.Error())
	}

	grade, gradersNum, err := gradeTrans(attrs[6])
	if err != nil {
		fmt.Println(err.Error())
	}

	res := ListEntity{
		Rank:       rank,
		Name:       attrs[1],
		WordNum:    wordNum,
		UpdateTime: updateTime,
		State:      attrs[5],
		Grade:      grade,
		GradersNum: gradersNum,
		Tags:       attrs[10],
	}

	return res
}

func wordNumTrans(s string) (float32, error) {

	// 先将 "万字" 替换为空格，然后进行字符串拆分
	parts := strings.Split(strings.ReplaceAll(s, "万字", ""), " ")
	// 解析数字部分
	num, err := strconv.ParseFloat(parts[0], 32)
	if err != nil {
		return 0, err
	}

	return float32(num), nil
}

func timeTrans(num string, unit string) (int, error) {
	res, err := strconv.Atoi(num)
	if err != nil {
		return -1, nil
	}

	switch unit {
	case "个月":
		res *= 30
	case "年":
		res *= 365
	}
	return res, nil
}

func gradeTrans(s string) (float32, int, error) {
	// 使用正则表达式匹配数字和人数
	re := regexp.MustCompile(`([\d.]+)\((\d+)人\)`)

	// 使用正则表达式提取评分和人数
	matches := re.FindStringSubmatch(s)

	if len(matches) == 3 {
		score := matches[1]
		numGraders_str := matches[2]
		grade, err := strconv.ParseFloat(score, 32)
		if err != nil {
			return 0, 0, err
		}
		numGraders, err := strconv.Atoi(numGraders_str)
		if err != nil {
			return 0, 0, err
		}
		return float32(grade), numGraders, nil
	} else {
		return 0, 0, errors.New("未找到匹配的评分和评分人数")
	}

}
