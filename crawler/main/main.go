package main

import (
	"crawler/analysis"
	"crawler/qidian"
	"crawler/structs"
	"crawler/utils"
	"fmt"
	"os"
	"strings"
)

func main() {

	getArchiveAnalyze()
}

func fetchRanks() {
	placeHolder := "{index}"
	// qidian.GetViableRankSpecifyClasses(structs.Ranks, structs.TagsPageName, placeHolder)
	qidian.GetViableRanks(structs.Ranks, placeHolder)
}

func getNewRank(dir string) {
	fileNames, err := utils.GetAllFileName(dir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var loadFiles []string
	for _, filename := range fileNames {
		fmt.Println(filename)
		if strings.Contains(filename, "畅销榜") ||
			strings.Contains(filename, "收藏榜历史") ||
			strings.Contains(filename, "书友榜") ||
			strings.Contains(filename, "推荐榜") ||
			strings.Contains(filename, "月票榜") ||
			strings.Contains(filename, "阅读指数榜") ||
			strings.Contains(filename, "VIP收藏榜") {
			loadFiles = append(loadFiles, filename)
		}
		// front := filename[:2]
		// if front == "畅销" ||
		// 	front == "收藏" ||
		// 	front == "书友" ||
		// 	front == "推荐" ||
		// 	front == "月票" ||
		// 	front == "阅读" ||
		// 	front == "VI" {
		// 	loadFiles = append(loadFiles, filename)
		// }
	}

	// // 加载
	// filenames := []string{
	// 	"畅销榜本日作品销量排行20240410230148.json",
	// 	"收藏榜历史总作品收藏数排行20240410230154.json",
	// 	"书友榜本周新增书友最多作品排行20240410230151.json",
	// 	"推荐榜本周作品推荐票数排行20240410230152.json",
	// 	"月票榜以起点平台投出月票为排序依据的榜单20240410230146.json",
	// 	"阅读指数榜本周阅读指数排行20240410230150.json",
	// 	"VIP收藏榜VIP作品被加入书架数量的排行20240410230157.json",
	// }

	var lists []structs.RankingList_qidian
	for _, name := range loadFiles {
		fmt.Println("loadFile:", name)
		list := utils.LoadJsonAsStruct(name, dir)
		lists = append(lists, list)
	}

	comprehensiveRank := analysis.GetComprehensiveRank(lists)
	for i, name := range comprehensiveRank {
		fmt.Printf("%d  %s\n", i+1, name)
	}
}

func dataArchive(sourceDir, newDir string) {
	if ok, _ := utils.FileDirExist(sourceDir); !ok {
		fmt.Println(sourceDir + "目录不存在")
		return
	}

	// 创建新目录
	if ok, _ := utils.FileDirExist(newDir); !ok {
		err := os.Mkdir(newDir, 0750)
		if err != nil {
			fmt.Println("failed to mkdir:", err.Error())
			return
		}
	}

	// 移动
	err := utils.MoveFiles(sourceDir, newDir)
	if err != nil {
		fmt.Println("文件存档失败：", err.Error())
	}
}

func getArchiveAnalyze() {
	dataArchive("./data/", "./data/archive/")
	fetchRanks()
	getNewRank("./data/")
}
