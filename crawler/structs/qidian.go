package structs

import "time"

type RankingList_qidian struct {
	Time     time.Time
	Name     string
	Url      string
	Entities []ListEntity_qidian
}

type ListEntity_qidian struct {
	Name       string
	ImgUrl     string
	Author     string
	WordNum    float32 // 单位：万
	State      string  // 连载状态
	UpdateTime string  // 20xx-xx-xx xx:xx
	MainTag    string  // 分区
	SubTag     string  // 子分类
	Grade      float32 // 评分
	GradersNum int     // 评分人数
	Rank       int     // 排名
	BookUrl    string	// 小说页面url
}
