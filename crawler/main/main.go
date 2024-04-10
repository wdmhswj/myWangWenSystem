package main

import (
	"crawler/analysis"
	"crawler/qidian"
	"crawler/structs"
	"crawler/utils"
	"fmt"
)

func main() {
	// 加载
	filenames := []string{
		"畅销榜本日作品销量排行20240410175834",
		"收藏榜历史总作品收藏数排行20240410175839",
		"书友榜本周新增书友最多作品排行20240410175836",
		"推荐榜本周作品推荐票数排行20240410175838",
		"月票榜以起点平台投出月票为排序依据的榜单20240410175832",
		"阅读指数榜本周阅读指数排行20240410175835",
		"VIP收藏榜VIP作品被加入书架数量的排行20240410175842",
	}
	for i, name := range filenames {
		filenames[i] = name + ".json"
	}
	var lists []structs.RankingList_qidian
	for _, name := range filenames {
		list := utils.LoadJsonAsStruct(name)
		lists = append(lists, list)
	}

	comprehensiveRank := analysis.GetComprehensiveRank(lists)
	for i, name := range comprehensiveRank {
		fmt.Printf("%d  %s\n", i+1, name)
	}



	// fetchRanks()
}

func fetchRanks() {
	placeHolder := "{index}"
	// qidian.GetViableRankSpecifyClasses(structs.Ranks, structs.TagsPageName, placeHolder)
	qidian.GetViableRanks(structs.Ranks, placeHolder)
}
