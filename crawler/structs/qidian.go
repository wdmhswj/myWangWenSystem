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

// const (
// 	yuepiao         string = "yuepiao"
// 	changxiao       string = "hotsales"
// 	yueduzhishu     string = "readindex"
// 	shuyou          string = "newfans"
// 	tuijian         string = "recom"
// 	shoucang        string = "collect"
// 	gengxin         string = "vipup"
// 	vipshoucang     string = "vipcollect"
// 	qianyuezuozhe   string = "signnewbook"
// 	gongzhongzuozhe string = "pubnewbook"
// 	xinrenqianyue   string = "newsign"
// 	xinrenzuozhe    string = "newauthor"
// )

var Ranks = []string{
	"yuepiao",
	"hotsales",
	"readindex",
	"newfans",
	"recom",
	"collect",
	"vipup",
	"vipcollect",
	"signnewbook",
	"pubnewbook",
	"newsign",
	"newauthor",
}

var TagsPageName = []string{
	"chn21",    // 玄幻
	"chn1",     // 奇幻
	"chn2",     // 武侠
	"chn22",    // 仙侠
	"chn4",     // 都市
	"chn15",    // 现实
	"chn6",     // 军事
	"chn5",     // 历史
	"chn7",     // 游戏
	"chn8",     // 体育
	"chn9",     // 科幻
	"chn20109", // 诸天无限
	"chn10",    // 悬疑
	"chn12",    // 轻小说
}

var TagsMap = map[string]string{
	"玄幻":   "chn21",
	"奇幻":   "chn1",
	"武侠":   "chn2",
	"仙侠":   "chn22",
	"现实":   "chn4",
	"军事":   "chn6",
	"历史":   "chn5",
	"游戏":   "chn7",
	"体育":   "chn8",
	"科幻":   "chn9",
	"诸天无限": "chn20109",
	"悬疑":   "chn10",
	"轻小说":  "chn12",
}

var ReversedMap = map[string]string{
	"chn21":    "玄幻",
	"chn1":     "奇幻",
	"chn2":     "武侠",
	"chn22":    "仙侠",
	"chn4":     "现实",
	"chn6":     "军事",
	"chn5":     "历史",
	"chn7":     "游戏",
	"chn8":     "体育",
	"chn9":     "科幻",
	"chn20109": "诸天无限",
	"chn10":    "悬疑",
	"chn12":    "轻小说",
}
