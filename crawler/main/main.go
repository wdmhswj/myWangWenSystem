package main

import (
	"crawler/analysis"
	"crawler/qidian"
	"crawler/structs"
	"crawler/utils"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	getArchiveAnalyzeByTag("chn22")
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
		// fmt.Println(sourceDir + "目录不存在")
		// return
		err := os.Mkdir(sourceDir, 0750)
		if err != nil {
			fmt.Println("failed to mkdir:", err.Error())
			return
		}
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

func getArchiveAnalyzeByTag(tag string) {
	// 存档
	sourceDir := filepath.Join("./data/", tag)
	newDir := filepath.Join("./data/archive/", tag)
	dataArchive(sourceDir, newDir)

	// 爬取
	placeHolder := "{index}"
	qidian.GetViableRankSpecifyClass(structs.Ranks, tag, placeHolder)

	// 分析
	fmt.Println(structs.ReversedMap[tag] + "综合榜如下：")
	getNewRank(sourceDir)

}
