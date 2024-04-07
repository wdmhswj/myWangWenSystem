package yousuu

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
	UpdateTime int     // 单位：天
	Tags       string  // 标签，可有多个
	Grade      float32 // 评分
	GradersNum int     // 评分人数
	Rank       int     // 排名
}
