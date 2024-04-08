package qidian

import "time"

type RankingList struct {
	Time     time.Time
	Name     string
	Url      string
	Entities []ListEntity
}

type ListEntity struct {
	Name       string
	ImgUrl     string
	Author     string
	WordNum    float32 // 单位：万
	State      string  // 连载状态
	UpdateTime string	// 20xx-xx-xx xx:xx
	MainTag    string  // 分区
	SubTag     string  // 子分类
	Grade      float32 // 评分
	GradersNum int     // 评分人数
	Rank       int     // 排名
}
