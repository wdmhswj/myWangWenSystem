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
	BookUrl    string  // 小说页面url
}

type rank_type string

const (
	yuepiao         rank_type = "yuepiao"
	changxiao       rank_type = "hotsales"
	yueduzhishu     rank_type = "readindex"
	shuyou          rank_type = "newfans"
	tuijian         rank_type = "recom"
	shoucang        rank_type = "collect"
	gengxin         rank_type = "vipup"
	vipshoucang     rank_type = "vipcollect"
	qianyuezuozhe   rank_type = "signnewbook"
	gongzhongzuozhe rank_type = "pubnewbook"
	xinrenqianyue   rank_type = "newsign"
	xinrenzuozhe    rank_type = "newauthor"
)
