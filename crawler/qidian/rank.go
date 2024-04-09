package qidian

import (
	"crawler/structs"
	. "crawler/structs"
	"crawler/utils"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

func getTodayPopularNovels(initialUrl, pageTemplete, allowDomain, placeHoler string, pageNumber int) (RankingList_qidian, error) {
	// pageTemplete := "https://www.qidian.com/rank/readindex/page{index}/"
	// pages := []string{
	// 	"https://www.qidian.com/rank/readindex/page1/",
	// 	"https://www.qidian.com/rank/readindex/page2/",
	// 	"https://www.qidian.com/rank/readindex/page3/",
	// 	"https://www.qidian.com/rank/readindex/page4/",
	// 	"https://www.qidian.com/rank/readindex/page5/",
	// }
	var ranklist structs.RankingList_qidian
	// ranklist.Url = "https://www.qidian.com/rank/readindex/"
	ranklist.Url = initialUrl
	ranklist.Time = time.Now()

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains(allowDomain),
	)

	// c.OnHTML("*", func(e *colly.HTMLElement) {
	// 	fmt.Printf("*: %s\n", e.Text)
	// })

	// 排行榜名称
	c.OnHTML("div.rank-header h3.lang", func(e *colly.HTMLElement) {
		if len(ranklist.Name) == 0 {
			fmt.Printf("rank name: %s\n", e.Text)
			ranklist.Name = e.Text
		}

	})

	c.OnHTML("div.rank-body div.rank-view-list div.book-img-text ul li", func(e *colly.HTMLElement) {
		// fmt.Printf("entity body: %s\n", e.Text)
		var entity structs.ListEntity_qidian

		// rank
		var index int
		requestURL := e.Request.URL.String()
		re, err := regexp.Compile(`page(\d+)`)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			match := re.FindStringSubmatch(requestURL)
			if len(match) > 1 {
				fmt.Println("匹配的数字部分:", match[1])
				index, err = strconv.Atoi(match[1])
				if err != nil {
					fmt.Println(err.Error())
				}
			} else {
				fmt.Println("未找到匹配")
			}

		}

		fmt.Println("requestURL:", requestURL)
		rank, _ := strconv.Atoi(e.Attr("data-rid"))
		fmt.Println("old rank:", rank)
		rank += 20 * (index - 1)
		fmt.Println("rank:", rank)
		entity.Rank = rank
		// // img src
		// imgUrl := e.DOM.Find("div.book-img-box").Text()
		// fmt.Println("imgUrl:", imgUrl)

		// book name
		name := e.DOM.Find("div.book-mid-info h2 a").Text()
		fmt.Println("book name:", name)
		entity.Name = name

		// author
		author := e.DOM.Find("div.book-mid-info p.author a.name").Text()
		fmt.Println("author:", author)
		entity.Author = author

		// mainTag
		mainTag := e.DOM.Find("div.book-mid-info p.author a:nth-child(4)").Text()
		fmt.Println("mainTag:", mainTag)
		entity.MainTag = mainTag

		// subTag
		subTag := e.DOM.Find("div.book-mid-info p.author a:nth-child(6)").Text()
		fmt.Println("subTag:", subTag)
		entity.SubTag = subTag

		// state
		state := e.DOM.Find("div.book-mid-info p.author span").Text()
		fmt.Println("state:", state)
		entity.State = state

		// updateTime
		updateTime := e.DOM.Find("div.book-mid-info p.update span").Text()
		fmt.Println("updateTime:", updateTime)
		entity.UpdateTime = updateTime

		// url
		url, _ := e.DOM.Find("div.book-right-info p.btn a.red-btn").Attr("href")
		url = url[2:]
		fmt.Println("url:", url)
		entity.BookUrl = url

		// img
		img, _ := e.DOM.Find("div.book-img-box a img").Attr("src")
		img = img[2:]
		fmt.Println("img:", img)
		entity.ImgUrl = img

		ranklist.Entities = append(ranklist.Entities, entity)

	})

	// c.OnHTML("ul.lbf-pagination-item-list li.lbf-pagination-item a", func(h *colly.HTMLElement) {
	// 	if !h.DOM.HasClass("lbf-pagination-disabled") && h.DOM.HasClass("lbf-pagination-next") {
	// 		fmt.Println("next url:", h.Attr())
	// 		c.Visit()
	// 	}
	// })

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
		r.Headers.Set("Connection", "keep-alive")
		r.Headers.Set("Content-Encoding", "gzip")
		r.Headers.Set("Content-Security-Policy", "frame-ancestors  'self' *.qidian.com  *.hongxiu.com  *.yuewen.com  *.qq.com  *.qdmm.com  *.readnovel.com  *.xs8.cn  *.xxsy.net  *.tingbook.com  *.lrts.me  *.ywurl.cn  *.qdwenxue.com  *.if.qidian.com  www.gameloop.com")
		r.Headers.Set("Server", "Lego Server")
		r.Headers.Set("Cookie", "supportwebp=true; supportWebp=true; qdrsnew=1%7C9%7C0%7C1%7C5; qdrs=1%7C12%7C0%7C1%7C5; _gid=GA1.2.1535091117.1712575717; e1=%7B%22l6%22%3A%22%22%2C%22l1%22%3A3%2C%22pid%22%3A%22qd_p_qidian%22%2C%22eid%22%3A%22qd_A16%22%7D; e2=%7B%22l6%22%3A%22%22%2C%22l1%22%3A3%2C%22pid%22%3A%22qd_p_qidian%22%2C%22eid%22%3A%22qd_A16%22%7D; _csrfToken=BubTnAmHRvG7CaATS2scvK05TQe1bg8CfucJa4tr; traffic_utm_referer=; newstatisticUUID=1712629150_1671393382; fu=2037405195; Hm_lvt_f00f67093ce2f38f215010b699629083=1710904960,1712056007,1712575717,1712629152; _yep_uuid=113fa9e4-ed8e-a554-912f-987940d8c0bd; Hm_lpvt_f00f67093ce2f38f215010b699629083=1712629722; _ga_FZMMH98S83=GS1.1.1712629152.15.1.1712629722.0.0.0; _ga_PFYW0QLV3P=GS1.1.1712629152.15.1.1712629722.0.0.0; _ga=GA1.2.1101290594.1694761175; _gat_gtag_UA_199934072_2=1; w_tsfp=ltvgWVEE2utBvS0Q6aLtk0OpETE7Z2R7xFw0D+M9Os09AqUjW5mH14V6ttfldCyCt5Mxutrd9MVxYnGDUNQsehMSQcSUb5tH1VPHx8NlntdKRQJtA5jcCwEXKrgk6WFCejxcJUHhiml3I9RDyrUyjA0Pu3Yn37ZlCa8hbMFbixsAqOPFm/97DxvSliPXAHGHM3wLc+6C6rgv8LlSgS3A9wqpcgQ2Xusewk+A1SgfDngj4RG7dOldNRytI86vWO0wrTPzwjn3apCs2RYx/UJk6EtuWZaxhCfAPXZMJQhoMAu01L4teqiuZeMi7DFNW/hGSAsW/w0b5rBo6wk=")
		r.Headers.Set("Host", "www.qidian.com")
		r.Headers.Set("Referer", "https://www.qidian.com/")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.0.0")
		r.Headers.Set("sec-ch-ua", "\"Microsoft Edge\";v=\"123\", \"Not:A-Brand\";v=\"8\", \"Chromium\";v=\"123\"")
		r.Headers.Set("Sec-Ch-Ua-Mobile", "?0")
		r.Headers.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
		r.Headers.Set("Sec-Fetch-Dest", "document")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br, zstd")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")

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
		fmt.Println("响应体: ...")
		// fmt.Println(string(r.Body))

	})

	// Start scraping on https://hackerspaces.org
	// c.Visit("https://www.qidian.com/rank/readindex/")
	for i := 1; i <= pageNumber; i++ {
		url := utils.ReplacePlaceholer(placeHoler, pageTemplete, strconv.Itoa(i))
		// fmt.Println("Visit:", url)
		c.Visit(url)
	}

	return ranklist, nil
}

var GetTodayPopularNovels = getTodayPopularNovels

func GetViableRanks(rankNames []string, placeHolder string) {
	for _, name := range rankNames {
		rankList, err := getTodayPopularNovels(
			"https://www.qidian.com/rank/"+name+"/",
			"https://www.qidian.com/rank/"+name+"/page"+placeHolder+"/",
			// "https://www.qidian.com/rank/readindex/page{index}/",
			"www.qidian.com",
			placeHolder,
			5,
		)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			utils.SaveAsJson(rankList.Name+rankList.Time.Format("20060102150405"), rankList)
		}

	}
}
